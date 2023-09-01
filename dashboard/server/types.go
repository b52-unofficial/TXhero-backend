package server

import (
	"github.com/b52-unofficial/TXhero-backend/dashboard/data"
	"time"
)

type TransactionInfoData struct {
	ID        int64     `json:"id"`
	TxHash    string    `json:"txHash"`
	GasFee    float64   `json:"gasFeeAmt"`
	FromAddr  string    `json:"fromAddr"`
	Timestamp time.Time `json:"timestamp"`
	Status    int       `json:"status"`
	Reward    float64   `json:"rewardAmt"`
}

type TransactionMetaData struct {
	TotalTx           int64   `json:"totalTxCnt"`
	TotalGas          float64 `json:"totalGasAmt"`
	TotalRewards      float64 `json:"totalRewardAmt"`
	TotalClaimableAmt float64 `json:"totalClaimableAmt"`
}

type TransactionAccumulatedData struct {
	TotalRewardAmt float64 `json:"totalRewardAmt"`
	AvgRewardAmt   float64 `json:"avgRewardAmt"`
}

type RoundBuilderInfo struct {
	Id             int       `json:"id"`
	BuilderName    string    `json:"builderName"`
	BuilderAddr    string    `json:"builderAddr"`
	Description    string    `json:"description"`
	StartTimestamp time.Time `json:"startTimestamp"`
	EndTimestamp   time.Time `json:"endTimestamp"`
	TopBid         float64   `json:"topBid"`
	TotalGasFee    float64   `json:"totalGasFee"`
}

type UserRewardData struct {
	Data []data.UserRewardData `json:"userRewardInfo"`
}

type UserRewardClaimData struct {
	Address string `json:"address"`
}
