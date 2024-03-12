package models

import (
	"time"
	// "gorm.io/gorm"
)

type Order struct {
	// gorm.Model
	OrderID      uint   `gorm:"primaryKey"`
	CustomerName string `gorm:"not null;type:varchar(255)"`
	Items        []Item
	OrderedAt    time.Time
}
