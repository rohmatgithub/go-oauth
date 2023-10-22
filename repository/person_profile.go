package repository

type PersonProfile struct {
	ID        int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Address1  string
	Address2  string
	CountryID int
}
