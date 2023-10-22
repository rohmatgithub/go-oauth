package repository

import "time"

type AbstractModel struct {
	CreatedBy int
	UpdatedBy int
	CreatedAt time.Time
	UpdatedAt time.Time
}
type AuthClient struct {
	ClientID     string `gorm:"primaryKey"`
	ClientSecret string
	SecretKey    string
	GrantType    string
	RedirectUri  string
	AbstractModel
}
