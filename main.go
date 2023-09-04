package main

import (
	"mygram/database"
	"mygram/handlers"
	"mygram/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Inisialisasi database
	db = database.SetupDatabase()

	// Initialize Gin
	r := gin.Default()

	authHandler := handlers.NewAuthHandler(db)

	// Define routes
	r.POST("/login", authHandler.Login)

	// Group endpoint yang memerlukan autentikasi
	authGroup := r.Group("/auth")
	authGroup.Use(middleware.LoginAuth)
	{
		authGroup.GET("/photos", authHandler.GetAllPhotos)
		authGroup.GET("/comments", authHandler.GetAllComments)
		authGroup.GET("/socialmedia", authHandler.GetAllSocialMedia)
	}

	// Endpoint untuk registrasi
	r.POST("/register", authHandler.Register)

	// Run the server
	r.Run(":8080")
}
