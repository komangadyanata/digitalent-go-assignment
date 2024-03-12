package models

// import "gorm.io/gorm"

type Item struct {
	// gorm.Model
	ItemID      uint   `gorm:"primaryKey" json:"lineItemId"`
	ItemCode    string `gorm:"not null;unique;type:varchar(10)"`
	Description string
	Quantity    uint
	OrderID     uint
}
