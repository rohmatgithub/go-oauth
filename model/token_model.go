package model

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-oauth/config"
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

func (input JWTToken) ParsingJwtToken(jwtTokenStr string, key string) (result PayloadJWTToken, errMdl ErrorModel) {
	token, err := jwt.ParseWithClaims(jwtTokenStr, &PayloadJWTToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			errMdl = GenerateExpiredTokenError()
			return
		}
		errMdl = GenerateUnknownError(err)
		return
	}

	result = *token.Claims.(*PayloadJWTToken)
	return
}
