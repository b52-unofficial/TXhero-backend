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

func GetTransactionMetaData(userAddr string) (*TransactionMetaData, error) {
	database := db.GetDB()
	var txMetadata *TransactionMetaData
	err := database.QueryRow(QueryUserMetaDataSQL, userAddr).Scan(&txMetadata)
	if err != nil {
		return nil, err
	}
	return txMetadata, nil
}

func GetTransactionAccumulatedInfo(month time.Time) (*TransactionAccumulatedData, error) {
	database := db.GetDB()

	var txAccumulatedData *TransactionAccumulatedData
	err := database.QueryRow(QueryAccumulatedDataSQL, month).Scan(&txAccumulatedData)
	if err != nil {
		return nil, err
	}
	return txAccumulatedData, nil
}

func GetUserRewardData(userAddr string) (*UserRewardData, error) {
	database := db.GetDB()
	var userReward *UserRewardData
	err := database.QueryRow(QueryUserRewardSQL, userAddr).Scan(&userReward)
	if err != nil {
		return nil, err
	}
	return userReward, nil
}

func GetCurrentRound() (*RoundInfo, error) {
	database := db.GetDB()
	var roundInfo *RoundInfo
	err := database.QueryRow(QueryCurrentRoundSQL).Scan(&roundInfo)
	if err != nil {
		return nil, err
	}
	return roundInfo, nil
}

func GetPrevRound(round string) (*RoundInfo, error) {
	database := db.GetDB()
	var roundInfo *RoundInfo
	err := database.QueryRow(QueryPrevRoundSQL, round).Scan(&roundInfo)
	if err != nil {
		return nil, err
	}

	return roundInfo, nil
}

func GetRoundInfo(round string) ([]*RoundBuilderInfo, error) {
	database := db.GetDB()
	var roundBuilder []*RoundBuilderInfo
	err := database.Select(&roundBuilder, QueryRoundBuilderInfoSQL, round)
	if err != nil {
		return nil, err
	}

	return roundBuilder, nil
}
