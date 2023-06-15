package model

import (
	"github.com/michaelcclary/blogsite/model/database"
	"gorm.io/gorm"
)

type BlogEntry struct {
	gorm.Model
	Title   string `gorm:"type:text" json:"title"`
	Card    string `gorm:"type:text" json:"card"`
	Content string `gorm:"type:text" json:"content"`
}

func (b *BlogEntry) Create() (*BlogEntry, error) {
	err := database.Database.Create(&b).Error
	if err != nil {
		return &BlogEntry{}, err
	}
	return b, nil
}

//	func (entryType *EntryType) Update(updatedEntryType EntryType) (*EntryType, error) {
//		err := database.Database.Model(&entryType).Updates(updatedEntryType).Error
//		if err != nil {
//			return &EntryType{}, err
//		}
//		return entryType, nil
//	}
//
//	func (entryType *EntryType) Delete() {
//		database.Database.Delete(&entryType)
//	}
func GetBlogEntry(id uint) (BlogEntry, error) {
	var entry BlogEntry
	err := database.Database.Where("ID=?", id).First(&entry).Error
	if err != nil {
		return BlogEntry{}, err
	}
	return entry, nil
}

//
//func FindAllEntryTypes() ([]EntryType, error) {
//	var entryTypes []EntryType
//	err := database.Database.Find(&entryTypes).Error
//	if err != nil {
//		return []EntryType{}, err
//	}
//
//	return entryTypes, nil
//}
