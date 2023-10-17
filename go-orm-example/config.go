package go_orm_example

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	dsn = "host=localhost user=postgres password=root dbname=db_belajar_golang port=5432 search_path=oauth sslmode=disable TimeZone=Asia/Shanghai"
)

func ConnectDB() (*gorm.DB, error) {
	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		//PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "oauth.", // schema name
			//SingularTable: false,
		}})
	if err != nil {
		return nil, err
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}
	return DB, nil
}
