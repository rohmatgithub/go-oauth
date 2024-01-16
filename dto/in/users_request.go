package in

import (
	"go-oauth/common"
	"go-oauth/dto"
	"go-oauth/model"
)

type UsersRequest struct {
	Username       string  `json:"username" validate:"required,min=8,max=50"`
	Password       string  `json:"password" validate:"required,min=8,max=20"`
	VerifyPassword string  `json:"verify_password" validate:"required,min=8,max=20"`
	Email          string  `json:"email" validate:"required,email"`
	Phone          string  `json:"phone" validate:"required,numeric,min=10,max=16"`
	Locale         string  `json:"locale" validate:"required"`
	IsAdmin        bool    `json:"is_admin"`
	FirstName      string  `json:"first_name" validate:"required,min=3,max=50"`
	LastName       string  `json:"last_name" validate:"required,min=3,max=50"`
	Address1       string  `json:"address_1" validate:"required,min=3,max=200"`
	Address2       string  `json:"address_2" validate:"required,min=3,max=200"`
	CountryID      int64   `json:"countryID" validate:"required"`
	CompanyID      int64   `json:"company_id" validate:"required"`
	BranchID       []int64 `json:"branch_id" validate:"required"`
	dto.AbstractDto
}

func (u *UsersRequest) ValidateInsert(contextModel *common.ContextModel) (resultMap map[string]string, errMdl model.ErrorModel) {
	resultMap = common.Validation.ValidationAll(*u, contextModel)

	if u.Password != u.VerifyPassword && resultMap["verify_password"] == "" {
		resultMap["verify_password"] = "password and verify password not match"
	}
	return
}

func (u *UsersRequest) ValidateUpdate(contextModel *common.ContextModel) (resultMap map[string]string, errMdl model.ErrorModel) {
	resultMap = common.Validation.ValidationAll(*u, contextModel)

	if u.Password != u.VerifyPassword && resultMap["verify_password"] == "" {
		resultMap["verify_password"] = "password and verify password not match"
	}

	errMdl = u.ValidateUpdateGeneral()
	return
}
