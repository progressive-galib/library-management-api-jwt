package database

import (
    "library-management-api-jwt/config"
    "library-management-api-jwt/models"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
    DB = config.InitDB()
    DB.AutoMigrate(&models.User{}, &models.Book{}, &models.Author{}, &models.Borrow{})
}