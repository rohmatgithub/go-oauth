package dao

import (
	"database/sql"
	"errors"
	"go-oauth/common"
	"go-oauth/model"
	"go-oauth/repository"
)

var AuthTokenDao authTokenDao

type authTokenDao struct {
	abstractDao
}

func (at authTokenDao) GetByToken(token string) (result repository.AuthToken, errMdl model.ErrorModel) {
	db := common.GormDB
	err := db.Where("tokens = ?", token).Find(&result).Error
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		errMdl = model.GenerateInternalDBServerError(err)
	}

	return
}
