package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/slot_machine/database"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "healthy",
	})
}

func LivenessCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "alive",
	})
}

func ReadinessCheck(c *fiber.Ctx) error {
	if err := database.GetMongoClient().Ping(context.Background(), nil); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "unready",
		})
	}
	if _, err := database.GetRedisClient().Ping(context.Background()).Result(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "unready",
		})
	}
	return c.JSON(fiber.Map{
		"status": "ready",
	})
}
