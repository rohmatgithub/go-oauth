package model

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-oauth/config"
	"go-oauth/constanta"
	"time"
)

type PayloadJWTToken struct {
	AuthID int64  `json:"auth_id"`
	Scope  string `json:"scope"`
	Locale string `json:"locale"`
	jwt.RegisteredClaims
}

type PayloadTokenInternal struct {
	Scope    string `json:"scope"`
	Locale   string `json:"locale"`
	ClientID string `json:"client_id"`
	Valid    bool   `json:"valid"`
	jwt.RegisteredClaims
}

type JWTToken struct {
}

func (input JWTToken) GenerateToken(payload jwt.Claims) (string, ErrorModel) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, payload)
	token, err := jwtToken.SignedString([]byte(config.ApplicationConfiguration.GetJwtConfig().TokenKey))
	if err != nil {
		return "", GenerateUnknownError(err)
	}
	return token, ErrorModel{}
}

type ValueRedis struct {
	Scope map[string]map[string][]string `json:"scp"`
}

func (input JWTToken) ParsingJwtTokenInternal(jwtTokenStr string) (result PayloadTokenInternal, errMdl ErrorModel) {
	token, err := jwt.ParseWithClaims(jwtTokenStr, &PayloadTokenInternal{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ApplicationConfiguration.GetJwtConfig().TokenKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			errMdl = GenerateExpiredTokenError()
			return
		}
		errMdl = GenerateUnknownError(err)
		return
	}

	result = *token.Claims.(*PayloadTokenInternal)
	return
}

func (input JWTToken) ParsingJwtToken(jwtTokenStr string) (result PayloadJWTToken, errMdl ErrorModel) {
	token, err := jwt.ParseWithClaims(jwtTokenStr, &PayloadJWTToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ApplicationConfiguration.GetJwtConfig().TokenKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			errMdl = GenerateExpiredTokenError()
			return
		}
		errMdl = GenerateUnauthorizedClientError()
		return
	}

	result = *token.Claims.(*PayloadJWTToken)
	return
}

func GetTokenInternal() (string, ErrorModel) {
	expJwtCode := time.Now().Add(constanta.ExpiredJWTCodeConstanta)
	jwtToken, errMdl := JWTToken{}.GenerateToken(
		PayloadTokenInternal{
			ClientID: "",
			Scope:    "",
			Locale:   "en-US",
			Valid:    true,
			RegisteredClaims: jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(expJwtCode),
				Issuer:    "auth",
			},
		})
	if errMdl.Error != nil {
		return "", errMdl
	}

	return jwtToken, ErrorModel{}
}
