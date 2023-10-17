package go_orm_example

import (
	"time"
)

type Categories struct {
	ID   int    `gorm:"primaryKey"`
	Code string `gorm:"unique"`
	Name string
	//CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID              int    `gorm:"primaryKey"`
	Username        string `gorm:"unique"`
	PersonProfileID int
	PersonProfile   PersonProfile `gorm:"foreignKey:PersonProfileID"`
}

type PersonProfile struct {
	ID          int `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Address     string
	PhoneNumber string
}

type Product struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique"`
	Name string
}
type Order struct {
	ID          uint   `gorm:"primaryKey"`
	OrderNumber string `gorm:"unique"`
	OrderDate   time.Time
	OrderItem   []OrderItem
}

type OrderItem struct {
	//gorm.Model
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Product   Product //`gorm:"foreignKey:ProductID"`
	NetAmount float64
}
