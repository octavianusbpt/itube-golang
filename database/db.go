package database

import (
	"os"

	"github.com/octavianusbpt/itube-golang/helpers"
	"github.com/octavianusbpt/itube-golang/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Initialize database connection
func InitializeDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	helpers.PanicIfError(err)
}

func SyncDatabase() {
	// Migrate database
	DB.AutoMigrate(&models.User{}, &models.Video{})
}
