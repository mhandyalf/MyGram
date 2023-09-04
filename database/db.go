package database

import (
	"mygram/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	// Gantilah konfigurasi database sesuai dengan Anda
	dsn := "user=postgres password=123456 dbname=mygram host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto-migrate the models
	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})

	return db
}
