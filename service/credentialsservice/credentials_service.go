package credentialsservice

import (
	"go-oauth/service"
)

var CredentialsService = credentialsService{}

type credentialsService struct {
	service.AbstractService
}
