package controllers

import (
	"backend-challenge/configs"
	"backend-challenge/internal/models"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func GetProductsByCategory(c *fiber.Ctx) error {
    category := c.Params("category")
    redisClient := configs.GetRedisClient()
    ctx := configs.GetRedisContext()

    cachedProducts, err := redisClient.Get(ctx, category).Result()
    if err == redis.Nil {
        var products []models.Product
        if err := configs.GetDB().Where("category = ?", category).Find(&products).Error; err != nil {
            return c.Status(500).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        productsJSON, err := json.Marshal(products)
        if err != nil {
            log.Println("Failed to marshal products:", err)
        }
        err = redisClient.Set(ctx, category, productsJSON, 0).Err()
        if err != nil {
            log.Println("Failed to cache products:", err)
        }

        return c.JSON(products)
    } else if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to retrieve data from Redis",
        })
    }

    var products []models.Product
    err = json.Unmarshal([]byte(cachedProducts), &products)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to unmarshal cached data",
        })
    }

    return c.JSON(products)
}
