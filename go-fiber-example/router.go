package go_fiber_example

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
)

func Router() {
	app := fiber.New(fiber.Config{
		//Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Oauth App v1.0.0",
		ColorScheme:   fiber.Colors{Green: ""},
	})

	startTest(app)
	subRoute1(app)
	subRoute2(app)
	err := app.Listen(":3000")
	if err != nil {
		log.Panic(err)
	}
}

func startTest(app *fiber.App) {
	app.Get("/hello/:value?", func(c *fiber.Ctx) error {
		return c.SendString("hello " + c.Params("value"))
	})

	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
		// => API path: user/john
	})

	app.Get("/test/error/:value?", func(c *fiber.Ctx) error {
		s := c.Params("value")
		if s != "" {
			return c.SendString("Hello :" + s)
		}
		return fiber.NewError(782, "Value cannot empty")
	})

	// Match any request
	// redirect example
	app.Use(func(c *fiber.Ctx) error {
		//fmt.Println("masuk middleware")
		url := c.OriginalURL()
		if strings.Contains(url, "//") {
			fmt.Println("masuk redirect")
			url = strings.ReplaceAll(url, "//", "/")
			return c.Redirect(url)
		}
		return c.Next()
	})

	//// Match request starting with /api
	//app.Use("/api", func(c *fiber.Ctx) error {
	//	fmt.Println("masuk middleware")
	//	return c.Next()
	//})
}

func subRoute1(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Route("/user",
		func(routeUser fiber.Router) {
			routeUser.Get("/get", func(c *fiber.Ctx) error {
				return c.SendString("result GET /v1/user/get")
			})

			routeUser.Post("/post", func(c *fiber.Ctx) error {
				return c.SendString("result POST /v1/user/post")
			})
		})
}

func subRoute2(app *fiber.App) {
	v2 := app.Group("/v2")
	v2.Route("/user",
		func(routeUser fiber.Router) {
			routeUser.Get("/get", func(c *fiber.Ctx) error {
				return c.SendString("result GET /v2/user/get")
			})

			routeUser.Post("/post", func(c *fiber.Ctx) error {
				return c.SendString("result POST /v2/user/post")
			})
		})
}
