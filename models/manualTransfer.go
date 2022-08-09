package models

import "errors"

type ManualTransfer struct {
	TransactionId string `gorm:"primaryKey; size: 44; not null" json:"transaction_id"`
	TokenId       string `gorm:"size: 41; not null" json:"token_id"`
	// Token         NFT
	UserID      uint `json:"user"`
	User        User
	Description string `json:"description"`
	Status      int    `gorm:"not null" json:"status"`
}

func GetManualTransfers() ([]ManualTransfer, error) {
	var mts []ManualTransfer
	if err := Db.Find(&mts).Error; err != nil {
		return mts, err
	}
	return mts, nil
}

func GetManualTransferById(transactionId string) (ManualTransfer, error) {
	var mt ManualTransfer

	if err := Db.First(&mt, "transaction_id=?", transactionId); err != nil {
		return mt, errors.New("manual transfer not found")
	}
	return mt, nil
}
