package controllers

import (
	"backend-challenge/configs"
	"backend-challenge/internal/models"

	"github.com/gofiber/fiber/v2"
)

func AddProductToCart(c *fiber.Ctx) error {
	cart := new(models.Cart)

	if err := c.BodyParser(cart); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := configs.GetDB().Create(&cart).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(cart)
}

func GetCartItems(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	var carts []models.Cart

	if err := configs.GetDB().Preload("Product").Where("user_id = ?", userID).Find(&carts).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(carts)
}

func DeleteCartItem(c *fiber.Ctx) error {
	cartID := c.Params("cart_id")
	var cartItem models.Cart

	if err := configs.GetDB().Where("id = ?", cartID).First(&cartItem).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Cart item not found",
		})
	}

	if err := configs.GetDB().Delete(&cartItem).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Cart item deleted successfully",
	})
}