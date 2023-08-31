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
	CreateDt  time.Time `json:"createDt" db:"create_dt"`
	UpdateDt  time.Time `json:"updateDt" db:"update_dt"`
}

type TransactionMetaData struct {
	TotalTx      int64   `json:"totalTx" db:"total_tx"`
	TotalGas     float64 `json:"totalGas" db:"total_gas"`
	TotalRewards float64 `json:"totalRewards" db:"total_rewards"`
}

type BidInfo struct {
	Round          int64     `json:"round" db:"round"`
	TopBid         float64   `json:"topBid" db:"top_bid"`
	TotalGasFee    float64   `json:"totalGasFee" db:"total_gas_fee"`
	BuilderId      int64     `json:"builderId" db:"builder_id"`
	StartTimestamp time.Time `json:"startTimestamp" db:"start_timestamp"`
	EndTimestamp   time.Time `json:"endTimestamp" db:"end_timestamp"`
	CreateDt       time.Time `json:"createDt" db:"create_dt"`
	UpdateDt       time.Time `json:"updateDt" db:"update_dt"`
}
