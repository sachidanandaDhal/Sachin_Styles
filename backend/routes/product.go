package routes

import (
	"backend/database"
	"context"

	"backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(c *gin.Context) {
	userID, exists := c.Get("userID") // Get Admin ID from JWT
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Define request structure
	var request struct {
		Name        string   `json:"name" binding:"required"`
		Description string   `json:"description" binding:"required"`
		Price       string   `json:"price" binding:"required"`
		Images      []string `json:"images"` // Base64 images
	}

	// Parse JSON request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create product object
	product := models.Product{
		ID:          primitive.NewObjectID(),
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Images:      request.Images,
		CreatedBy:   userID.(string), // Store Admin ID
		CreatedAt:   time.Now(),
	}

	_, err := database.ProductCollection.InsertOne(context.TODO(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": product})
}

// 📌 Get All Products (Available to Everyone)
func GetProducts(c *gin.Context) {
	var products []models.Product
	cursor, err := database.ProductCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var product models.Product
		cursor.Decode(&product)
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

// 📌 Delete Product (Admin Only)
func DeleteProduct(c *gin.Context) {
	// Check if user is admin
	userRole, exists := c.Get("role")
	if !exists || userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can delete products"})
		return
	}

	productID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	_, err = database.ProductCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
