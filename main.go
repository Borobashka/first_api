package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  )

type Employee struct{
	OrderNumber		string     `gorm:"primaryKey"`
	Name 			string    
//	CreationDate	time.Time  `json:"dateTime"`
	Exhausted		bool	   `json:"exhausted" sql:"DEFAULT:false;type:boolean;index" gorm:"not null"` 
}

func main() {
	dsn := " host = localhost user = postgres password = 12345 dbname = db_1 port = 5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
 	   panic("failed to connect database")
 	 }

	//миграция
	db.AutoMigrate(&Employee{})
	
	//создание
	db.Create(&Employee{OrderNumber: "1", Name: "Alex", Exhausted: true})

	//чтение
	var employee Employee
	db.First(&employee, "1")
	db.First(&employee, "order_number = ?", "1")

	//обновление
	db.Model(&employee).Update("Name", "Misha")

	//удаление
	//db.Delete(&employee, "1")

}