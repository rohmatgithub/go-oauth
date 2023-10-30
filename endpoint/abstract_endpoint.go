package endpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"go-oauth/common"
	"go-oauth/config"
	"go-oauth/constanta"
	"go-oauth/dao"
	"go-oauth/dto/out"
	"go-oauth/model"
	"go-oauth/repository"
	"strings"
	"time"
)

type AbstractEndpoint struct {
}

func (ae AbstractEndpoint) EndpointWhiteList(c *fiber.Ctx, runFunc func(*fiber.Ctx, *common.ContextModel) (out.Payload, model.ErrorModel)) error {

	return ae.serve(c, func(*common.ContextModel) model.ErrorModel {
		return model.ErrorModel{}
	}, runFunc)
}

func (ae AbstractEndpoint) EndpointClientCredentials(c *fiber.Ctx, runFunc func(*fiber.Ctx, *common.ContextModel) (out.Payload, model.ErrorModel)) error {
	// validate client_id
	tokenStr := c.Get(constanta.TokenHeaderNameConstanta)
	destresource := c.Get(constanta.HeaderDestResourceKey)

	validateFunc := func(contextModel *common.ContextModel) (errMdl model.ErrorModel) {
		// cek token expired
		_, errMdl = model.JWTToken{}.ParsingJwtTokenInternal(tokenStr)
		if errMdl.Error != nil {
			return
		}

		ctx := context.Background()
		// get value token from storage
		redis := common.RedisClient.Get(ctx, tokenStr)
		var valueToken string
		if redis != nil {
			valueToken = redis.Val()
		}

		if valueToken == strings.TrimSpace(valueToken) {
			// get to db
			var authTokenDB repository.AuthToken
			authTokenDB, errMdl = dao.AuthTokenDao.GetByToken(tokenStr)
			if errMdl.Error != nil {
				return
			}

			valueToken = authTokenDB.ValueToken.String
		}
		var valueModel model.ValueRedis
		err := json.Unmarshal([]byte(valueToken), &valueModel)
		if err != nil {
			log.Error(err)
			errMdl = model.GenerateUnauthorizedClientError()
			return
		}
		// validate scope token
		scope := valueModel.Scope[destresource]
		log.Debug(fmt.Sprintf("dest-resource : [%s], value-scope : [%v]", destresource, scope))
		if scope == nil {
			errMdl = model.GenerateUnauthorizedClientError()
			return
		}
		return
	}

	return ae.serve(c, validateFunc, runFunc)
}

func (ae AbstractEndpoint) serve(c *fiber.Ctx,
	validateFunc func(contextModel *common.ContextModel) model.ErrorModel,
	runFunc func(*fiber.Ctx, *common.ContextModel) (out.Payload, model.ErrorModel)) (err error) {
	var (
		response     out.StandardResponse
		payload      out.Payload
		contextModel common.ContextModel
	)

	requestID := c.Locals("requestid").(string)
	logModel := c.Context().Value(constanta.ApplicationContextConstanta).(*common.LoggerModel)

	contextModel.LoggerModel = *logModel
	response.Header = out.HeaderResponse{
		RequestID: requestID,
		Version:   config.ApplicationConfiguration.GetServerConfig().Version,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	defer func() {
		response.Payload = payload

		adaptor.CopyContextToFiberContext(logModel, c.Context())
		err = c.JSON(response)
	}()
	// validate
	errMdl := validateFunc(&contextModel)
	if errMdl.Error != nil {
		generateEResponseError(c, logModel, &payload, errMdl)
		return
	}
	payload, errMdl = runFunc(c, &contextModel)
	if errMdl.Error != nil {
		generateEResponseError(c, logModel, &payload, errMdl)
	} else {
		payload.Status = out.StatusPayload{
			Success: true,
			Code:    "OK",
		}
	}
	return
}

func generateEResponseError(c *fiber.Ctx, logModel *common.LoggerModel, payload *out.Payload, errMdl model.ErrorModel) {
	logModel.Code = errMdl.Error.Error()
	logModel.Class = errMdl.Line
	if errMdl.CausedBy != nil {
		logModel.Message = errMdl.CausedBy.Error()
	}
	// write failed
	c.Status(errMdl.Code)
	payload.Status = out.StatusPayload{
		Success: false,
		Code:    errMdl.Error.Error(),
		Message: "",
	}
}
