package go_fiber_example

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http/httptest"
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
	routerTest(app)
	err := app.Listen(":3000")
	if err != nil {
		log.Panic(err)
	}
}

func startTest(app *fiber.App) {

	app.Get("/getcounthandler", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("%d", app.HandlersCount()))
	})
	app.Get("/hello/:value?", func(c *fiber.Ctx) error {
		return c.SendString("hello " + c.Params("value"))
	}).Name("hello")

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
	}).Name("test-error")

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

	app.Get("/getstack/:routeName?", func(c *fiber.Ctx) error {
		var data []byte
		routeName := c.Params("routeName")
		if routeName != "" {
			data, _ = json.MarshalIndent(app.GetRoute(routeName), "", " ")
		} else {
			data, _ = json.MarshalIndent(app.Stack(), "", "  ")
		}
		//fmt.Println(string(data))
		return c.SendString(string(data))
	})

	//// Match request starting with /api
	//app.Use("/api", func(c *fiber.Ctx) error {
	//	fmt.Println("masuk middleware")
	//	return c.Next()
	//})
}

func routerTest(app *fiber.App) {
	// Create route with GET method for test:
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c.BaseURL())              // => http://google.com
		fmt.Println(c.Get("X-Custom-Header")) // => hi

		return c.SendString("hello, World!")
	})

	// http.Request
	req := httptest.NewRequest("GET", "http://localhost:3000", nil)
	req.Header.Set("X-Custom-Header", "hi")

	// http.Response
	resp, _ := app.Test(req)

	// Do something with results:
	if resp.StatusCode == fiber.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("-->>" + string(body)) // => Hello, World!
	}
}

func subRoute1(app *fiber.App) {
	v1 := app.Group("/v1")

	v1.Use(func(c *fiber.Ctx) error {
		content := c.Accepts("text/plain")

		fmt.Println(content)
		if content == "" {
			return errors.New("error")
		}
		return c.Next()
	})
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
