package handlers

import (
	"context"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/slot_machine/database"
	"github.com/slot_machine/models"
)

var (
	totalSpins  int64
	totalPayout int64
)

func PlaySlotMachine(c *fiber.Ctx) error {
	var playReq models.PlayRequest
	if err := json.Unmarshal(c.Body(), &playReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}

	// Example slot machine logic
	rand.Seed(time.Now().UnixNano())
	outcome := rand.Intn(100) // Example probability logic
	payout := 0

	if outcome < 5 { // 5% chance to win big
		payout = 1000
		logStatistic("win big", payout)
	} else if outcome < 20 { // 15% chance to win small
		payout = 100
		logStatistic("win small", payout)
	} else if outcome < 30 { // 10% chance to lose big
		payout = -500
		logStatistic("lose big", payout)
	} else { // 70% chance to lose small
		payout = -50
		logStatistic("lose small", payout)
	}

	totalSpins++
	totalPayout += int64(payout)

	// Update RTP calculation
	rtp := float64(totalPayout) / float64(totalSpins)
	if rtp > 0.975 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "RTP limit exceeded",
		})
	}

	return c.JSON(models.PlayResponse{
		Result:  "win/lose",
		Credits: payout,
	})
}

func logStatistic(result string, payout int) {
	key := "stat:" + result
	database.GetRedisClient().IncrBy(context.Background(), key, int64(payout))
}
