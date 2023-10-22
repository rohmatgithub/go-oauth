package model

import "errors"

type ErrorModel struct {
	Code                  int
	Error                 error
	FileName              string
	FuncName              string
	CausedBy              error
	ErrorParameter        []ErrorParameter
	AdditionalInformation interface{}
}

type ErrorParameter struct {
	ErrorParameterKey   string
	ErrorParameterValue string
}

func GenerateErrorModel(code int, err string, fileName string, funcName string, causedBy error) ErrorModel {
	var errModel ErrorModel
	errModel.Code = code
	errModel.Error = errors.New(err)
	errModel.FileName = fileName
	errModel.FuncName = funcName
	errModel.CausedBy = causedBy
	return errModel
}

func GenerateErrorModelWithErrorParam(code int, err string, fileName string, funcName string, errorParam []ErrorParameter) ErrorModel {
	var errModel ErrorModel
	errModel.Code = code
	errModel.Error = errors.New(err)
	errModel.FileName = fileName
	errModel.FuncName = funcName
	errModel.ErrorParameter = errorParam
	return errModel
}
