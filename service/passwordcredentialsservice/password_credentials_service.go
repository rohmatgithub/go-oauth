package passwordcredentialsservice

import (
	"go-oauth/service"
)

var PasswordCredentialsService = passwordCredentialsService{}

type passwordCredentialsService struct {
	service.AbstractService
}
