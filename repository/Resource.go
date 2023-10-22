package repository

type Resource struct {
	ResourceID  string `gorm:"primaryKey"`
	Description string
	AbstractModel
}
