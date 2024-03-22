package model

import (
	"gorm.io/gorm"
)

// User model
// @Model
// @Title User
// @Description User model
// @Success 200 {object} model.User
// @Router /user [get]
// @Router /user [post]
// @Router /user/{id} [get]
// @Router /user/{id} [put]
// @Router /user/{id} [delete]
// @Param id path int true "User ID"
// @Param user body model.User true "User"

type User struct {
	gorm.Model
    ID            uint   `gorm:"primaryKey"`
    Username      string `gorm:"unique;not null"`
    PasswordHash  string `gorm:"not null"`
    Salt          string `gorm:"not null"`
}

// Menu model
// @Model
// @Title Menu
// @Description Menu model
// @Success 200 {object} model.Menu
// @Router /menu [get]
// @Router /menu [post]
// @Router /menu/{id} [get]
// @Router /menu/{id} [put]
// @Router /menu/{id} [delete]
// @Param id path int true "Menu ID"
// @Param menu body model.Menu true "Menu"

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