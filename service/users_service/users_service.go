package users_service

import (
	"database/sql"
	"go-oauth/common"
	"go-oauth/constanta"
	"go-oauth/dao"
	"go-oauth/dto"
	"go-oauth/dto/in"
	"go-oauth/dto/out"
	"go-oauth/model"
	"go-oauth/repository"
	"go-oauth/service"
	"go-oauth/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var UsersService = usersService{}

type usersService struct {
	service.AbstractService
}

func (u usersService) InsertUser(c *fiber.Ctx, ctxModel *common.ContextModel) (payload out.Payload, errMdl model.ErrorModel) {
	var dtoIn in.UsersRequest
	errMdl = u.Initialize(c, ctxModel).ReadBodyAndValidate(&dtoIn).ErrMdl
	if errMdl.Error != nil {
		return
	}

	validated, errMdl := dtoIn.ValidateInsert(ctxModel)
	if errMdl.Error != nil {
		return
	}
	if validated != nil {
		payload.Status.Detail = validated
		errMdl = model.GenerateFailedValidate()
		return
	}

	// cek username
	userDB, errMdl := dao.UserDao.GetDataByUsername(dtoIn.Username, "auth")
	if errMdl.Error != nil {
		return
	}
	if userDB.ID.Int64 > 0 {
		errMdl = model.GenerateDuplicateDataError(constanta.Username)
		return
	}

	timeNow := time.Now()

	tx := common.GormDB.Begin()
	if tx.Error != nil {
		errMdl = model.GenerateUnknownError(tx.Error)
		return
	}

	defer func() {
		var err *gorm.DB
		if errMdl.Error != nil {
			err = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		if err.Error != nil {
			errMdl = model.GenerateInternalDBServerError(err.Error)
		}
	}()

	// insert person profile
	repoPersonProfile := repository.PersonProfile{
		FirstName: sql.NullString{String: dtoIn.FirstName, Valid: true},
		LastName:  sql.NullString{String: dtoIn.LastName, Valid: true},
		Address1:  sql.NullString{String: dtoIn.Address1, Valid: true},
		Address2:  sql.NullString{String: dtoIn.Address2, Valid: true},
		CountryID: sql.NullInt64{Int64: dtoIn.CountryID, Valid: true},
		AbstractModel: repository.AbstractModel{
			CreatedAt: sql.NullTime{Time: timeNow, Valid: true},
			UpdatedAt: sql.NullTime{Time: timeNow, Valid: true},
			CreatedBy: sql.NullInt64{Int64: ctxModel.AuthAccessTokenModel.ResourceUserID, Valid: true},
			UpdatedBy: sql.NullInt64{Int64: ctxModel.AuthAccessTokenModel.ResourceUserID, Valid: true},
		},
	}
	errMdl = dao.PersonProfileDao.Insert(tx, &repoPersonProfile)
	if errMdl.Error != nil {
		return
	}

	// insert user
	salt := util.GetUUID()
	password := util.HashingPassword(dtoIn.Password, salt)
	repoUser := repository.Users{
		Username:        sql.NullString{String: dtoIn.Username, Valid: true},
		Password:        sql.NullString{String: password, Valid: true},
		Salt:            sql.NullString{String: salt, Valid: true},
		Email:           sql.NullString{String: dtoIn.Email, Valid: true},
		Phone:           sql.NullString{String: dtoIn.Phone, Valid: true},
		Locale:          sql.NullString{String: dtoIn.Locale, Valid: true},
		ClientID:        sql.NullString{String: util.GetUUID(), Valid: true},
		PersonProfileID: sql.NullInt64{Int64: repoPersonProfile.ID.Int64, Valid: true},
		AbstractModel: repository.AbstractModel{
			CreatedAt: sql.NullTime{Time: timeNow, Valid: true},
			UpdatedAt: sql.NullTime{Time: timeNow, Valid: true},
			CreatedBy: sql.NullInt64{Int64: ctxModel.AuthAccessTokenModel.ResourceUserID, Valid: true},
			UpdatedBy: sql.NullInt64{Int64: ctxModel.AuthAccessTokenModel.ResourceUserID, Valid: true},
		},
	}
	errMdl = dao.UserDao.Insert(tx, &repoUser)
	if errMdl.Error != nil {
		return
	}

	// insert access permission
	repoAccPermission := repository.AccessPermission{
		UserID:    repoUser.ID.Int64,
		CompanyID: dtoIn.CompanyID,
		BranchID:  dtoIn.BranchID,
		IsAdmin:   dtoIn.IsAdmin,
	}
	errMdl = dao.AccessPermissionsDao.InsertData(tx, repoAccPermission)
	if errMdl.Error != nil {
		return
	}

	crRepo := repository.ClientResource{
		ClientID:    sql.NullString{String: repoUser.ClientID.String, Valid: true},
		ResourceID:  sql.NullString{String: "auth", Valid: true},
		Authorities: sql.NullString{String: `{"ALL":["ALL"]}`, Valid: true},
		AbstractModel: repository.AbstractModel{
			CreatedAt: sql.NullTime{Time: timeNow, Valid: true},
			UpdatedAt: sql.NullTime{Time: timeNow, Valid: true},
			CreatedBy: sql.NullInt64{Int64: ctxModel.AuthAccessTokenModel.ResourceUserID, Valid: true},
			UpdatedBy: sql.NullInt64{Int64: ctxModel.AuthAccessTokenModel.ResourceUserID, Valid: true},
		},
	}
	errMdl = dao.ClientResourceDao.Insert(tx, &crRepo)
	payload.Status.Message = service.InsertI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (u usersService) ListUser(c *fiber.Ctx, ctxModel *common.ContextModel) (payload out.Payload, errMdl model.ErrorModel) {

	dtoList, listSearch, errMdl := service.ValidateList(c, []string{"id", "username", "updated_at"}, dto.ValidOperatorUser)
	if errMdl.Error != nil {
		return
	}
	resultDao, errMdl := dao.UserDao.GetList(dtoList, listSearch)
	if errMdl.Error != nil {
		return
	}

	var listOut []out.UsersList

	for _, data := range resultDao {
		temp := data.(repository.Users)
		listOut = append(listOut, out.UsersList{
			ID:        temp.ID.Int64,
			Username:  temp.Username.String,
			FirstName: temp.PersonProfile.FirstName.String,
			LastName:  temp.PersonProfile.LastName.String,
			Address1:  temp.PersonProfile.Address1.String,
			CreatedAt: temp.CreatedAt.Time,
			UpdatedAt: temp.UpdatedAt.Time,
		})
	}

	payload.Data = listOut
	payload.Status.Message = service.ListI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}
