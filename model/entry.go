package model

import (
	"github.com/michaelcclary/blogsite/model/database"
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Title   string `gorm:"type:text" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

func (entry *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}
