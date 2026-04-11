package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Maintenance struct {
	ID                string     `gorm:"primaryKey;type:char(36)" json:"id"`
	AssetID           string     `gorm:"type:char(36);index" json:"asset_id"`
	IssueDescription  string     `gorm:"type:text" json:"issue_description"`
	MaintenanceStatus string     `gorm:"type:enum('pending','progress','done');default:'pending'" json:"maintenance_status"`
	StartDate         *time.Time `json:"start_date,omitempty"`
	EndDate           *time.Time `json:"end_date,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Asset Asset `gorm:"foreignKey:AssetID" json:"asset"`
}

func (Maintenance) TableName() string {
	return "maintenances"
}

func (m *Maintenance) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	return
}
