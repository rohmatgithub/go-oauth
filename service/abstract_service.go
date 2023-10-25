package service

import (
	"github.com/gofiber/fiber/v2"
	"go-oauth/common"
	"go-oauth/model"
	"reflect"
	"strconv"
)

type AbstractService struct {
	request      *fiber.Ctx
	contextModel *common.ContextModel
	ErrMdl       model.ErrorModel
	value        interface{}
}

func (as *AbstractService) Initialize(request *fiber.Ctx, contextModel *common.ContextModel) *AbstractService {
	as.request = request
	as.contextModel = contextModel
	return as
}

func (as *AbstractService) ReadIDParam(v interface{}) *AbstractService {
	id, _ := strconv.Atoi(as.request.Params(":ID"))
	reflect.ValueOf(v).Elem().SetInt(int64(id))

	if id < 1 {
		as.ErrMdl = model.GenerateUnknownDataError("id")
		return as
	}

	return as
}

func (as *AbstractService) ReadBodyAndValidate(v interface{}) *AbstractService {
	//var stringBody string

	if err := as.request.BodyParser(v); err != nil {
		as.ErrMdl = model.GenerateInvalidRequestError(err)
		return as
	}
	as.value = v
	return as
}
