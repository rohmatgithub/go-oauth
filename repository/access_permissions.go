package repository

type AccessPermission struct {
	UserID    int64 `gorm:"primaryKey"`
	CompanyID int64
	BranchID  []int64 `gorm:"type:bigint[]"`
	IsAdmin   bool
}
