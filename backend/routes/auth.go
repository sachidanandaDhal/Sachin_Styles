package routes

import (
	"backend/database"
	"backend/models"
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// JWT Secret Key (Change this for production)
var jwtSecret = []byte("my_secret_key")

// Password hashing function
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Password comparison function
func checkPassword(providedPassword, storedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}

// Generate JWT Token
func generateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":           user.ID.Hex(),
		"email":        user.Email,
		"mobileNumber": user.MobileNumber,
		"role":         user.Role,
		"exp":          time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// 📌 Register User API
func Register(c *gin.Context) {
	var input struct {
		Username     string `json:"username" binding:"required"`
		MobileNumber string `json:"mobileNumber" binding:"required"`
		Email        string `json:"email" binding:"required"`
		Password     string `json:"password" binding:"required"`
	}

	// Parse request JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if user already exists
	count, _ := database.UserCollection.CountDocuments(context.TODO(), bson.M{"email": input.Email})
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Hash the password
	hashedPassword, _ := hashPassword(input.Password)

	// Create user object
	user := models.User{
		ID:           primitive.NewObjectID(),
		Username:     input.Username,
		MobileNumber: input.MobileNumber,
		Email:        input.Email,
		Password:     hashedPassword,
		Role:         "user", // Default role is "user"
	}

	// Insert into MongoDB
	_, err := database.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	// Generate JWT token
	token, _ := generateToken(user)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "token": token})
}

// 📌 Login User API
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Parse request JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Find user in database
	var user models.User
	err := database.UserCollection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Verify password
	if !checkPassword(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, _ := generateToken(user)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token, "role": user.Role})
}
