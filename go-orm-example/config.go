package go_orm_example

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

const (
	dsn = "host=localhost user=postgres password=root dbname=db_belajar_golang port=5432 search_path=oauth sslmode=disable TimeZone=Asia/Shanghai"
)

func ConnectDB() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Adjust the output writer as needed
		logger.Config{
			SlowThreshold: time.Second, // Set the threshold for slow query logging
			LogLevel:      logger.Info, // Set the log level to Info to log queries
			Colorful:      true,        // Enable colored output
		})
	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		//PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "oauth.", // schema name
			//SingularTable: false,
		},
		Logger: newLogger})
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
