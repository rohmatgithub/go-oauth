package dao

import (
	"go-oauth/model"
	"gorm.io/gorm"
)

type interfaceDao interface {
	ViewDetail(interface{}) model.ErrorModel
	List() (interface{}, model.ErrorModel)
	Insert(*gorm.DB, interface{}) model.ErrorModel
	Update(*gorm.DB, interface{}) model.ErrorModel
	Delete(*gorm.DB, interface{}) model.ErrorModel
}
