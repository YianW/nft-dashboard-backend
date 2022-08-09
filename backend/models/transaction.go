package models

import (
	"errors"
	"time"
)

type Transaction struct {
	TransactionId string `gorm:"primaryKey; size:44; not null" json:"transaction_id"`
	TokenId       string `gorm:"size:41; not null" json:"token_id"`
	// Token         NFT       `json:"-"`
	Type        int       `gorm:"not null" json:"type"`
	FunctionIdx int8      `gorm:"not null" json:"function_idx"`
	Timestamp   time.Time `gorm:"not null" json:"timestamp"`
	By          string    `gorm:"size:35; not null" json:"by"`
	Receiver    string    `gorm:"size:35" json:"receiver"`
	Status      string    `gorm:"size:120" json:"status"`
}

func GetTransactions() ([]Transaction, error) {
	var ts []Transaction
	if err := Db.Find(&ts).Error; err != nil {
		return ts, err
	}
	return ts, nil
}

func GetTransactionById(transactionId string) (Transaction, error) {
	var t Transaction

	if err := Db.First(&t, transactionId).Error; err != nil {
		return t, errors.New("transaction not found")
	}
	return t, nil
}
