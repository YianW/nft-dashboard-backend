package models

import (
	"gorm.io/gorm"
	"time"
)

type NFTMintHistory struct {
	gorm.Model
	ContractId       string      `json:"contract_id"`
	Contract         Contract    `json:"-"`
	TokenDescription string      `json:"token_description"`
	Attachment       string      `json:"attachment"`
	Status           int         `json:"status"` // TODO: Enum this field
	ExecuteTime      time.Time   `json:"execute_time"`
	TransactionId    string      `json:"transaction_id"`
	Transaction      Transaction `json:"-"`
}
