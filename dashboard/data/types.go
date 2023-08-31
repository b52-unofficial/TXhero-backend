package data

import (
	"time"
)

type TransactionInfo struct {
	ID        int64     `db:"id"`
	TxHash    string    `db:"tx_hash"`
	GasFee    float64   `db:"gas_fee"`
	FromAddr  string    `db:"from_address"`
	Timestamp time.Time `db:"timestamp"`
	Status    int       `db:"status"`
	Reward    float64   `db:"reward"`
}

type TransactionMetaData struct {
	TotalTx      int64   `db:"total_tx"`
	TotalGas     float64 `db:"total_gas"`
	TotalRewards float64 `db:"total_reward"`
}

type TransactionAccumulatedData struct {
	TotalRewardAmt float64 `db:"total_reward_amt"`
	AvgRewardAmt   float64 `db:"avg_reward_amt"`
}

type UserRewardData struct {
	Reward float64 `db:"reward"`
}
