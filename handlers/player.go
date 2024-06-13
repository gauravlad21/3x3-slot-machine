package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/slot_machine/database"
	"github.com/slot_machine/models"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func CreatePlayer(c *fiber.Ctx) error {
	var player models.Player
	if err := json.Unmarshal(c.Body(), &player); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}

	player.ID = primitive.NewObjectID().Hex()
	_, err := database.GetMongoClient().Database("slotmachine").Collection("players").InsertOne(context.Background(), player)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot create player",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(player)
}

func GetPlayer(c *fiber.Ctx) error {
	id := c.Params("id")

	player := models.Player{}
	cacheKey := "player:" + id

	if val, err := database.GetRedisClient().Get(context.Background(), cacheKey).Result(); err == nil {
		if err := json.Unmarshal([]byte(val), &player); err == nil {
			return c.JSON(player)
		}
	}

	filter := bson.M{"_id": id}
	if err := database.GetMongoClient().Database("slotmachine").Collection("players").FindOne(context.Background(), filter).Decode(&player); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "player not found",
		})
	}

	if val, err := json.Marshal(player); err == nil {
		database.GetRedisClient().Set(context.Background(), cacheKey, val, 0)
	}

	return c.JSON(player)
}

func SuspendPlayer(c *fiber.Ctx) error {
	id := c.Params("id")

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": "suspended"}}

	result, err := database.GetMongoClient().Database("slotmachine").Collection("players").UpdateOne(context.Background(), filter, update)
	if err != nil || result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "player not found",
		})
	}

	database.GetRedisClient().Del(context.Background(), "player:"+id)

	return c.JSON(fiber.Map{
		"status": "player suspended",
	})
}
