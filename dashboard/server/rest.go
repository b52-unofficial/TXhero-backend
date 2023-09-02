package server

import (
	"fmt"
	"github.com/b52-unofficial/TXhero-backend/dashboard/data"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"strings"
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
		return err
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
		return err
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
		TotalClaimableAmt: reward.Reward,
	}
	return ctx.JSON(txMetaData)
}

func TransactionAccumulatedInfo(ctx *fiber.Ctx) error {
	avgMonth := ctx.Query("avg_month", "3")
	tmp, err := strconv.ParseInt(avgMonth, 10, 64)
	month := time.Now().AddDate(0, int(-tmp), 0)

	res, err := data.GetTransactionAccumulatedInfo(month)
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

func TransactionChartInfo(ctx *fiber.Ctx) error {
	userAddr := ctx.Query("address")
	period := ctx.Query("period")

	tmp := strings.Split(period, "-")
	if len(tmp) < 2 {
		return fmt.Errorf("")
	}

	num, err := strconv.ParseInt(tmp[0], 10, 64)
	var date time.Time
	switch tmp[1] {
	case "year":
		date = time.Now().AddDate(-int(num), 0, 0)
	case "month":
		date = time.Now().AddDate(0, -int(num), 0)
	case "day":
		date = time.Now().AddDate(0, 0, -int(num))
	case "week":
		date = time.Now().AddDate(0, 0, -int(num)*7)
	default:
		return fmt.Errorf("%s is not exist", tmp[1])
	}

	res, err := data.GetTransactionChartInfo(userAddr, date)
	if err != nil {
		return err
	}
	return ctx.JSON(res)
}

func CurrentRound(ctx *fiber.Ctx) error {
	res, err := data.GetCurrentRound()
	if err != nil {
		return err
	}
	return ctx.JSON(res)
}

func PrevRound(ctx *fiber.Ctx) error {
	round := ctx.Query("round")
	if round == "" {
		return ctx.Status(404).JSON(fiber.Map{
			"err": "round is not nil",
		})
	}
	res, err := data.GetPrevRound(round)
	if err != nil {
		return err
	}
	if res.Round == 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"err": "round is not found",
		})
	}
	return ctx.JSON(res)
}

func RoundInfo(ctx *fiber.Ctx) error {
	round := ctx.Query("round")
	res, err := data.GetRoundInfo(round)
	if err != nil {
		return err
	}
	return ctx.JSON(res)
}

func UserReward(ctx *fiber.Ctx) error {
	var rewards []*data.UserRewardData
	err := ctx.BodyParser(&rewards)
	if err != nil {
		return err
	}
	err = data.SaveUserReward(rewards)
	return err
}

func UserRewardClaim(ctx *fiber.Ctx) error {
	var userAddr UserRewardClaimData

	err := ctx.BodyParser(&userAddr)
	if err != nil {
		return err
	}

	if userAddr.Address == "" {
		return ctx.Status(404).JSON(fiber.Map{
			"err": fmt.Sprintf("user address is nil"),
		})
	}

	err = data.UpdateUserReward(userAddr.Address)
	return err
}
