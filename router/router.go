package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
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
	})

	return app.Listen(fmt.Sprintf(":%d", config.ApplicationConfiguration.GetServerConfig().Port))
}
