package models

import "time"

type Category struct {
	CategoryID int       `gorm:"primaryKey;autoIncrement"`
	Name       string    `gorm:"column:category_name;type:varchar(255);not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`

	// Relasi (optional, tapi bagus untuk preload)
	Assets []Asset `gorm:"foreignKey:CategoryID"`
}
