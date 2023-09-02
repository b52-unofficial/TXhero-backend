package data

import (
	"time"
)

type TransactionInfo struct {
	ID        int64     `db:"id" json:"id"`
	TxHash    string    `db:"tx_hash" json:"txHash"`
	GasFee    float64   `db:"gas_fee" json:"gasFee"`
	FromAddr  string    `db:"from_address" json:"fromAddr"`
	Timestamp time.Time `db:"timestamp" json:"timestamp"`
	Status    int       `db:"status" json:"status"`
	Reward    float64   `db:"reward" json:"reward"`
}

type TransactionMetaData struct {
	TotalTx      int64   `db:"total_tx" json:"totalTxCnt"`
	TotalGas     float64 `db:"total_gas" json:"totalGasAmt"`
	TotalRewards float64 `db:"total_reward" json:"totalRewardAmt"`
}

type TransactionAccumulatedData struct {
	TotalRewardAmt float64 `db:"total_reward_amt" json:"totalRewardAmt"`
	AvgRewardAmt   float64 `db:"avg_reward_amt" json:"avgRewardAmt"`
}

type TxChartData struct {
	Date           time.Time `db:"timestamp" json:"date"`
	TotalGasAmt    float64   `db:"total_gas_amt" json:"totalGasAmt"`
	TotalRebateAmt float64   `db:"total_rebate_amt" json:"totalRebateAmt"`
}

type UserRewardData struct {
	Address string  `db:"address"  json:"address"`
	Reward  float64 `db:"reward" json:"reward"`
}

type RoundInfo struct {
	Round          int64     `db:"round" json:"round"`
	StartTimestamp time.Time `db:"start_timestamp" json:"startTimestamp"`
	EndTimestamp   time.Time `db:"end_timestamp" json:"endTimestamp"`
	TotalTxCount   int64     `db:"total_tx_cnt" json:"totalTxCnt"`
}

type RoundBuilderInfo struct {
	Id             int       `db:"id" json:"id"`
	BuilderName    string    `db:"builder_name" json:"builderName"`
	BuilderAddr    string    `db:"address" json:"builderAddr"`
	Description    string    `db:"description" json:"description"`
	StartTimestamp time.Time `db:"start_timestamp" json:"startTimestamp"`
	EndTimestamp   time.Time `db:"end_timestamp" json:"endTimestamp"`
	TopBid         float64   `db:"top_bid" json:"topBid"`
	TotalGasFee    float64   `db:"total_gas_fee" json:"totalGasFee"`
}
