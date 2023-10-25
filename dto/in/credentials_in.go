package in

import (
	"go-oauth/constanta"
	"go-oauth/model"
)

type PasswordCredentialsIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *PasswordCredentialsIn) ValidateLogin() (errMdl model.ErrorModel) {
	if req.Username == "" {
		return model.GenerateEmptyFieldError(constanta.Username)
	}

	if req.Password == "" {
		return model.GenerateEmptyFieldError(constanta.Password)
	}
	return model.ErrorModel{}
}
