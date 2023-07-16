package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"./model"
)


func main() {
	dsn := " host = localhost user = postgres password = 12345 dbname = db_1 port = 5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
 	   panic("failed to connect database")
 	 }

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