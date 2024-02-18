package router

import (
	"fmt"
	"go-oauth/config"
	"go-oauth/endpoint"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))
	//app.Use(recoverfiber.New())
	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				customErrorHandler(c, fmt.Errorf("%v", r))
			}
		}()
		return c.Next()
	})
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

	oauth := app.Group("/v1/oauth")
	credentialsRouter(oauth)
	usersRouter(oauth)

	master := app.Group("/v1/master", endpoint.MiddlewareOtherService)
	masterDataRouter(master)

	trans := app.Group("/v1/trans", endpoint.MiddlewareOtherService)
	transactionRouter(trans)

	app.Use(NotFoundHandler)
	return app.Listen(fmt.Sprintf(":%d", config.ApplicationConfiguration.GetServerConfig().Port))
}
