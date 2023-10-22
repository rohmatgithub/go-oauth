package common

import (
	"database/sql"
	"gorm.io/gorm"
)

var (
	ConnectionDB              *sql.DB
	GormDB                    *gorm.DB
	SQLMigrationResolutionDir string
)

func SetServerAttribute() error {
	err := ConnectDB()
	if err != nil {
		return err
	}

	return nil
}
