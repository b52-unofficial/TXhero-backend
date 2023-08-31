package data

const (
	QueryUserTransactionSQL     = "SELECT * FROM transaction_info WHERE from_address = $1"
	QueryUserDateTransactionSQL = "SELECT * FROM transaction_info WHERE from_address = $1 and date = $2"
	QueryUserMetaData           = "SELECT count(tx_hash) as total_tx, sum(gas_fee) as total_gas, sum(reward) as total_reward FROM transaction_info WHERE from_address = $1"
)
