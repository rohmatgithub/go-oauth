package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"go-oauth/common"
	"go-oauth/config"
	"go-oauth/constanta"
	"go-oauth/dto/out"
	"go-oauth/model"
	"time"
)

type AbstractEndpoint struct {
}

func (ae AbstractEndpoint) EndpointWhiteList(c *fiber.Ctx, runFunc func(*fiber.Ctx, *common.ContextModel) (out.Payload, model.ErrorModel)) error {

	return ae.serve(c, runFunc)
}

// func (ae AbstractEndpoint) EndpointClientCredentials(c *fiber.Ctx, runFunc func(*fiber.Ctx) model.ErrorModel) error {
//
//		return nil
//	}
func (ae AbstractEndpoint) serve(c *fiber.Ctx, runFunc func(*fiber.Ctx, *common.ContextModel) (out.Payload, model.ErrorModel)) error {
	var response out.StandardResponse

	requestID := c.Locals("requestid").(string)
	logModel := c.Context().Value(constanta.ApplicationContextConstanta).(*common.LoggerModel)

	response.Header = out.HeaderResponse{
		RequestID: requestID,
		Version:   config.ApplicationConfiguration.GetServerConfig().Version,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	payload, errMdl := runFunc(c, nil)
	if errMdl.Error != nil {
		logModel.Code = errMdl.Error.Error()
		logModel.Class = errMdl.Line
		// write failed
		c.Status(errMdl.Code)
		payload.Status = out.StatusPayload{
			Success: false,
			Code:    errMdl.Error.Error(),
			Message: "",
		}
	} else {
		payload.Status = out.StatusPayload{
			Success: true,
			Code:    "OK",
		}
	}
	response.Payload = payload

	adaptor.CopyContextToFiberContext(logModel, c.Context())
	return c.JSON(response)
}
