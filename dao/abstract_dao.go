package dao

import (
	"go-oauth/model"
	"gorm.io/gorm"
)

type abstractDao struct {
	interfaceDao
}

func (ad abstractDao) Insert(tx *gorm.DB, repo interface{}) (errMdl model.ErrorModel) {

	err := tx.Create(repo).Error
	if err != nil {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}

func (ad abstractDao) Update(tx *gorm.DB, repo interface{}) (errMdl model.ErrorModel) {
	err := tx.Save(repo).Error
	if err != nil {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}

func (ad abstractDao) Delete(tx *gorm.DB, repo interface{}) (errMdl model.ErrorModel) {
	err := tx.Delete(repo).Error
	if err != nil {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}
