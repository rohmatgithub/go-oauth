package service

import (
	"go-oauth/dto"
	"go-oauth/model"

	"github.com/gofiber/fiber/v2"
)

func ValidateList(c *fiber.Ctx, validOrderBy []string, validOperator map[string]dto.DefaultOperator) (dtoList dto.GetListRequest, listSearch []dto.SearchByParam, errMdl model.ErrorModel) {
	dtoList = dto.GetListRequest{
		Page:    c.QueryInt("page"),
		Limit:   c.QueryInt("limit"),
		OrderBy: c.Query("order_by"),
		Filter:  c.Query("filter"),
	}

	errMdl = dtoList.ValidateInputPageLimitAndOrderBy([]int{10, 30, 50, 100, -99}, validOrderBy)
	if errMdl.Error != nil {
		return
	}

	listSearch, errMdl = dtoList.ValidateFilter(validOperator)
	return
}

func ValidateCount(c *fiber.Ctx, validOperator map[string]dto.DefaultOperator) (listSearch []dto.SearchByParam, errMdl model.ErrorModel) {
	dtoList := dto.GetListRequest{
		Filter: c.Query("filter"),
	}

	listSearch, errMdl = dtoList.ValidateFilter(validOperator)
	return
}
