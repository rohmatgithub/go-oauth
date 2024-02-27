package router

import (
	"bytes"
	"go-oauth/config"
	"go-oauth/constanta"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func reportRouter(app fiber.Router) {
	app.Get("/*", func(c *fiber.Ctx) error {

		req, err := http.NewRequest("GET", config.ApplicationConfiguration.GetUriResouce().Report+c.OriginalURL(), nil)
		if err != nil {
			return err
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Set(constanta.TokenHeaderNameConstanta, c.Locals(constanta.TokenInternalHeaderName).(string))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp != nil {
			httpRespToFiberCtx(resp, c)
		}
		return nil
	})

	app.Post("/*", func(c *fiber.Ctx) error {
		req, err := http.NewRequest("POST", config.ApplicationConfiguration.GetUriResouce().Report+c.OriginalURL(), bytes.NewReader(c.Body()))
		if err != nil {
			return err
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", c.Get("Content-Type"))
		req.Header.Set(constanta.TokenHeaderNameConstanta, c.Locals(constanta.TokenInternalHeaderName).(string))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp != nil {
			httpRespToFiberCtx(resp, c)
		}
		return nil
	})

	app.Put("/*", func(c *fiber.Ctx) error {
		req, err := http.NewRequest("PUT", config.ApplicationConfiguration.GetUriResouce().Report+c.OriginalURL(), bytes.NewReader(c.Body()))
		if err != nil {
			return err
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", c.Get("Content-Type"))
		req.Header.Set(constanta.TokenHeaderNameConstanta, c.Locals(constanta.TokenInternalHeaderName).(string))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp != nil {
			httpRespToFiberCtx(resp, c)
		}
		return nil
	})
}
