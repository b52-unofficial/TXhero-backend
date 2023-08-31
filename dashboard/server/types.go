package server

import "time"

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
