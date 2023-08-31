package data

const (
	QueryUserTransactionSQL     = "SELECT * FROM transaction_info WHERE from_address = $1"
	QueryUserDateTransactionSQL = "SELECT * FROM transaction_info WHERE from_address = $1 and date = $2"
	QueryUserMetaDataSQL        = "SELECT count(tx_hash) as total_tx, sum(gas_fee) as total_gas, sum(reward) as total_reward FROM transaction_info WHERE from_address = $1"
	QueryAccumulatedDataSQL     = "SELECT coalesce(sum(reward), 0) as total_reward_amt, coalesce(avg(reward), 0) as avg_reward_amt FROM transaction_info WHERE $1 < timestamp"
)
