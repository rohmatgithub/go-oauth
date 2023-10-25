package passwordcredentialsservice

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-oauth/common"
	"go-oauth/dto/in"
	"go-oauth/dto/out"
	"go-oauth/model"
)

func (pcs passwordCredentialsService) VerifyService(c *fiber.Ctx, contextModel *common.ContextModel) (payload out.Payload, errMdl model.ErrorModel) {
	var dtoIn in.PasswordCredentialsIn
	errMdl = pcs.Initialize(c, contextModel).ReadBodyAndValidate(&dtoIn).ErrMdl
	if errMdl.Error != nil {
		return
	}

	errMdl = dtoIn.ValidateLogin()
	if errMdl.Error != nil {
		return
	}
	fmt.Printf("username : %s, password : %s", dtoIn.Username, dtoIn.Password)
	return
}
