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

type Order struct {
    Data  string `json:"data"`
    From  string `json:"from"`
    To    string `json:"to"`
    CustomerID uint `json:"customer_id"`
    CustomerName string `json:"customer_name"`
    OrderID uint `json:"order_id"`
    OrderStatus OrderStatus `json:"order_status"`
    OrderTable string `json:"order_table"`
    OrderItems []OrderItem `json:"order_items"`
}

type OrderItem struct {
    MenuID uint `json:"menu_id"`
    Quantity int `json:"quantity"`
    Size ItemSize `json:"size"`
    Sweetness ItemSweetness `json:"sweetness"`
    SpicyLevel ItemSpicyLevel `json:"spicy_level"`
    Note string `json:"note"`
}

type OrderStatus string

type ItemSize string

type ItemSweetness string

type ItemSpicyLevel string

const (
    OrderStatusOnGoing OrderStatus = "ONGOING"
    OrderStatusFinished OrderStatus = "FINISHED"
)

const (
    ItemSizeSmall ItemSize = "SMALL"
    ItemSizeMedium ItemSize = "MEDIUM"
    ItemSizeLarge ItemSize = "LARGE"
)

const (
    ItemSweetnessNormal ItemSweetness = "NORMAL"
    ItemSweetnessLess ItemSweetness = "LESS"
    ItemSweetnessMore ItemSweetness = "MORE"
)

const (
    ItemSpicyLevelNormal ItemSpicyLevel = "NORMAL"
    ItemSpicyLevelLess ItemSpicyLevel = "LESS"
    ItemSpicyLevelMore ItemSpicyLevel = "MORE"
)