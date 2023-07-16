package model

import (
	"time"
	"gorm.io/gorm"
)
type Document struct {
	Id 				int `gorm:"primaryKey"`
	Year 			int `gorm:"primaryKey"`
	Path			string
	name 			string
	CreationDate 	time.Time
}

func MigrateDocument(db *gorm.DB) error{
	err := db.AutoMigrate(&Document{})
	return err
}
