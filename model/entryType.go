package model

import (
	"github.com/michaelcclary/blogsite/model/database"
	"gorm.io/gorm"
)

type EntryType struct {
	gorm.Model
	Name string `gorm:"type:text" json:"name"`
}

func (entryType *EntryType) Save() (*EntryType, error) {
	err := database.Database.Create(&entryType).Error
	if err != nil {
		return &EntryType{}, err
	}
	return entryType, nil
}
