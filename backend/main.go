package main

import (
	"backend/database"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	r := gin.Default()

	// Routes
	r.POST("/register", routes.Register)
	r.POST("/login", routes.Login)

	r.Run(":8080")
}
