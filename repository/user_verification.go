package repository

import "database/sql"

type UserVerification struct {
	ID               sql.NullInt64 `gorm:"primaryKey"`
	UserID           sql.NullInt64
	Email            sql.NullString
	EmailCode        sql.NullString
	EmailExpires     sql.NullInt64
	EmailVerifiedAt  sql.NullInt64
	Phone            sql.NullString
	PhoneCode        sql.NullString
	PhoneExpires     sql.NullInt64
	PhoneVerifiedAt  sql.NullInt64
	ForgetCode       sql.NullString
	ForgetExpires    sql.NullInt64
	ForgetVerifiedAt sql.NullInt64
	AbstractModel
}
