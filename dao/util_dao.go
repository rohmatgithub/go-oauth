package dao

import (
	"database/sql"
	"fmt"
	"go-oauth/dto"
	"go-oauth/model"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func GetListDataDefault(gormDB *gorm.DB, query string, queryParam []interface{},
	dtoList dto.GetListRequest, searchBy []dto.SearchByParam,
	wrap func(rows *sql.Rows) (interface{}, error)) (result []interface{}, errMdl model.ErrorModel) {

	queryParam, queryCondition := SearchByParamToQuery(searchBy, queryParam)
	query += queryCondition + fmt.Sprintf(" ORDER BY %s ", dtoList.OrderBy)

	if dtoList.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %d OFFSET %d ", dtoList.Limit, countOffset(dtoList.Page, dtoList.Limit))
	}

	return ExecuteQuery(gormDB, query, queryParam, wrap)
}

func GetCountDataDefault(gormDB *gorm.DB, query string, queryParam []interface{}, searchBy []dto.SearchByParam) (result int64, errMdl model.ErrorModel) {

	queryParam, queryCondition := SearchByParamToQuery(searchBy, queryParam)
	query += queryCondition

	var temp sql.NullInt64
	gormCallBack := gormDB.Raw(query, queryParam...).Scan(&temp)
	if gormCallBack.Error != nil {
		errMdl = model.GenerateUnknownError(gormCallBack.Error)
		return
	}

	result = temp.Int64
	return
}

func SearchByParamToQuery(searchByParam []dto.SearchByParam, queryParam []interface{}) (resultQueryParam []interface{}, result string) {
	if len(queryParam) == 0 && len(searchByParam) > 0 {
		result = "WHERE \n"
	}
	var (
		operator          string
		searchConditionOr []dto.SearchByParam
	)
	index := len(queryParam)
	for i := 0; i < len(searchByParam); i++ {
		if searchByParam[i].Condition == "OR" {
			searchConditionOr = append(searchConditionOr, searchByParam[i])
			continue
		}
		if index > 0 {
			result += " AND "
		}
		index++
		if searchByParam[i].DataType == "enum" {
			searchByParam[i].SearchKey = "cast( " + searchByParam[i].SearchKey + " AS VARCHAR)"
		}
		searchByParam[i], operator = getOperator(searchByParam[i])

		if searchByParam[i].SearchOperator == "between" {
			operator = "between"
			result += fmt.Sprintf(" %s %s $%d AND $%d ", searchByParam[i].SearchKey, operator, len(queryParam)+1, len(queryParam)+2)
			searchValueSplit := strings.Fields(searchByParam[i].SearchValue)
			queryParam = append(queryParam, searchValueSplit[0])
			if len(searchValueSplit) > 1 {
				queryParam = append(queryParam, searchValueSplit[1])
			}
		} else if searchByParam[i].SearchOperator == "in" || searchByParam[i].SearchOperator == "not_in" {
			queryList := getQueryListValue(len(queryParam)+1, searchByParam[i].ListValue)
			result += " " + searchByParam[i].SearchKey + " " + operator + " (" + queryList + ")"
			queryParam = append(queryParam, searchByParam[i].ListValue...)
		} else {
			result += fmt.Sprintf(" %s %s $%d ", searchByParam[i].SearchKey, operator, len(queryParam)+1)
			queryParam = append(queryParam, searchByParam[i].SearchValue)
		}

	}

	queryOR := ""
	for i := 0; i < len(searchConditionOr); i++ {
		searchConditionOr[i], operator = getOperator(searchConditionOr[i])
		if searchConditionOr[i].SearchOperator == "between" {
			queryOR += fmt.Sprintf(" %s %s $%d AND $%d", searchConditionOr[i].SearchKey, searchConditionOr[i].SearchOperator, len(queryParam)+1, len(queryParam)+2)
			searchValueSplit := strings.Fields(searchConditionOr[i].SearchValue)
			queryParam = append(queryParam, searchValueSplit[0])
			if len(searchValueSplit) > 1 {
				queryParam = append(queryParam, searchValueSplit[1])
			}
		} else {
			switch operator {
			case "in", "not in":
				queryList := getQueryListValue(len(queryParam)+1, searchConditionOr[i].ListValue)
				queryOR += fmt.Sprintf(" %s %s (%s) ", searchConditionOr[i].SearchKey, operator, queryList)
				queryParam = append(queryParam, searchConditionOr[i].ListValue...)
			default:
				queryOR += fmt.Sprintf(" %s %s $%d ", searchConditionOr[i].SearchKey, operator, len(queryParam)+1)
				queryParam = append(queryParam, searchConditionOr[i].SearchValue)
			}
		}
		if i < len(searchConditionOr)-1 {
			queryOR += " OR "
		}
	}
	if len(searchConditionOr) > 0 {
		queryOR = fmt.Sprintf(" ( %s )", queryOR)
		if index > 0 {
			queryOR = " AND " + queryOR
		}
		result += queryOR
	}
	resultQueryParam = queryParam
	return
}

func getOperator(searchByParam dto.SearchByParam) (result dto.SearchByParam, operator string) {
	operator = searchByParam.SearchOperator
	switch operator {
	case "like":
		//searchByParam.SearchKey = "LOWER(" + searchByParam.SearchKey + ")"
		operator = "ilike"
		//searchByParam.SearchValue = strings.ToLower(searchByParam.SearchValue)
		searchByParam.SearchValue = "%" + searchByParam.SearchValue + "%"
	case "eq":
		operator = "="
	case "not_eq":
		operator = "!="
	case "not_like":
		operator = "not ilike "
		searchByParam.SearchKey = "LOWER(" + searchByParam.SearchKey + ")"
		//searchByParam.SearchValue = strings.ToLower(searchByParam.SearchValue)
		searchByParam.SearchValue = "%" + searchByParam.SearchValue + "%"
	case "between":
		operator = "between"
	//20-02-2022 -NEXCORE
	//-- Start Perubahan
	case "in":
		operator = "in"
	case "not_in":
		operator = "not in"
	}
	return searchByParam, operator
}

func getQueryListValue(currentIndex int, listValue []interface{}) (result string) {
	index := currentIndex
	for i := 0; i < len(listValue); i++ {
		if i != len(listValue)-1 {
			result += " $" + strconv.Itoa(index) + ","
		} else {
			result += " $" + strconv.Itoa(index)
		}
		index++
	}
	return result
}

func ExecuteQuery(gormDB *gorm.DB, query string, queryParam []interface{},
	wrap func(rows *sql.Rows) (interface{}, error)) (result []interface{}, errMdl model.ErrorModel) {

	rows, err := gormDB.Raw(query, queryParam...).Rows()
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
		return
	}

	if rows != nil {
		defer func() {
			errMdl = closeRow(rows, errMdl)
		}()
		var temp interface{}
		for rows.Next() {
			temp, err = wrap(rows)
			if err != nil {
				errMdl = model.GenerateInternalDBServerError(err)
				return
			}
			result = append(result, temp)
		}
	}

	return
}

func closeRow(rows *sql.Rows, inputErr model.ErrorModel) (errMdl model.ErrorModel) {
	err := rows.Close()
	if err != nil {
		errMdl = model.GenerateInternalDBServerError(err)
	} else {
		errMdl = inputErr
	}
	return
}

func countOffset(page int, limit int) int {
	return (page - 1) * limit
}
