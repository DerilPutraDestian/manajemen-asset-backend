package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID        string    `gorm:"primaryKey;type:char(36)" json:"id"`
	Name      string    `gorm:"column:employee_name" json:"name"`
	Email     string    `gorm:"column:email" json:"email"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return
}
