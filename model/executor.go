package model

import (
	"time"
	"gorm.io/gorm"
)
type Executor struct {
	gorm.Model
	Employee Employee  `gorm:"foreignKey:OrderNumberID"`
	Activity 		bool 		
	StartDate		time.Time 
	StopDate		time.Time
}

func MigrateExecutor(db *gorm.DB) error{
	err := db.AutoMigrate(&Executor{})
	return err
}