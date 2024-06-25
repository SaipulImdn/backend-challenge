package controllers

import (
	"backend-challenge/configs"
	"backend-challenge/internal/models"

	"github.com/gofiber/fiber/v2"
)

type CheckoutRequest struct {
    UserID uint    `json:"user_id"`
    Total  float64 `json:"total"`
}

func Checkout(c *fiber.Ctx) error {
    var req CheckoutRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    checkout := models.Checkout{
        UserID: req.UserID,
        Total:  req.Total,
    }

    if err := configs.GetDB().Create(&checkout).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(checkout)
}
