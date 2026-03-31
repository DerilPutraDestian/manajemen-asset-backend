package models

import "time"

type AssetLoan struct {
	LoanID     int `gorm:"primaryKey"`
	AssetID    int
	EmployeeID int
	LoanDate   time.Time
	ReturnDate time.Time
	Status     string
	CreatedAt  time.Time
}
