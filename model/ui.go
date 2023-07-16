package model

import (
	"gorm.io/gorm"
)

type UI struct {
	UiId			string `gorm:"primaryKey"`
	Name			string
	Type 			string
	Create 			bool
	Read			bool
	Update 			bool
	Delete 			bool 
}

func MigrateUI(db *gorm.DB) error{
	err := db.AutoMigrate(&UI{})
	return err
}