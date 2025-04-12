package models

import (
	"time"

	"gorm.io/gorm"
)

type Budget struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Descript  string         `json:"descript"`
	Valor     float64        `json:"valor"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
