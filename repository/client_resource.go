package repository

import "database/sql"

type ClientResource struct {
	ID          sql.NullInt64 `gorm:"primaryKey"`
	ClientID    sql.NullString
	ResourceID  sql.NullString
	Authorities sql.NullString
	AbstractModel
}

func (ClientResource) TableName() string {
	return "client_resource"
}
