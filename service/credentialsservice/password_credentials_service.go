package credentialsservice

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-oauth/common"
	"go-oauth/constanta"
	"go-oauth/dao"
	"go-oauth/dto/in"
	"go-oauth/dto/out"
	"go-oauth/model"
	"go-oauth/util"
	"time"
)

func (cs credentialsService) VerifyService(c *fiber.Ctx, contextModel *common.ContextModel) (payload out.Payload, errMdl model.ErrorModel) {
	var dtoIn in.PasswordCredentialsIn
	errMdl = cs.Initialize(c, contextModel).ReadBodyAndValidate(&dtoIn).ErrMdl
	if errMdl.Error != nil {
		return
	}

	errMdl = dtoIn.ValidateLogin()
	if errMdl.Error != nil {
		return
	}

	token, errMdl := validateUsername(dtoIn, "auth")
	if errMdl.Error != nil {
		return
	}

	c.Set(constanta.TokenHeaderNameConstanta, token)
	return
}

func validateUsername(dtoIn in.PasswordCredentialsIn, resourceID string) (token string, errMdl model.ErrorModel) {
	// get user by username in db
	userDB, errMdl := dao.UserDao.GetDataByUsername(dtoIn.Username, resourceID)
	if errMdl.Error != nil {
		return
	}

	// validate username & password
	isMatch := util.CheckIsPasswordMatch(dtoIn.Password, userDB.Password.String, userDB.Salt.String)
	if !isMatch {
		errMdl = model.GenerateVerifyPasswordInvalidError()
		return
	}

	expJwtCode := time.Now().Add(constanta.ExpiredJWTCodeConstanta)
	jwtToken, errMdl := model.JWTToken{}.GenerateToken(
		model.PayloadJWTToken{
			AuthID: userDB.ID.Int64,
			Scope:  "",
			Locale: userDB.Locale.String,
			RegisteredClaims: jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(expJwtCode),
				Issuer:    "auth",
			},
		})
	if errMdl.Error != nil {
		return
	}

	// insert token to db
	// insert token to redis

	return jwtToken, model.ErrorModel{}
}
