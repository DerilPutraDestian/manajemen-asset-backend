package models

import "time"

type Employee struct {
	EmployeeID int    `gorm:"primaryKey"`
	Name       string `gorm:"column:employee_name"`
	Email      string
	Phone      string
	CreatedAt  time.Time
}
