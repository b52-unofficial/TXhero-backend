package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/b52-unofficial/TXhero-backend/dashboard/server"
	"github.com/gofiber/fiber/v2"
)

func APIRoute(router *fiber.App) {
	router.Get("/swagger/*", swagger.HandlerDefault)
	router.Get("/health", server.HealthCheck)

	TxHandler(router)
	RewardHandler(router)
	BidHandler(router)
}

func TxHandler(router *fiber.App) {
	txHandler := router.Group("tx")
	txHandler.Get("user", server.TransactionInfo)
	txHandler.Get("metadata", server.TransactionMetadata)
	txHandler.Get("accumulated_info", server.TransactionAccumulatedInfo)
	txHandler.Get("chart_info", server.TransactionChartInfo)
}

func RewardHandler(router *fiber.App) {
	rewardHandler := router.Group("reward")
	rewardHandler.Post("", server.UserReward)
	rewardHandler.Put("claim", server.UserRewardClaim)
}

func BidHandler(router *fiber.App) {
	bidHandler := router.Group("bid")
	bidHandler.Get("current_round", server.CurrentRound)
	bidHandler.Get("rounds", server.PrevRound)
	bidHandler.Get("round_info", server.RoundInfo)
}
