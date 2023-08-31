package server

import (
	"github.com/b52-unofficial/TXhero-backend/dashboard/data"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"time"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
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
		return ctx.Status(500).JSON(fiber.Map{
			"err": err,
		})
	}

	txInfoData := &TransactionInfoData{
		ID:        res[0].ID,
		TxHash:    res[0].TxHash,
		GasFee:    res[0].GasFee,
		FromAddr:  res[0].FromAddr,
		Timestamp: res[0].Timestamp,
		Status:    res[0].Status,
		Reward:    res[0].Reward,
	}
	return ctx.JSON(txInfoData)
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
		return ctx.Status(500).JSON(fiber.Map{
			"err": err,
		})
	}

	reward, err := data.GetUserRewardData(userAddr)
	if err != nil {
		log.Println(err)
		return err
	}

	txMetaData := &TransactionMetaData{
		TotalTx:           res[0].TotalTx,
		TotalGas:          res[0].TotalGas,
		TotalRewards:      res[0].TotalRewards,
		TotalClaimableAmt: reward[0].Reward,
	}
	return ctx.JSON(txMetaData)
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
		return ctx.Status(500).JSON(fiber.Map{
			"err": err,
		})
	}

	return ctx.JSON(res)
}
