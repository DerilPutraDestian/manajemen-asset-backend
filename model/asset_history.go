package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssetHistory struct {
	ID        string    `gorm:"primaryKey;type:char(36)" json:"id"`
	AssetID   string    `gorm:"type:char(36);index" json:"asset_id"`
	OldStatus string    `gorm:"type:enum('available','broken','maintenance')" json:"old_status"`
	NewStatus string    `gorm:"type:enum('available','broken','maintenance')" json:"new_status"`
	Note      string    `gorm:"size:50" json:"note"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Asset Asset `gorm:"foreignKey:AssetID;references:ID" json:"asset"`
}

func (h *AssetHistory) BeforeCreate(tx *gorm.DB) (err error) {
	h.ID = uuid.New().String()
	return
}
