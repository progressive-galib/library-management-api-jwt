// main.go
package main

import (
    "github.com/gofiber/fiber/v2"
    "library-management-api-jwt/database"
    "library-management-api-jwt/routes"
)

func main() {
    app := fiber.New()
    database.InitDatabase()
    routes.SetupRoutes(app)
    app.Listen(":3000")
}