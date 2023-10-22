package passwordcredentialsservice

import (
	"github.com/gofiber/fiber/v2"
	"go-oauth/model"
)

func (pcs passwordCredentialsService) VerifyService(c *fiber.Ctx) (errMdl model.ErrorModel) {
	pcs.FileName = "verify_service.go"
	pcs.FuncName = "VerifyService"

	return errMdl
}
