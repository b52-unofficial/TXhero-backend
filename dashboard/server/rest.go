package server

import (
	"github.com/b52-unofficial/TXhero-backend/dashboard/data"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"time"
)

func HealthCheck(ctx *fiber.Ctx) error {
	if err := ctx.SendString("OK"); err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"status": "ok",
	})
}

func TransactionInfo(ctx *fiber.Ctx) error {
	userAddr := ctx.Query("address")
	if userAddr == "" {
		return ctx.Status(404).JSON(fiber.Map{
			"err": "user address is not nil",
		})
	}

	date := ctx.Query("date")
	res, err := data.GetTransactionData(userAddr, date)
	if err != nil {
		log.Println(err)
	}

	return ctx.JSON(res)
}

func TransactionMetadata(ctx *fiber.Ctx) error {
	userAddr := ctx.Query("address")
	if userAddr == "" {
		return ctx.Status(404).JSON(fiber.Map{
			"err": "user address is not nil",
		})
	}

	res, err := data.GetTransactionMetaData(userAddr)
	if err != nil {
		log.Println(err)
	}
	return ctx.JSON(res)
}

func TransactionAccumulatedInfo(ctx *fiber.Ctx) error {
	avgMonth := ctx.Query("avg_month")
	if avgMonth == "" {
		avgMonth = "3"
	}

	tmp, err := strconv.ParseInt(avgMonth, 10, 64)
	month := time.Now().AddDate(0, int(-tmp), 0)

	res, err := data.GetTransactionAccumulatedInfo(month)
	if err != nil {
		log.Println(err)
	}

	return ctx.JSON(res)
}
