package model

// ==================== ERROR DTO ===================

func GenerateEmptyFieldError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-DTO-001", fileName, funcName, errorParam)
}

func GenerateFormatFieldError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-DTO-002", fileName, funcName, errorParam)
}

func GenerateFieldFormatWithRuleError(fileName string, funcName string, ruleName string, fieldName string, additionalInfo string) ErrorModel {
	errorParam := make([]ErrorParameter, 3)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	errorParam[1].ErrorParameterKey = "RuleName"
	errorParam[1].ErrorParameterValue = ruleName
	errorParam[2].ErrorParameterKey = "Other"
	errorParam[2].ErrorParameterValue = additionalInfo
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-DTO-003", fileName, funcName, errorParam)
}

func GenerateInvalidRequestError(fileName string, funcName string, causedBy error) ErrorModel {
	return GenerateErrorModel(400, "E-4-AUT-DTO-004", fileName, funcName, causedBy)
}

// ====================== ERROR SERVICE  ===================
func GenerateUnauthorizedClientError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModel(401, "E-1-AUT-SRV-001", fileName, funcName, nil)
}

func GenerateVerifyPasswordInvalidError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModel(401, "E-1-AUT-SRV-002", fileName, funcName, nil)
}

func GenerateExpiredTokenError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModel(401, "E-1-AUT-SRV-003", fileName, funcName, nil)
}

func GenerateCannotUsedError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-001", fileName, funcName, errorParam)
}

func GenerateUnknownDataError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-002", fileName, funcName, errorParam)
}

func GenerateDuplicateDataError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-003", fileName, funcName, errorParam)
}

func GenerateNotAccessError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-004", fileName, funcName, errorParam)
}

func GenerateFieldInvalid(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-005", fileName, funcName, errorParam)
}

func GenerateNotChangedDataError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-006", fileName, funcName, errorParam)
}

func GenerateNotDeleteDataError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-007", fileName, funcName, errorParam)
}

func GenerateUnknownError(fileName string, funcName string, causedBy error) ErrorModel {
	return GenerateErrorModel(500, "E-5-AUT-SRV-001", fileName, funcName, causedBy)
}

func GenerateInternalDBServerError(fileName string, funcName string, causedBy error) ErrorModel {
	return GenerateErrorModel(500, "E-5-AUT-DBS-001", fileName, funcName, causedBy)
}
