package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Maintenance struct {
	ID          string     `gorm:"primaryKey;type:char(36)" json:"id"`
	AssetID     string     `gorm:"type:char(36);index" json:"asset_id"`
	ReportedBy  string     `gorm:"type:char(36);index" json:"reported_by"`
	Description string     `gorm:"type:text" json:"description"`
	Status      string     `gorm:"type:enum('pending','progress','done');default:'pending'" json:"status"`
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Asset Asset `gorm:"foreignKey:AssetID" json:"asset"`
	User  User  `gorm:"foreignKey:ReportedBy" json:"user"`
}

func (Maintenance) TableName() string {
	return "maintenances"
}

func (m *Maintenance) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	return
}
