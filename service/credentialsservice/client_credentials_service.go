package credentialsservice

import (
	context2 "context"
	"database/sql"
	"encoding/json"
	"go-oauth/common"
	"go-oauth/constanta"
	"go-oauth/dao"
	"go-oauth/dto/out"
	"go-oauth/model"
	"go-oauth/repository"
	"go-oauth/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

func (cs credentialsService) ClientCredentialsService(c *fiber.Ctx, _ *common.ContextModel) (payload out.Payload, errMdl model.ErrorModel) {

	clientID := c.Get(constanta.HeaderClientIdKey)
	clientSecret := c.Get(constanta.HeaderClientSecretKey)
	destResource := c.Get(constanta.HeaderDestResourceKey)

	repo := repository.AuthClient{
		ClientID: sql.NullString{String: clientID, Valid: true},
	}

	//var repo []repository.AuthClient
	errMdl = dao.AuthClientDao.ViewDetail(&repo)
	if errMdl.Error != nil {
		return
	}

	if repo.ClientID.String == "" || repo.ClientSecret.String != clientSecret {
		errMdl = model.GenerateUnauthorizedClientError()
		return
	}

	scope := make(map[string]map[string][]string)
	var isExistDest bool
	for _, resource := range repo.ClientResource {
		var authorities map[string][]string
		_ = json.Unmarshal([]byte(resource.Authorities.String), &authorities)
		scope[resource.ResourceID.String] = authorities
		if resource.ResourceID.String == destResource {
			isExistDest = true
		}
	}

	if !isExistDest {
		errMdl = model.GenerateUnauthorizedClientError()
		return
	}

	expJwtCode := time.Now().Add(constanta.ExpiredJWTCodeConstanta)
	jwtToken, errMdl := model.JWTToken{}.GenerateToken(
		model.PayloadTokenInternal{
			Locale: "en-US",
			Valid:  true,
			RegisteredClaims: jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(expJwtCode),
				Issuer:    "auth",
			},
		})
	if errMdl.Error != nil {
		return
	}

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
		CompanyID: 0,
		BranchID:  0,
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
	c.Set(constanta.TokenHeaderNameConstanta, jwtToken)
	return
}
