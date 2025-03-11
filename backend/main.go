package main

import (
	"backend/database"
	"backend/middleware"
	"backend/routes"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	r := gin.Default()

	// Enable CORS for React frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.POST("/register", routes.Register)
	r.POST("/login", routes.Login)
	r.GET("/admin-info", middleware.AuthMiddleware("admin"), func(c *gin.Context) {
		username, _ := c.Get("username")
		c.JSON(http.StatusOK, gin.H{"name": username})
	})

	r.GET("/products", routes.GetProducts)
	r.POST("/products", middleware.AuthMiddleware("admin"), routes.CreateProduct)       // ✅ Only admins
	r.DELETE("/products/:id", middleware.AuthMiddleware("admin"), routes.DeleteProduct) // ✅ Only admins

	r.Run(":8080")
}
