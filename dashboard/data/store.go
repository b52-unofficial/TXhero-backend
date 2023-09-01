package data

import (
	"github.com/b52-unofficial/TXhero-backend/common/db"
)

func SaveUserReward(userData []UserRewardData) error {
	database := db.GetDB()

	tx := database.MustBegin()
	for _, res := range userData {
		_, err := tx.NamedExec(SaveUserRewardSQL, &res)
		if err != nil {
			return err
		}
	}
	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserReward(userAddr string) error {
	database := db.GetDB()

	tx := database.MustBegin()
	_, err := tx.Exec(UserRewardClaimSQL, &userAddr)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
