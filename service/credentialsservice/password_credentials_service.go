package credentialsservice

import (
	context2 "context"
	"database/sql"
	"encoding/json"
	"go-oauth/common"
	"go-oauth/config"
	"go-oauth/constanta"
	"go-oauth/dao"
	"go-oauth/dto/in"
	"go-oauth/dto/out"
	"go-oauth/model"
	"go-oauth/repository"
	"go-oauth/util"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
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

	token, data, errMdl := validateUsername(dtoIn, "auth")
	if errMdl.Error != nil {
		return
	}

	payload.Data = data
	c.Set(constanta.TokenHeaderNameConstanta, token)
	return
}

func validateUsername(dtoIn in.PasswordCredentialsIn, resourceID string) (token string, data interface{}, errMdl model.ErrorModel) {
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
		CompanyID: userDB.CompanyID.Int64,
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

	// get company from master
	// Create an HTTP client
	client := &http.Client{}

	tokenInternal, errMdl := model.GetTokenInternal(userDB.ID.Int64, userDB.CompanyID.Int64)
	if errMdl.Error != nil {
		return
	}
	uri := config.ApplicationConfiguration.GetUriResouce().MasterData + "/v1/master/company/" + strconv.Itoa(int(userDB.CompanyID.Int64))
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Error("Error creating request:", err)
		errMdl = model.GenerateUnknownError(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(constanta.TokenHeaderNameConstanta, tokenInternal)

	// Make the request
	response, err := client.Do(req)
	if err != nil {
		log.Error("Error making request:", err)
		errMdl = model.GenerateUnknownError(err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error("Error reading response:", err)
		errMdl = model.GenerateUnknownError(err)
		return
	}

	var company out.CompanyResponse
	err = json.Unmarshal(body, &company)
	if err != nil {
		log.Error("Error unmarshalling JSON:", err)
		errMdl = model.GenerateUnknownError(err)
		return
	}
	return jwtToken, company.Payload.Data, model.ErrorModel{}
}
