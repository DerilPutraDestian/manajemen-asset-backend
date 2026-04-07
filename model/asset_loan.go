package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssetLoan struct {
	ID         string `gorm:"primaryKey;type:char(36)" json:"id"`
	AssetID    string `gorm:"type:char(36);index;not null" json:"asset_id"`
	EmployeeID string `gorm:"type:char(36);index;not null" json:"employee_id"`

	LoanDate   time.Time  `json:"loan_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"`

	Status string `gorm:"type:enum('borrowed','returned');default:'borrowed'" json:"status"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Asset    Asset    `gorm:"foreignKey:AssetID;->:false" json:"asset"`
	Employee Employee `gorm:"foreignKey:EmployeeID;->:false" json:"employee"`
}

func (AssetLoan) TableName() string {
	return "asset_loans"
}

func (l *AssetLoan) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.New().String()
	return
}
