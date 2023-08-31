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

func GetTransactionDataByStatus(status int) ([]*TransactionInfo, error) {
	database := db.GetDB()
	var txInfo []*TransactionInfo
	err := database.Select(&txInfo, QueryTransactionByStatus, status)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func UpdateTxConfirmed(txHash string, gasFee float64) error {
	database := db.GetDB()
	_, err := database.Exec(UpdateTxConfirmedSQL, txHash, gasFee)
	if err != nil {
		return err
	}
	return nil
}

func InsertNextBidInfo(bidInfo *BidInfo) error {
	database := db.GetDB()
	_, err := database.Exec(InsertNextBidInfoSQL, bidInfo.TopBid, bidInfo.BuilderId)
	if err != nil {
		return err
	}
	return nil
}
