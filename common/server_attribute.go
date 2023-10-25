package common

import (
	"database/sql"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
	"io"
	"os"
)

var (
	ConnectionDB              *sql.DB
	GormDB                    *gorm.DB
	SQLMigrationResolutionDir string
)

//type logWriter struct {
//}
//
//func (writer logWriter) Write(bytes []byte) (int, error) {
//	return fmt.Print(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + string(bytes))
//	//return os.Stdout.Write([]byte(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + string(bytes)))
//}

func SetServerAttribute() error {
	err := ConnectDB()
	if err != nil {
		return err
	}

	// set log fiber
	// Output to ./test.log file
	file, _ := os.OpenFile("fiber.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(iw)
	return nil
}
