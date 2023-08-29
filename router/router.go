package router

import "github.com/gofiber/fiber/v2"

func APIRoute(router *fiber.App) {
	// health check
	router.Get("/health", func(ctx *fiber.Ctx) error {
		if err := ctx.SendString("OK"); err != nil {
			return err
		}
		return nil
	})

	TxHandler(router)
	BuilderHandler(router)
}

func TxHandler(router *fiber.App) {
	_ = router.Group("tx/v1")
}

func BuilderHandler(router *fiber.App) {
	_ = router.Group("builder/v1")
}
