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
	err := database.Select(&txMetadata, QueryUserMetaData, userAddr)
	if err != nil {
		return nil, err
	}
	return txMetadata, nil
}
