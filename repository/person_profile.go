package repository

import "database/sql"

type PersonProfile struct {
	ID        sql.NullInt64 `gorm:"primaryKey"`
	FirstName sql.NullString
	LastName  sql.NullString
	Address1  sql.NullString
	Address2  sql.NullString
	CountryID sql.NullInt64
}
