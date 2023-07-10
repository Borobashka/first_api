package model

import (
	"time"
)

type Employee struct{
	OrderNumber		string     `json:"id"`
	Name 			string     `json:"name"`	
	CreationDate	time.Time  `json:"dateTime"`
	Exhausted		bool	   `json:"exhausted" sql:"DEFAULT:false;type:boolean;index" gorm:"not null"` 
}

//var Employees = []employee{
//	{OrderNumber: "1", Name: "Alexey Semenov", CreationDate: "2023-07-06", Exhausted: "false",},
//}