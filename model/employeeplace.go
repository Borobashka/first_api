package model

import (
	"gorm.io/gorm"
)

type EmployeePlace struct{
	EmplEmail  		string		`gorm:"primaryKey"`
	Role 			string
	Admin			bool		`json:"exhausted" sql:"DEFAULT:false;type:boolean;index" gorm:"not null"` 
	EmplPhone       int
}

func MigrateEmployeePlace(db *gorm.DB) error{
	err := db.AutoMigrate(&EmployeePlace{})
	return err
}