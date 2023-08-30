package middleware

import (
	"github.com/b52-unofficial/TXhero-backend/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"log"
)

func Restricted() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.GetConfig().Secret),
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			name := claims["name"].(string)
			log.Printf("user '%s' accessing to '%s'", name, ctx.Request().URI().String())
			return ctx.Next()
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  fiber.StatusUnauthorized,
				"message": "unauthorized",
			})
		},
	})
}
