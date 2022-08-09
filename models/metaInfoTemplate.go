package models

import (
	"errors"

	"gorm.io/gorm"
)

type MetaInfoTemplate struct {
	gorm.Model
	TemplateName string `gorm:"size:255;not null;primaryKey" json:"template_name"`
	MetaKey      string `gorm:"not null" json:"meta_key"`
}

func GetTemplateByName(name string) (MetaInfoTemplate, error) {
	var m MetaInfoTemplate

	if err := Db.First(&m, "template_name = ?", name).Error; err != nil {
		return m, errors.New("template not found")
	}

	return m, nil
}
