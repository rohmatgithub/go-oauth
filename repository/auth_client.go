package repository

import (
	"database/sql"
)

type AbstractModel struct {
	CreatedBy sql.NullInt64
	UpdatedBy sql.NullInt64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
type AuthClient struct {
	ClientID       sql.NullString `gorm:"primaryKey"`
	ClientSecret   sql.NullString
	SecretKey      sql.NullString
	GrantType      sql.NullString
	RedirectUri    sql.NullString
	ClientResource []ClientResource `gorm:"foreignKey:ClientID"`
	AbstractModel
}

func (AuthClient) TableName() string {
	return "auth_client"
}
