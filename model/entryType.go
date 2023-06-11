package model

import (
	"github.com/michaelcclary/blogsite/model/database"
	"gorm.io/gorm"
)

type EntryType struct {
	gorm.Model
	Name string `gorm:"type:text" json:"name"`
}

func (entryType *EntryType) Create() (*EntryType, error) {
	err := database.Database.Create(&entryType).Error
	if err != nil {
		return &EntryType{}, err
	}
	return entryType, nil
}

func (entryType *EntryType) Update(updatedEntryType EntryType) (*EntryType, error) {
	err := database.Database.Model(&entryType).Updates(updatedEntryType).Error
	if err != nil {
		return &EntryType{}, err
	}
	return entryType, nil
}

func (entryType *EntryType) Delete() {
	database.Database.Delete(&entryType)
}

func FindEntryTypeById(id uint) (EntryType, error) {
	var entryType EntryType
	err := database.Database.Where("ID=?", id).First(&entryType).Error
	if err != nil {
		return EntryType{}, err
	}
	return entryType, nil
}

func FindAllEntryTypes() ([]EntryType, error) {
	var entryTypes []EntryType
	err := database.Database.Find(&entryTypes).Error
	if err != nil {
		return []EntryType{}, err
	}

	return entryTypes, nil
}
