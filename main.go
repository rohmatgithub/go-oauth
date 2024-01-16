package main

import (
	"encoding/json"
	"fmt"
	"go-oauth/common"
	"go-oauth/config"
	"go-oauth/dto"
	"go-oauth/router"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	var arguments = "development"
	args := os.Args
	if len(args) > 1 {
		arguments = args[1]
	}

	config.GenerateConfiguration(arguments)
	err := common.SetServerAttribute()
	if err != nil {
		fmt.Println("ERROR common server attribute : ", err)
		os.Exit(3)
	}

	common.Validation = common.NewGoValidator()
	dto.GenerateValidOperator()
	loadBundleI18N()
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

func loadBundleI18N() {
	prefixPath := config.ApplicationConfiguration.GetLanguageDirectoryPath()
	var err error

	//------------ constanta bundle
	common.ConstantaBundle = i18n.NewBundle(language.Indonesian)
	common.ConstantaBundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err = common.ConstantaBundle.LoadMessageFile(prefixPath + "/common/constanta/en-US.json")
	readError(err)

	_, err = common.ConstantaBundle.LoadMessageFile(prefixPath + "/common/constanta/id-ID.json")
	readError(err)

	//------------ constanta bundle
	common.ErrorBundle = i18n.NewBundle(language.Indonesian)
	common.ErrorBundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err = common.ErrorBundle.LoadMessageFile(prefixPath + "/common/error/en-US.json")
	readError(err)

	_, err = common.ErrorBundle.LoadMessageFile(prefixPath + "/common/error/id-ID.json")
	readError(err)
}

func readError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
