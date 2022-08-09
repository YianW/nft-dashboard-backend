package models

import (
	"gorm.io/gorm"
)

type APIkey struct {
	gorm.Model
	Key string `gorm:"size:255;not null;primaryKey" json:"key"`
}

func (k *APIkey) InsertKey() (*APIkey, error) {
	err := Db.Create(&k).Error
	if err != nil {
		return &APIkey{}, err
	}
	return k, nil
}

func GetAllKeys() ([]APIkey, error) {
	var keys []APIkey
	if err := Db.Find(&keys).Error; err != nil {
		return keys, err
	}
	return keys, nil
}

func (k *APIkey) UpdateKey() (*APIkey, error) {
	err := Db.Save(&k).Error
	if err != nil {
		return &APIkey{}, err
	}
	return k, nil
}

func (k *APIkey) RemoveKey() (*APIkey, error) {
	err := Db.Delete(&k).Error
	if err != nil {
		return &APIkey{}, err
	}
	return k, nil
}
