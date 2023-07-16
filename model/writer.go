package model

import (
	"gorm.io/gorm"
)

type Writer struct {
	ChatId			string
}

func MigrateWriter(db *gorm.DB) error{
	err := db.AutoMigrate(&Writer{})
	return err
}

