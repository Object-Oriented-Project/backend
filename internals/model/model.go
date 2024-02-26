package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
    ID            uint   `gorm:"primaryKey"`
    Username      string `gorm:"unique;not null"`
    PasswordHash  string `gorm:"not null"`
    Salt          string `gorm:"not null"`
}

// Menu model
type Menu struct {
	gorm.Model
    ID              uint    `gorm:"primaryKey"`
    ItemName        string  `gorm:"not null"`
    ItemPriceSmall  float64 `gorm:"not null"`
    ItemPriceMedium float64 `gorm:"not null"`
    ItemPriceLarge  float64 `gorm:"not null"`
	ItemType        string  `gorm:"not null"`
    IsRecommended   bool    `gorm:"not null"`
    PictureName     string
    ItemDescription string
}