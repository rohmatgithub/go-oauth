package common

import (
	"fmt"
	"go-oauth/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func ConnectDB() error {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Adjust the output writer as needed
		logger.Config{
			SlowThreshold: time.Second, // Set the threshold for slow query logging
			LogLevel:      logger.Info, // Set the log level to Info to log queries
			Colorful:      true,        // Enable colored output
		})
	GormDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: config.ApplicationConfiguration.GetPostgresqlConfig().Address + fmt.Sprintf(" search_path=%s", config.ApplicationConfiguration.GetPostgresqlConfig().DefaultSchema),
		//PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "oauth.", // schema name
			//SingularTable: false,
		},
		Logger: newLogger,
	})
	if err != nil {
		return err
	}

	ConnectionDB, err = GormDB.DB()
	return err
}
