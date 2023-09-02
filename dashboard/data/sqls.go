package data

const (
	// Query TransactionInfo Table
	QueryUserTransactionSQL     = "SELECT tx_hash, gas_fee, reward, from_address, timestamp, status, id FROM transaction_info WHERE from_address = $1"
	QueryUserDateTransactionSQL = "SELECT tx_hash, gas_fee, reward, from_address, timestamp, status, id FROM transaction_info WHERE from_address = $1 and date = $2"
	QueryUserMetaDataSQL        = "SELECT coalesce(count(tx_hash), 0) as total_tx, coalesce(sum(gas_fee), 0) as total_gas, coalesce(sum(reward), 0) as total_reward FROM transaction_info WHERE from_address = $1"
	QueryAccumulatedDataSQL     = "SELECT coalesce(sum(reward), 0) as total_reward_amt, coalesce(avg(reward), 0) as avg_reward_amt FROM transaction_info WHERE $1 < timestamp"
	QueryChartInfoSQL           = "SELECT sum(gas_fee) as total_gas_amt, sum(reward) as total_rebate_amt, DATE(timestamp) as timestamp FROM transaction_info WHERE from_address = $1 and DATE(timestamp) >= DATE($2) Group By DATE(timestamp)"
	QueryRoundTxSQL             = "SELECT count(id) as total_tx_cnt FROM transaction_info WHERE timestamp >= $1 and timestamp <= $2"

	// Query Bid Table
	QueryCurrentRoundSQL     = "SELECT round, end_timestamp FROM bid WHERE round = (SELECT MAX(round) FROM bid)"
	QueryPrevRoundSQL        = "SELECT round, start_timestamp, end_timestamp FROM bid WHERE round = $1"
	QueryRoundBuilderInfoSQL = "SELECT b.start_timestamp, b.end_timestamp, b.top_bid, b.total_gas_fee, bd.builder_name, bd.address, bd.description, bd.id FROM bid b join builder bd on b.builder_id = bd.id where b.round = $1"

	// TODO: Change Contract Query
	QueryUserRewardSQL = "SELECT coalesce(sum(reward), 0) as reward FROM reward WHERE address = $1"
	SaveUserRewardSQL  = "INSERT INTO reward (address, reward) VALUES (:address, :reward) ON conflict (address) DO UPDATE SET (address, reward) = (excluded.address, excluded.reward + reward.reward)"
	UserRewardClaimSQL = "UPDATE reward SET reward = 0 WHERE address = $1"
)
