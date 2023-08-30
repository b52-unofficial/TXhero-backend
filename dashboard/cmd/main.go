package main

import (
	"github.com/b52-unofficial/TXhero-backend/common/db"
	"github.com/b52-unofficial/TXhero-backend/config"
	"github.com/b52-unofficial/TXhero-backend/dashboard/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"sync"

	_ "github.com/b52-unofficial/TXhero-backend/docs"
)

func main() {
	// DB connection
	db.ConnectDB()

	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)
	go RunAPI()
}

func RunAPI() {
	cfg := config.GetConfig()
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET",
	}))
	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	router.APIRoute(app)

	err := app.Listen(cfg.App.Host + ":" + cfg.App.Port)
	if err != nil {
		log.Println(err)
		return
	}
}
