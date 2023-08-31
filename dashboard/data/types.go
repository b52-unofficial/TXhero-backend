package data

import "time"

type TransactionInfo struct {
	ID        int64     `json:"id" db:"id"`
	TxHash    string    `json:"txHash" db:"tx_hash"`
	GasFee    float64   `json:"gasFee" db:"gas_fee"`
	FromAddr  string    `json:"fromAddr" db:"from_address"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	Status    int       `json:"status" db:"status"`
	Reward    float64   `json:"reward" db:"reward"`
	Claimable bool      `json:"claimable" db:"claimable"`
}

type TransactionMetaData struct {
	TotalTx      int64   `json:"totalTx" db:"total_tx"`
	TotalGas     float64 `json:"totalGas" db:"total_gas"`
	TotalRewards float64 `json:"totalReward" db:"total_reward"`
}
