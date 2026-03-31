package models

import "time"

type AssetHistory struct {
	HistoryID int `gorm:"primaryKey"`
	AssetID   int
	OldStatus string
	NewStatus string
	Note      string
	ChangedAt time.Time
}
