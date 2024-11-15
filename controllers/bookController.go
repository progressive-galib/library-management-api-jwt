package controllers

import (
    "github.com/gofiber/fiber/v2"
    "library-management-api-jwt/database"
    "library-management-api-jwt/models"
)

// AddBook handles adding a new book to the database
func AddBook(c *fiber.Ctx) error {
    book := new(models.Book)
    if err := c.BodyParser(book); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }
    database.DB.Create(&book)
    return c.Status(fiber.StatusCreated).JSON(book)
}

// ListBooks handles listing all books in the database
func ListBooks(c *fiber.Ctx) error {
    var books []models.Book
    database.DB.Find(&books)
    return c.JSON(books)
}

