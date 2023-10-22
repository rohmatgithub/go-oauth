package repository

type ClientResource struct {
	ID          int `gorm:"primaryKey"`
	ClientID    string
	ResourceID  string
	Authorities string
	AbstractModel
}
