package passwordcredentialsservice

import (
	"go-oauth/common"
)

var PasswordCredentialsService = passwordCredentialsService{}

type passwordCredentialsService struct {
	common.AbstractStruct
}
