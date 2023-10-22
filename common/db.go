package common

import (
	"fmt"
	"go-oauth/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDB() error {
	var err error
	GormDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: config.ApplicationConfiguration.GetPostgresqlConfig().Address + fmt.Sprintf(" search_path=%s", config.ApplicationConfiguration.GetPostgresqlConfig().DefaultSchema),
		//PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "oauth.", // schema name
			//SingularTable: false,
		}})
	if err != nil {
		return err
	}

	ConnectionDB, err = GormDB.DB()
	return err
}
