package router

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"go-oauth/model"
	"net/http"
)

func masterDataRouter(app fiber.Router) {
	app.Get("/*", func(c *fiber.Ctx) error {
		req, err := http.NewRequest("GET", "http://localhost:9092"+c.OriginalURL(), nil)
		if err != nil {
			return err
		}

		token, errMdl := model.GetTokenInternal()
		if errMdl.Error != nil {
			return errMdl.Error
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", token)

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
		req, err := http.NewRequest("POST", "http://localhost:9092"+c.OriginalURL(), bytes.NewReader(c.Body()))
		if err != nil {
			return err
		}

		token, errMdl := model.GetTokenInternal()
		if errMdl.Error != nil {
			return errMdl.Error
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", c.Get("Content-Type"))
		req.Header.Set("Authorization", token)

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
		req, err := http.NewRequest("PUT", "http://localhost:9092"+c.OriginalURL(), bytes.NewReader(c.Body()))
		if err != nil {
			return err
		}

		token, errMdl := model.GetTokenInternal()
		if errMdl.Error != nil {
			return errMdl.Error
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", c.Get("Content-Type"))
		req.Header.Set("Authorization", token)

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
