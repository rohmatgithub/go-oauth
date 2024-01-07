package credentialsservice

import (
	context2 "context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"go-oauth/common"
	"go-oauth/constanta"
	"go-oauth/dao"
	"go-oauth/dto/in"
	"go-oauth/dto/out"
	"go-oauth/model"
	"go-oauth/repository"
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

	if userDB.ID.Int64 == 0 {
		errMdl = model.GenerateUnauthorizedClientError()
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
	tx := common.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil || errMdl.Error != nil {
			tx.Rollback()
		} else {
			// save to redis
			err := tx.Commit().Error
			if err != nil {
				errMdl = model.GenerateInternalDBServerError(err)
				return
			}
		}
	}()

	valueRedis := model.ValueRedis{
		Scope: nil,
	}
	valueToken := util.JsonToString(valueRedis)
	authTokenRepo := repository.AuthToken{
		Tokens:      sql.NullString{String: jwtToken, Valid: true},
		ExpiredTime: sql.NullInt64{Int64: expJwtCode.Unix(), Valid: true},
		ValueToken:  sql.NullString{String: valueToken, Valid: true},
		CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	}
	errMdl = dao.AuthTokenDao.Insert(tx, &authTokenRepo)
	if errMdl.Error != nil {
		return
	}

	// insert to redis
	context := context2.Background()
	redisStatus := common.RedisClient.Set(context, jwtToken, valueToken, expJwtCode.Sub(time.Now()))
	if redisStatus != nil && redisStatus.Err() != nil {
		log.Error(redisStatus.Err())

	}

	return jwtToken, model.ErrorModel{}
}
