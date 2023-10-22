package main

import (
	"fmt"
	"go-oauth/common"
	"go-oauth/config"
	"go-oauth/router"
	"os"
)

func main() {
	var arguments = "development"
	args := os.Args
	if len(args) > 1 {
		arguments = args[1]
	}

	config.GenerateConfiguration(arguments)
	common.SetLoggerServer(config.ApplicationConfiguration.GetLogFile())
	err := common.SetServerAttribute()
	if err != nil {
		fmt.Println("ERROR common server attribute : ", err)
		os.Exit(3)
	}

	err = common.MigrateSchema(common.ConnectionDB, config.ApplicationConfiguration.GetSqlMigrateDirPath(), config.ApplicationConfiguration.GetPostgresqlConfig().DefaultSchema)
	if err != nil {
		fmt.Println("ERROR migrate sql : ", err)
		os.Exit(3)
	}

	err = router.Router()
	if err != nil {
		fmt.Println("ERROR router : ", err)
		os.Exit(3)
	}
}
