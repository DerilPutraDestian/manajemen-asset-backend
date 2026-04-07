package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Asset struct {
	ID         string    `gorm:"primaryKey;type:char(36)" json:"id"`
	AssetCode  string    `gorm:"size:100;uniqueIndex;not null" json:"asset_code"`
	Name       string    `gorm:"size:255;not null" json:"name"`
	CategoryID string    `gorm:"type:char(36);index" json:"category_id"`
	Status     string    `gorm:"type:enum('available','broken','maintenance');default:'available'" json:"status"`
	Condition  string    `gorm:"type:enum('good','fair','poor');default:'good'" json:"condition"`
	QRCode     string    `gorm:"size:255" json:"qr_code"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Category     Category       `gorm:"foreignKey:CategoryID" json:"category"`
	Maintenances []Maintenance  `gorm:"foreignKey:AssetID" json:"maintenances,omitempty"`
	Histories    []AssetHistory `gorm:"foreignKey:AssetID" json:"histories,omitempty"`
}

func (Asset) TableName() string {
	return "assets"
}

func (a *Asset) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New().String()
	return
}
