package model

import (
	"errors"
	"go-oauth/config"
	"go-oauth/constanta"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type PayloadJWTToken struct {
	AuthID int64  `json:"uid"`
	Locale string `json:"locale"`
	jwt.RegisteredClaims
}

type PayloadTokenInternal struct {
	Locale    string `json:"locale"`
	UserID    int64  `json:"uid"`
	CompanyID int64  `json:"cid"`
	BranchID  int64  `json:"bid"`
	Valid     bool   `json:"valid"`
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
	CompanyID int64 `json:"cpid"`
	BranchID  int64 `json:"brid"`
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

func GetTokenInternal(userID, companyID, branchID int64) (string, ErrorModel) {
	expJwtCode := time.Now().Add(constanta.ExpiredJWTCodeConstanta)
	jwtToken, errMdl := JWTToken{}.GenerateToken(
		PayloadTokenInternal{
			Locale:    "en-US",
			UserID:    userID,
			CompanyID: companyID,
			BranchID:  branchID,
			Valid:     true,
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
