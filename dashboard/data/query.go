package data

import (
	"github.com/b52-unofficial/TXhero-backend/common/db"
	"time"
)

func GetTransactionData(userAddr string, date string) ([]*TransactionInfo, error) {
	database := db.GetDB()
	var txInfo []*TransactionInfo

	if date == "" {
		err := database.Select(&txInfo, QueryUserTransactionSQL, userAddr)
		if err != nil {
			return nil, err
		}
	} else if date != "" {
		tmp, err := time.Parse(time.DateOnly, date)
		if err != nil {
			return nil, err
		}
		err = database.Select(&txInfo, QueryUserDateTransactionSQL, userAddr, tmp)
		if err != nil {
			return nil, err
		}
	}

	return txInfo, nil
}

func GetTransactionMetaData(userAddr string) ([]*TransactionMetaData, error) {
	database := db.GetDB()
	var txMetadata []*TransactionMetaData
	err := database.Select(&txMetadata, QueryUserMetaDataSQL, userAddr)

	return txMetadata, err
}

func GetTransactionAccumulatedInfo(month time.Time) (*TransactionAccumulatedData, error) {
	database := db.GetDB()

	var txAccumulatedData TransactionAccumulatedData
	err := database.QueryRow(QueryAccumulatedDataSQL, month).Scan(&txAccumulatedData.AvgRewardAmt, &txAccumulatedData.TotalRewardAmt)

	return &txAccumulatedData, err
}

func GetUserRewardData(userAddr string) (UserRewardData, error) {
	database := db.GetDB()
	var userReward UserRewardData
	err := database.QueryRow(QueryUserRewardSQL, userAddr).Scan(&userReward.Reward)

	return userReward, err
}

func GetCurrentRound() (RoundInfo, error) {
	database := db.GetDB()
	var roundInfo RoundInfo
	err := database.QueryRow(QueryCurrentRoundSQL).Scan(&roundInfo.Round, &roundInfo.EndTimestamp)

	return roundInfo, err
}

func GetPrevRound(round string) (RoundInfo, error) {
	database := db.GetDB()
	var roundInfo RoundInfo
	err := database.QueryRow(QueryPrevRoundSQL, round).Scan(&roundInfo.Round, &roundInfo.EndTimestamp)

	return roundInfo, err
}

func GetRoundInfo(round string) ([]*RoundBuilderInfo, error) {
	database := db.GetDB()
	var roundBuilder []*RoundBuilderInfo
	err := database.Select(&roundBuilder, QueryRoundBuilderInfoSQL, round)

	return roundBuilder, err
}
