package dao

import (
	"go-oauth/common"
	"go-oauth/model"
	"go-oauth/repository"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

var AccessPermissionsDao = accessPermissionsDao{}

type accessPermissionsDao struct {
	abstractDao
}

func (a accessPermissionsDao) InsertData(tx *gorm.DB, repo repository.AccessPermission) (errMdl model.ErrorModel) {
	query := "INSERT INTO access_permissions (user_id, company_id, branch_id, is_admin) " +
		" VALUES ($1, $2, $3, $4) "
	param := []interface{}{repo.UserID, repo.CompanyID, pq.Array(repo.BranchID), repo.IsAdmin}
	err := tx.Exec(query, param...)
	if err.Error != nil {
		errMdl = model.GenerateInternalDBServerError(err.Error)
	}

	return
}

func (a accessPermissionsDao) GetAccessPermissions(userID int64) (result repository.AccessPermission, errMdl model.ErrorModel) {
	db := common.GormDB
	var pqArray pq.Int64Array
	query := "SELECT user_id, company_id, branch_id, is_admin FROM access_permissions WHERE user_id = $1"
	err := db.Raw(query, userID).Row().Scan(&result.UserID, &result.CompanyID, &pqArray, &result.IsAdmin)
	if err != nil {
		errMdl = model.GenerateInternalDBServerError(err)
	}

	result.BranchID = pqArray
	return
}
