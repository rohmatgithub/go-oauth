package model

// ==================== ERROR DTO ===================

func GenerateEmptyFieldError(fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-DTO-001", errorParam)
}

func GenerateFormatFieldError(fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-DTO-002", errorParam)
}

func GenerateFieldFormatWithRuleError(ruleName string, fieldName string, additionalInfo string) ErrorModel {
	errorParam := make([]ErrorParameter, 3)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	errorParam[1].ErrorParameterKey = "RuleName"
	errorParam[1].ErrorParameterValue = ruleName
	errorParam[2].ErrorParameterKey = "Other"
	errorParam[2].ErrorParameterValue = additionalInfo
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-DTO-003", errorParam)
}

func GenerateInvalidRequestError(causedBy error) ErrorModel {
	return GenerateErrorModel(400, "E-4-AUT-DTO-004", causedBy)
}

// ====================== ERROR SERVICE  ===================

func GenerateUnauthorizedClientError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModel(401, "E-1-AUT-SRV-001", nil)
}

func GenerateVerifyPasswordInvalidError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModel(401, "E-1-AUT-SRV-002", nil)
}

func GenerateExpiredTokenError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModel(401, "E-1-AUT-SRV-003", nil)
}

func GenerateCannotUsedError(fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-001", errorParam)
}

func GenerateUnknownDataError(fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-002", errorParam)
}

func GenerateDuplicateDataError(fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-003", errorParam)
}

func GenerateNotAccessError(fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-004", errorParam)
}

func GenerateFieldInvalid(fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-005", errorParam)
}

func GenerateNotChangedDataError(fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-006", errorParam)
}

func GenerateNotDeleteDataError(fieldName string) ErrorModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-007", errorParam)
}

func GenerateUnknownError(causedBy error) ErrorModel {
	return GenerateErrorModel(500, "E-5-AUT-SRV-001", causedBy)
}

func GenerateInternalDBServerError(causedBy error) ErrorModel {
	return GenerateErrorModel(500, "E-5-AUT-DBS-001", causedBy)
}
