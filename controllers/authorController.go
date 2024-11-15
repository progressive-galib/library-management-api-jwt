package controllers

import (
    "github.com/gofiber/fiber/v2"
    "library-management-api-jwt/database"
    "library-management-api-jwt/models"
)

// AddAuthor handles adding a new author
func AddAuthor(c *fiber.Ctx) error {
    author := new(models.Author)
    if err := c.BodyParser(author); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }
    database.DB.Create(&author)
    return c.Status(fiber.StatusCreated).JSON(author)
}

// ListAuthors lists all authors in the database
func ListAuthors(c *fiber.Ctx) error {
    var authors []models.Author
    database.DB.Find(&authors)
    return c.JSON(authors)
}

