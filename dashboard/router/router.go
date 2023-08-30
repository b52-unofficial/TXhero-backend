package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func APIRoute(router *fiber.App) {
	// health check
	router.Get("/health", func(ctx *fiber.Ctx) error {
		if err := ctx.SendString("OK"); err != nil {
			return err
		}
		return nil
	})

	router.Get("/swagger/*", swagger.HandlerDefault)

	TxHandler(router)
	SmartContractHandler(router)
}

func TxHandler(router *fiber.App) {
	_ = router.Group("tx/v1")
}

func SmartContractHandler(router *fiber.App) {
	_ = router.Group("builder/v1")
}
