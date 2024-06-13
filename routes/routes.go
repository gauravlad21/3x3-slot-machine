package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/slot_machine/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/players", handlers.CreatePlayer)
	app.Get("/players/:id", handlers.GetPlayer)
	app.Put("/players/:id/suspend", handlers.SuspendPlayer)
	app.Post("/play", handlers.PlaySlotMachine)

	app.Get("/health", handlers.HealthCheck)
	app.Get("/liveness", handlers.LivenessCheck)
	app.Get("/readiness", handlers.ReadinessCheck)
}
