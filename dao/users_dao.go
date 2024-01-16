package dao

import (
	"database/sql"
	"errors"
	"go-oauth/common"
	"go-oauth/dto"
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

func (u usersDao) GetDataByUsername(username string, resourceID string) (result repository.UsersDetail, errMdl model.ErrorModel) {
	db := common.GormDB
	err := db.Raw("SELECT u.id, u.password, u.salt, u.client_id, "+
		"ac.secret_key, ac.grant_type, cr.resource_id, "+
		"cr.authorities, u.phone, u.locale, ap.company_id "+
		"FROM users u "+
		"LEFT JOIN auth_client ac ON u.client_id = ac.client_id "+
		"LEFT JOIN client_resource cr ON u.client_id = cr.client_id "+
		"LEFT JOIN access_permissions ap ON ap.user_id = u.id "+
		"WHERE u.username = @username AND cr.resource_id = @resourceID",
		sql.Named("username", username),
		sql.Named("resourceID", resourceID)).Row().
		Scan(&result.ID, &result.Password, &result.Salt, &result.ClientID,
			&result.AuthClient.SecretKey, &result.AuthClient.GrantType, &result.ClientResource.ResourceID,
			&result.ClientResource.Authorities, &result.Phone, &result.Locale, &result.CompanyID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}

func (u usersDao) GetList(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel) {
	query := "SELECT u.id, u.username, pp.first_name, " +
		"pp.last_name, pp.address_1, u.created_at, u.updated_at  " +
		"FROM users u " +
		"LEFT JOIN person_profile pp ON u.person_profile_id = pp.id  "

	return GetListDataDefault(common.GormDB, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp repository.Users
			err := rows.Scan(&temp.ID, &temp.Username, &temp.PersonProfile.FirstName,
				&temp.PersonProfile.LastName, &temp.PersonProfile.Address1,
				&temp.CreatedAt, &temp.UpdatedAt)
			return temp, err
		})
}
