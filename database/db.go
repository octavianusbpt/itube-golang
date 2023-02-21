package database

import (
	"fmt"
	"os"

	"github.com/octavianusbpt/itube-golang/helpers"
	"github.com/octavianusbpt/itube-golang/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Initialize database connection
func InitializeDatabase() {
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	fmt.Println(DB)
	helpers.PanicIfError(err)
}

// Migrate database
func SyncDB() {
	DB.AutoMigrate(&models.User{}, &models.Video{})
}
