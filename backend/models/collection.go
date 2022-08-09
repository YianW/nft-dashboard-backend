package models

import (
	"errors"

	"gorm.io/gorm"
)

// TODO: add relationship between collection.MetaInfoType and MetaInfoTemplate
type Collection struct {
	gorm.Model
	CollectionName string `gorm:"size:255;not null;primaryKey" json:"collection_name"`
	MetaInfoType   string `gorm:"size:255;not null" json:"metainfo_type"`
	DesFileKey     string `gorm:"not null" json:"des_file_key"`
}

func (c *Collection) InsertCollect() (*Collection, error) {
	err := Db.Create(&c).Error
	if err != nil {
		return &Collection{}, err
	}
	return c, nil
}

func GetAllCollect() ([]Collection, error) {
	var collections []Collection
	if err := Db.Find(&collections).Error; err != nil {
		return collections, err
	}
	return collections, nil
}

func GetCollectByName(name string) (Collection, error) {
	var c Collection

	if err := Db.First(&c, "collection_name = ?", name).Error; err != nil {
		return c, errors.New("collection not found")
	}

	return c, nil
}
