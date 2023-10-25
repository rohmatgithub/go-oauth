package router

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	recoverfiber "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go-oauth/config"
)

func Router() error {
	app := fiber.New(fiber.Config{
		//Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Oauth App v1.0.0",
		ColorScheme:   fiber.Colors{Green: ""},
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})
	app.Use(requestid.New())
	app.Use(recoverfiber.New())
	app.Use(middleware)
	//file, err := os.OpenFile("fiber.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalf("error opening file: %v", err)
	//}
	//iw := io.MultiWriter(os.Stdout, file)
	//defer file.Close()
	//app.Use(logger.New(logger.Config{
	//	Format:     "[${time}] pid:${pid}, request-id:${locals:requestid}, status:${status}, method:${method}, path:${path}, error-message:[${error}]\n",
	//	TimeFormat: time.DateTime,
	//	TimeZone:   "Asia/Jakarta",
	//	Output:     iw,
	//}))

	credentialsRouter(app)
	app.Use(NotFoundHandler)
	return app.Listen(fmt.Sprintf(":%d", config.ApplicationConfiguration.GetServerConfig().Port))
}