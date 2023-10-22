package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"go-oauth/common"
)

type AbstractEndpoint struct {
	common.AbstractStruct
}

func (ae AbstractEndpoint) EndpointWhiteList(c *fiber.Ctx) error {

	return nil
}
