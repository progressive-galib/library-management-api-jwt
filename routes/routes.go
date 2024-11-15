package routes

import (
    "github.com/gofiber/fiber/v2"
    "library-management-api-jwt/controllers"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api/v1")

    // Book routes
    api.Post("/books", controllers.AddBook)        // Add a new book
    api.Get("/books", controllers.ListBooks)       // List all books

    // Borrow routes
    api.Post("/borrow", controllers.BorrowBook)    // Borrow a book
    api.Get("/borrowed/:user_id", controllers.GetUserBorrowedBooks) // List books borrowed by a user

    // Additional routes can go here, e.g., for authors
    api.Post("/authors", controllers.AddAuthor)    // Add a new author
    api.Get("/authors", controllers.ListAuthors)   // List all authors
}

