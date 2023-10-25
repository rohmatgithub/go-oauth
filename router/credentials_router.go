package router

import (
	"github.com/gofiber/fiber/v2"
	"go-oauth/endpoint"
	"go-oauth/service/passwordcredentialsservice"
)

func credentialsRouter(app *fiber.App) {
	var ae endpoint.AbstractEndpoint
	app.Post("/verify", func(ctx *fiber.Ctx) error {
		return ae.EndpointWhiteList(ctx, passwordcredentialsservice.PasswordCredentialsService.VerifyService)
	})
}
