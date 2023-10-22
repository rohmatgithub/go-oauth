package common

import (
	"database/sql"
	"errors"
	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
	"go-oauth/config"
	"go-oauth/constanta"
	"reflect"
	"strconv"
	"strings"
)

func MigrateSchema(db *sql.DB, pathFile string, schemaName string) error {
	class := "[MigrateSql.go,MigrateSchema]"
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations_"+schemaName, pathFile),
	}
	if db == nil {
		return errors.New("error because db is null")
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	if SQLMigrationResolutionDir == "" {
		box := reflect.Indirect(reflect.ValueOf(migrations)).FieldByName("Box")
		resolution := reflect.Indirect(box).Interface().(*packr.Box)
		splitData := strings.Split(resolution.ResolutionDir, "\\")
		SQLMigrationResolutionDir = strings.Join(splitData[0:len(splitData)-1], "\\")
	}
	logModel := GenerateLogModel(config.ApplicationConfiguration.GetServerConfig().Version, config.ApplicationConfiguration.GetServerConfig().ResourceID)
	logModel.Status = 200
	logModel.Class = class
	logModel.Message = "Applied " + strconv.Itoa(n) + " migrations!"
	PrintLogWithLevel(constanta.LogLevelInfo, logModel)

	return nil
}
