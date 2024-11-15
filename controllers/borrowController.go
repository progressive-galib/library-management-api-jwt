package controllers

import (
    "github.com/gofiber/fiber/v2"
    "library-management-api-jwt/database"
    "library-management-api-jwt/models"
    "time"
)

// BorrowBook handles borrowing a book by a user
func BorrowBook(c *fiber.Ctx) error {
    borrow := new(models.Borrow)
    if err := c.BodyParser(borrow); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }
    borrow.DueDate = time.Now().AddDate(0, 0, 14) // 2 weeks by default
    database.DB.Create(&borrow)
    return c.Status(fiber.StatusCreated).JSON(borrow)
}

// GetUserBorrowedBooks lists all books borrowed by a user
func GetUserBorrowedBooks(c *fiber.Ctx) error {
    userID := c.Params("user_id")
    var borrows []models.Borrow
    database.DB.Where("user_id = ?", userID).Find(&borrows)
    return c.JSON(borrows)
}

