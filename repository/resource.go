package repository

import "database/sql"

type Resource struct {
	ResourceID  sql.NullString `gorm:"primaryKey"`
	Description sql.NullString
	AbstractModel
}
