package router

import (
	"go-oauth/endpoint"
	"go-oauth/service/credentialsservice"
	"go-oauth/service/users_service"

	"github.com/gofiber/fiber/v2"
)

func usersRouter(app fiber.Router) {
	var ae endpoint.AbstractEndpoint
	app.Post("/users", func(ctx *fiber.Ctx) error {
		return ae.EndpointJwtToken(ctx, users_service.UsersService.InsertUser)
	})

	app.Get("/users", func(ctx *fiber.Ctx) error {
		return ae.EndpointJwtToken(ctx, users_service.UsersService.ListUser)
	})

	app.Get("/authorize", func(ctx *fiber.Ctx) error {
		return ae.EndpointJwtToken(ctx, credentialsservice.CredentialsService.GetBranchID)
	})

	app.Post("/selectbranch", func(ctx *fiber.Ctx) error {
		return ae.EndpointJwtToken(ctx, credentialsservice.CredentialsService.SelectBranchID)
	})
}
