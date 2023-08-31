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
	SmartContractHandler(router)
}

func TxHandler(router *fiber.App) {
	txHandler := router.Group("tx")
	txHandler.Get("user", server.TransactionInfo)
	txHandler.Get("metadata", server.TransactionMetadata)
	txHandler.Get("accumulated_info", server.TransactionAccumulatedInfo)
}

func SmartContractHandler(router *fiber.App) {
	_ = router.Group("builder/v1")
}
