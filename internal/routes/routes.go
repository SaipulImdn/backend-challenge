package routes

import (
	"backend-challenge/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/products/:category", controllers.GetProductsByCategory)
	app.Post("/cart", controllers.AddProductToCart)
	app.Get("/cart/:user_id", controllers.GetCartItems)
	app.Delete("/cart/:cart_id", controllers.DeleteCartItem)
	app.Post("/checkout", controllers.Checkout)
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
}
