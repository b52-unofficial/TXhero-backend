package data

const (
	QueryUserTransactionSQL               = "SELECT * FROM transaction_info WHERE from_address = $1"
	QueryUserDateTransactionSQL           = "SELECT * FROM transaction_info WHERE from_address = $1 and date = $2"
	QueryUserMetaData                     = "SELECT count(tx_hash) as total_tx, sum(gas_fee) as total_gas, sum(reward) as total_rewards FROM transaction_info WHERE from_address = $1"
	QueryTransactionByStatusSQL           = "SELECT * FROM transaction_info WHERE status = $1"
	UpdateTxConfirmedSQL                  = "UPDATE transaction_info SET status = 1, gas_fee = $2, update_dt = now() WHERE tx_hash = $1"
	InsertNextBidInfoSQL                  = "INSERT INTO bid(top_bid, total_gas_fee, builder_id, start_timestamp, end_timestamp) VALUES ($1, 0, $2, date_trunc('day', current_timestamp), date_trunc('day', current_timestamp) + interval '1 day')"
	QueryPrevRoundGasFeeGroupByAddressSQL = "SELECT from_address, sum(gas_fee) as gas_fee FROM transaction_info WHERE status = 1 AND timestamp >= $1 AND timestamp < $2 GROUP BY from_address"
	QueryPrevRoundTotalGasFeeSQL          = "SELECT sum(gas_fee) as gas_fee FROM transaction_info WHERE status = 1 AND  timestamp >= $1 AND timestamp < $2"
	QueryPrevRoundInfoSQL                 = "SELECT round, top_bid, total_gas_fee, builder_id, start_timestamp, end_timestamp FROM bid WHERE start_timestamp >= date_trunc('day', CURRENT_DATE - INTERVAL '1 day') AND end_timestamp <= date_trunc('day', CURRENT_DATE)"
	QueryPrevRoundConfirmedTxsSQL         = "SELECT * FROM transaction_info WHERE status = 1 AND timestamp >= $1 AND timestamp < $2"
	UpdateTxRewardSQL                     = "UPDATE transaction_info SET reward = $2, update_dt = now() WHERE tx_hash = $1"
)
