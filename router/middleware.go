package router

import (
	"context"
	"fmt"
	"go-oauth/common"
	"go-oauth/config"
	"go-oauth/constanta"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func middleware(c *fiber.Ctx) error {
	logModel := &common.LoggerModel{
		Pid:         strconv.Itoa(os.Getpid()),
		RequestID:   c.Locals("requestid").(string),
		Resource:    "",
		Application: config.ApplicationConfiguration.GetServerConfig().ResourceID,
		Version:     config.ApplicationConfiguration.GetServerConfig().Version,
		ByteIn:      len(c.Body()),
		//Path:        c.BaseURL(),
	}
	logger := context.WithValue(c.Context(), constanta.ApplicationContextConstanta, logModel)
	adaptor.CopyContextToFiberContext(logger, c.Context())

	err := c.Next()
	if err != nil {
		return err
	}

	c.Set("Access-Control-Expose-Headers", "Authorization")
	//c.Set("Access-Control-Expose-Headers", "X-Request-Id")
	logModel = c.Context().Value(constanta.ApplicationContextConstanta).(*common.LoggerModel)
	logModel.Status = c.Response().StatusCode()
	logModel.Path = c.OriginalURL()
	log.Info(common.GenerateLogModel(*logModel))
	return err
}

func NotFoundHandler(c *fiber.Ctx) error {
	if strings.HasPrefix(c.OriginalURL(), "/v1/master") {
		fmt.Println(c.OriginalURL())
		return c.Redirect("localhost:9092/v1/master/company?page=1&order_by=name DESC")
	}

	// Customize the response for the 404 error
	return c.Status(fiber.StatusNotFound).SendString("404 Not Found")
}

func customErrorHandler(c *fiber.Ctx, err error) {
	// Handle the error here
	fmt.Printf("Error: %v\n", err)

	// Return a custom error response
	c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Something went wrong",
	})
}

func httpRespToFiberCtx(resp *http.Response, c *fiber.Ctx) {
	defer resp.Body.Close()
	err := c.JSON(resp.Body)
	if err != nil {
		log.Error(err)
		return
	}

	c.Status(resp.StatusCode)
	// Iterate over the headers and set each one on the Ctx
	for key, values := range resp.Header {
		for _, value := range values {
			c.Append(key, value)
		}
	}

	// Get the body writer
	bodyWriter := c.Response().BodyWriter()

	// Copy the response body
	_, err = io.Copy(bodyWriter, resp.Body)
	if err != nil {
		log.Error(err)
		return
	}
}

//func middlewareOtherService(c *fiber.Ctx) error {
//	// authorize
//
//	tokenStr := c.Get(constanta.TokenHeaderNameConstanta)
//	fmt.Println(tokenStr)
//	return c.Next()
//}
