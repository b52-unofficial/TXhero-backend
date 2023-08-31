package data

import (
	"time"
)

type TransactionInfo struct {
	ID        int64     `json:"id" db:"id"`
	TxHash    string    `json:"txHash" db:"tx_hash"`
	GasFee    float64   `json:"gasFeeAmt" db:"gas_fee"`
	FromAddr  string    `json:"fromAddr" db:"from_address"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	Status    int       `json:"status" db:"status"`
	Reward    float64   `json:"rewardAmt" db:"reward"`
	Claimable bool      `json:"claimable" db:"claimable"`
}

type TransactionMetaData struct {
	TotalTx      int64   `json:"totalTxCnt" db:"total_tx"`
	TotalGas     float64 `json:"totalGasAmt" db:"total_gas"`
	TotalRewards float64 `json:"totalRewardAmt" db:"total_reward"`
}

type TransactionAccumulatedData struct {
	TotalRewardAmt float64 `json:"totalRewardAmt" db:"total_reward_amt"`
	AvgRewardAmt   float64 `json:"avgRewardAmt" db:"avg_reward_amt"`
}
