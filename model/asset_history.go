package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssetHistory struct {
	ID          string `gorm:"primaryKey;type:char(36)" json:"id"`
	AssetID     string `gorm:"type:char(36);index" json:"asset_id"`
	Action      string `gorm:"size:100" json:"action"`
	Description string `gorm:"type:text" json:"description"`
	PerformedBy string `gorm:"type:char(36);index" json:"performed_by"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Asset Asset `gorm:"foreignKey:AssetID" json:"asset"`
	User  User  `gorm:"foreignKey:PerformedBy" json:"user"`
}

func (AssetHistory) TableName() string {
	return "asset_histories"
}

func (h *AssetHistory) BeforeCreate(tx *gorm.DB) (err error) {
	h.ID = uuid.New().String()
	return
}
