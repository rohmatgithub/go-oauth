package repository

import "database/sql"

type PersonProfile struct {
	ID        sql.NullInt64  `gorm:"primaryKey"`
	FirstName sql.NullString `gorm:"column:first_name"`
	LastName  sql.NullString `gorm:"column:last_name"`
	Address1  sql.NullString `gorm:"column:address_1"`
	Address2  sql.NullString `gorm:"column:address_2"`
	CountryID sql.NullInt64  `gorm:"column:country_id"`
	AbstractModel
}

func (PersonProfile) TableName() string {
	return "person_profile"
}
