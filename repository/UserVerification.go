package repository

type UserVerification struct {
	ID               int `gorm:"primaryKey"`
	UserID           int
	Email            string
	EmailCode        string
	EmailExpires     int
	EmailVerifiedAt  int
	Phone            string
	PhoneCode        string
	PhoneExpires     int
	PhoneVerifiedAt  int
	ForgetCode       string
	ForgetExpires    int
	ForgetVerifiedAt int
	AbstractModel
}
