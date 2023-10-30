package repository

import "database/sql"

type AuthToken struct {
	ID          sql.NullInt64 `gorm:"primaryKey"`
	Tokens      sql.NullString
	ExpiredTime sql.NullInt64
	ValueToken  sql.NullString
	CreatedAt   sql.NullTime
}

func (AuthToken) TableName() string {
	return "auth_token"
}
