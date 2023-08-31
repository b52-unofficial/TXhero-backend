package data

const (
	QueryUserTransactionSQL     = "SELECT * FROM transaction_info WHERE from_address = $1"
	QueryUserDateTransactionSQL = "SELECT * FROM transaction_info WHERE from_address = $1 and date = $2"
	QueryUserMetaData           = "SELECT count(tx_hash) as total_tx, sum(gas_fee) as total_gas, sum(reward) as total_rewards FROM transaction_info WHERE from_address = $1"
	QueryTransactionByStatus    = "SELECT * FROM transaction_info WHERE status = $1"
	UpdateTxConfirmedSQL        = "UPDATE transaction_info SET status = 1, gas_fee = $2, update_dt = now() WHERE tx_hash = $1"
	InsertNextBidInfoSQL        = "INSERT INTO bid_info(top_bid, total_gas_fee, builder_id, start_timestamp, end_timestamp) VALUES ($1, 0, $2, date_trunc('day', current_timestamp), date_trunc('day', current_timestamp) + interval '1 day')"
)
