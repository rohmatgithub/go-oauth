package repository

type Users struct {
	ID              int `gorm:"primaryKey"`
	Username        string
	Password        string
	Salt            string
	Email           string
	Phone           string
	Status          string
	Locale          string
	ClientID        string
	PersonProfileID int
	AbstractModel
}
