package main

import (
	"backend-challenge/configs"
	"backend-challenge/internal/models"
	"backend-challenge/internal/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    app := fiber.New()

    configs.ConnectDB()

    configs.ConnectRedis()

    configs.GetDB().AutoMigrate(&models.Product{}, &models.Cart{}, &models.User{}, &models.Checkout{})

    routes.SetupRoutes(app)

    port := os.Getenv("PORT")
    if port == "" {
        port = "5000" 
    }
    log.Fatal(app.Listen(":" + port))
}
