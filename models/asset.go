package models

import "time"

type Asset struct {
	AssetID      int    `gorm:"primaryKey;autoIncrement"`
	Code         string `gorm:"column:asset_code;type:varchar(100);unique;not null"`
	Name         string `gorm:"column:asset_name;type:varchar(255);not null"`
	CategoryID   int    `gorm:"not null"`
	PurchaseYear int
	Condition    string    `gorm:"type:varchar(50)"`
	Status       string    `gorm:"type:varchar(50)"`
	QRCode       string    `gorm:"type:text"`
	Image        string    `gorm:"type:text"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	// FK ke Category
	// Category Category `gorm:"foreignKey:CategoryID;references:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
