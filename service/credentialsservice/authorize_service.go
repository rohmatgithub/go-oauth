package credentialsservice

import (
	"context"
	"encoding/json"
	"go-oauth/common"
	"go-oauth/config"
	"go-oauth/constanta"
	"go-oauth/dao"
	"go-oauth/dto/out"
	"go-oauth/model"
	"go-oauth/repository"
	"go-oauth/util"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (cs credentialsService) GetBranchID(c *fiber.Ctx, ctxModel *common.ContextModel) (payload out.Payload, errMdl model.ErrorModel) {

	// get list branchID to DB
	dataDB, errMdl := dao.AccessPermissionsDao.GetAccessPermissions(ctxModel.AuthAccessTokenModel.ResourceUserID)
	if errMdl.Error != nil {
		return
	}

	var strSlice []string
	for _, intValue := range dataDB.BranchID {
		strSlice = append(strSlice, strconv.Itoa(int(intValue)))
	}

	// Create an HTTP client
	client := &http.Client{}

	tokenInternal, errMdl := model.GetTokenInternal(dataDB.UserID, dataDB.CompanyID)
	if errMdl.Error != nil {
		return
	}

	// Join the string values with commas
	stringListID := strings.Join(strSlice, ",")
	uri := config.ApplicationConfiguration.GetUriResouce().MasterData + "/v1/master/companybranch?page=1&limit=-99&order_by=code ASC&list_id=" + stringListID
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Error("Error creating request:", err)
		errMdl = model.GenerateUnknownError(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(constanta.TokenHeaderNameConstanta, tokenInternal)

	// Make the request
	response, err := client.Do(req)
	if err != nil {
		log.Error("Error making request:", err)
		errMdl = model.GenerateUnknownError(err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error("Error reading response:", err)
		errMdl = model.GenerateUnknownError(err)
		return
	}

	var companyBranch out.CompanyBranchResponse
	err = json.Unmarshal(body, &companyBranch)
	if err != nil {
		log.Error("Error unmarshalling JSON:", err)
		errMdl = model.GenerateUnknownError(err)
		return
	}

	mapsResult := make(map[string]interface{})
	mapsResult["list"] = companyBranch.Payload.Data
	mapsResult["is_admin"] = dataDB.IsAdmin

	payload.Data = mapsResult
	return
}

func (cs credentialsService) SelectBranchID(c *fiber.Ctx, ctxModel *common.ContextModel) (payload out.Payload, errMdl model.ErrorModel) {
	type dtoStruct struct {
		BranchID int64 `json:"branch_id"`
	}
	var dtoIn dtoStruct

	errMdl = cs.Initialize(c, ctxModel).ReadBodyAndValidate(&dtoIn).ErrMdl
	if errMdl.Error != nil {
		return
	}

	tokenStr := c.Get(constanta.TokenHeaderNameConstanta)
	// get value from redis
	ctx := context.Background()
	// get value token from storage
	redis := common.RedisClient.Get(ctx, tokenStr)
	var valueRedis string
	if redis != nil {
		valueRedis = redis.Val()
	}

	// get to db
	var authTokenDB repository.AuthToken
	authTokenDB, errMdl = dao.AuthTokenDao.GetByToken(tokenStr)
	if errMdl.Error != nil {
		return
	}
	if strings.TrimSpace(valueRedis) == "" {
		valueRedis = authTokenDB.ValueToken.String
	}
	var valueModel model.ValueRedis
	err := json.Unmarshal([]byte(valueRedis), &valueModel)
	if err != nil {
		log.Error(err)
		errMdl = model.GenerateUnauthorizedClientError()
		return
	}
	// valueModel.BranchID = dtoIn.BranchID

	// update value redis
	redisStatus := common.RedisClient.Set(ctx, tokenStr, util.JsonToString(valueModel), 0)
	if redisStatus != nil && redisStatus.Err() != nil {
		log.Error(redisStatus.Err())

	}

	// update value to db
	// insert token to db
	tx := common.GormDB.Begin()
	defer func() {
		if r := recover(); r != nil || errMdl.Error != nil {
			tx.Rollback()
		} else {
			// save to redis
			err := tx.Commit().Error
			if err != nil {
				errMdl = model.GenerateInternalDBServerError(err)
				return
			}
		}
	}()

	errMdl = dao.AuthTokenDao.Update(tx, &authTokenDB)
	if errMdl.Error != nil {
		return
	}

	payload.Status.Message = "Success Authorization"
	return
}
