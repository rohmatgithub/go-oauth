package dao

import (
	"database/sql"
	"go-oauth/common"
	"go-oauth/model"
	"go-oauth/repository"
)

var UserDao = usersDao{}

type usersDao struct {
	abstractDao
}

func (usersDao) ViewDetail(repo interface{}) (result interface{}, errMdl model.ErrorModel) {

	return
}

func (usersDao) GetDataByUsername(username string, resourceID string) (result repository.Users, errMdl model.ErrorModel) {
	db := common.GormDB
	err := db.Raw("SELECT u.id, u.password, u.salt, u.client_id, "+
		"ac.secret_key, ac.grant_type, cr.resource_id, "+
		"cr.authorities, u.phone, u.locale "+
		"FROM users u "+
		"LEFT JOIN auth_client ac ON u.client_id = ac.client_id "+
		"LEFT JOIN client_resource cr ON u.client_id = cr.client_id "+
		"WHERE u.username = @username AND cr.resource_id = @resourceID",
		sql.Named("username", username),
		sql.Named("resourceID", resourceID)).Row().
		Scan(&result.ID, &result.Password, &result.Salt, &result.ClientID,
			&result.AuthClient.SecretKey, &result.AuthClient.GrantType, &result.ClientResource.ResourceID,
			&result.ClientResource.Authorities, &result.Phone, &result.Locale)
	if err != nil {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}
