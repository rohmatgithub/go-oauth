package dao

import (
	"go-oauth/common"
	"go-oauth/model"
)

var AuthClientDao = authClientDao{}

type authClientDao struct {
	abstractDao
}

func (ac authClientDao) ViewDetail(repo interface{}) (errMdl model.ErrorModel) {
	db := common.GormDB
	//temp := repo.(*repository.AuthClient)
	resultDB := db.
		Preload("ClientResource").
		Find(repo)
	if resultDB.Error != nil {
		errMdl = model.GenerateInternalDBServerError(resultDB.Error)
		return
	}

	//result = repo
	return
}

//func (ac authClientDao) CheckByResourceID(repo *repository.AuthClient) (isExists bool, errMdl model.ErrorModel) {
//	db := common.GormDB
//	err := db.Raw("SELECT ac.grant_type, cr.resource_id, cr.authorities "+
//		"FROM auth_client ac "+
//		"LEFT JOIN client_resource cr ON ac.client_id = cr.client_id "+
//		"WHERE ac.client_id = @clientID AND ac.client_secret = @clientSecret AND "+
//		"cr.resource_id = @resourceID AND grant_type = @grantType ",
//		sql.Named("clientID", repo.ClientID.String), sql.Named("clientSecret", repo.ClientSecret.String),
//		sql.Named("resourceID", repo.ClientResource.ResourceID.String), sql.Named("grantType", repo.GrantType.String)).Row().
//		Scan(&repo.GrantType, &repo.ClientResource.ResourceID, &repo.ClientResource.Authorities)
//	if err != nil && !errors.Is(err, sql.ErrNoRows) {
//		errMdl = model.GenerateInternalDBServerError(err)
//		return
//	}
//
//	isExists = err == nil
//
//	return
//}
