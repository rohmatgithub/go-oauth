package repository

import "database/sql"

type Users struct {
	ID              sql.NullInt64 `gorm:"primaryKey"`
	Username        sql.NullString
	Password        sql.NullString
	Salt            sql.NullString
	Email           sql.NullString
	Phone           sql.NullString
	Status          sql.NullString
	Locale          sql.NullString
	ClientID        sql.NullString
	PersonProfileID sql.NullInt64
	AuthClient      AuthClient    `gorm:"foreignKey:ClientID"`
	PersonProfile   PersonProfile `gorm:"foreignKey:PersonProfileID"`
	ClientResource  ClientResource
	AbstractModel
}
