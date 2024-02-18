package endpoint

import (
	"context"
	"encoding/json"
	"go-oauth/common"
	"go-oauth/config"
	"go-oauth/constanta"
	"go-oauth/dao"
	"go-oauth/dto/out"
	"go-oauth/model"
	"go-oauth/repository"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

type AbstractEndpoint struct {
}

func (ae AbstractEndpoint) serve(c *fiber.Ctx,
	validateFunc func(c *fiber.Ctx, contextModel *common.ContextModel) model.ErrorModel,
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
		if r := recover(); r != nil {
			contextModel.LoggerModel.Message = string(debug.Stack())
			generateEResponseError(c, &contextModel, &payload, model.GenerateUnknownError(nil))
		}
		response.Payload = payload

		adaptor.CopyContextToFiberContext(context.WithValue(c.Context(), constanta.ApplicationContextConstanta, &contextModel.LoggerModel), c.Context())
		err = c.JSON(response)
	}()
	// validate
	errMdl := validateFunc(c, &contextModel)
	if errMdl.Error != nil {
		generateEResponseError(c, &contextModel, &payload, errMdl)
		return
	}
	payload, errMdl = runFunc(c, &contextModel)
	if errMdl.Error != nil {
		generateEResponseError(c, &contextModel, &payload, errMdl)
	} else {
		payload.Status.Success = true
		payload.Status.Code = "OK"
	}
	return
}

func generateEResponseError(c *fiber.Ctx, ctxModel *common.ContextModel, payload *out.Payload, errMdl model.ErrorModel) {
	ctxModel.LoggerModel.Code = errMdl.Error.Error()
	ctxModel.LoggerModel.Class = errMdl.Line
	if errMdl.CausedBy != nil {
		ctxModel.LoggerModel.Message = errMdl.CausedBy.Error()
	}
	// write failed
	c.Status(errMdl.Code)
	payload.Status.Success = false
	payload.Status.Code = errMdl.Error.Error()
	payload.Status.Message = common.GenerateI18NErrorMessage(errMdl, ctxModel.AuthAccessTokenModel.Locale)

}

func (ae AbstractEndpoint) EndpointWhiteList(c *fiber.Ctx, runFunc func(*fiber.Ctx, *common.ContextModel) (out.Payload, model.ErrorModel)) error {

	return ae.serve(c, func(*fiber.Ctx, *common.ContextModel) model.ErrorModel {
		return model.ErrorModel{}
	}, runFunc)
}

func (ae AbstractEndpoint) EndpointClientCredentials(c *fiber.Ctx, runFunc func(*fiber.Ctx, *common.ContextModel) (out.Payload, model.ErrorModel)) error {
	// validate client_id

	validateFunc := func(cx *fiber.Ctx, contextModel *common.ContextModel) (errMdl model.ErrorModel) {
		// cek token expired
		tokenStr := cx.Get(constanta.TokenHeaderNameConstanta)
		// destresource := cx.Get(constanta.HeaderDestResourceKey)
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
		// // validate scope token
		// scope := valueModel.Scope[destresource]
		// log.Debug(fmt.Sprintf("dest-resource : [%s], value-scope : [%v]", destresource, scope))
		// if scope == nil {
		// 	errMdl = model.GenerateUnauthorizedClientError()
		// 	return
		// }
		return
	}

	return ae.serve(c, validateFunc, runFunc)
}

func validatePermissionUser(c *fiber.Ctx, contextModel *common.ContextModel) (errMdl model.ErrorModel) {
	tokenStr := c.Get(constanta.TokenHeaderNameConstanta)

	if tokenStr == "" {
		return model.GenerateUnauthorizedClientError()
	}
	// cek token expired
	parsedToken, errMdl := model.JWTToken{}.ParsingJwtToken(tokenStr)
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

	if strings.TrimSpace(valueToken) == "" {
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

	contextModel.AuthAccessTokenModel.ResourceUserID = parsedToken.AuthID
	contextModel.AuthAccessTokenModel.CompanyID = valueModel.CompanyID

	return
}
func (ae AbstractEndpoint) EndpointJwtToken(c *fiber.Ctx, runFunc func(*fiber.Ctx, *common.ContextModel) (out.Payload, model.ErrorModel)) error {

	return ae.serve(c, validatePermissionUser, runFunc)
}

func MiddlewareOtherService(c *fiber.Ctx) (err error) {
	var (
		response     out.StandardResponse
		payload      out.Payload
		contextModel common.ContextModel
		errMdl       model.ErrorModel
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
		if r := recover(); r != nil {
			contextModel.LoggerModel.Message = string(debug.Stack())
			errMdl = model.GenerateUnknownError(nil)
			generateEResponseError(c, &contextModel, &payload, errMdl)
		}
		if errMdl.Error != nil {
			response.Payload = payload

			adaptor.CopyContextToFiberContext(context.WithValue(c.Context(), constanta.ApplicationContextConstanta, &contextModel.LoggerModel), c.Context())
			err = c.JSON(response)
		}

	}()
	errMdl = validatePermissionUser(c, &contextModel)
	if errMdl.Error != nil {
		generateEResponseError(c, &contextModel, &payload, errMdl)
		return
	}

	tokenInternal, errMdl := model.GetTokenInternal(contextModel.AuthAccessTokenModel.ResourceUserID,
		contextModel.AuthAccessTokenModel.CompanyID)
	if errMdl.Error != nil {
		generateEResponseError(c, &contextModel, &payload, errMdl)
		return
	}

	c.Locals(constanta.TokenInternalHeaderName, tokenInternal)

	return c.Next()
}
