package model

import (
	"time"
	"gorm.io/gorm"
)

type Employee struct{
	gorm.Model
	OrderNumber		string     `gorm:"primaryKey"`
	Name 			string    
	CreationDate	time.Time  `json:"dateTime"`
	Exhausted		bool	   `json:"exhausted" sql:"DEFAULT:false;type:boolean;index" gorm:"not null"` 
}

func MigrateEmployee(db *gorm.DB) error{
	err := db.AutoMigrate(&Employee{})
	return err
}
