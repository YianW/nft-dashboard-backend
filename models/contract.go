package models

import (
	"errors"
)

type Contract struct {
	ContractId                string `gorm:"primaryKey; size: 35; not null" json:"contract_id"`
	Name                      string `gorm:"size: 50" json:"name"`
	RegistrationTransactionId string `gorm:"size: 44; not null" json:"registration_transaction_id"`
	Maker                     string `gorm:"size: 35; not null" json:"maker"`
	ContractDescription       string `json:"contract_description"`
	Active                    bool   `gorm:"default: true" json:"active"`
}

func (c *Contract) UpdateContract() (*Contract, error) {
	err := Db.Save(&c).Error
	if err != nil {
		return &Contract{}, err
	}
	return c, nil
}

func (c *Contract) SaveContract() (*Contract, error) {
	err := Db.Create(&c).Error
	if err != nil {
		return &Contract{}, err
	}
	return c, nil
}

func GetContracts() ([]Contract, error) {
	var ctrts []Contract
	if err := Db.Find(&ctrts).Error; err != nil {
		return ctrts, err
	}
	return ctrts, nil
}

func GetContractById(tokenId string) (Contract, error) {
	var c Contract
	if err := Db.First(&c, tokenId).Error; err != nil {
		return c, errors.New("contract not found")
	}
	return c, nil
}
