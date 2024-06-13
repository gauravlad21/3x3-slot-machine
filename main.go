package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/slot_machine/config"
	"github.com/slot_machine/database"
	"github.com/slot_machine/routes"
)

func main() {
	config.LoadConfig()
	database.InitMongo(config.AppConfig.MongoURI)
	database.InitRedis(config.AppConfig.RedisURI)

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
