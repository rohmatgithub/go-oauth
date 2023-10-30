package router

import (
	"github.com/gofiber/fiber/v2"
	"go-oauth/endpoint"
	"go-oauth/service/credentialsservice"
)

func credentialsRouter(app fiber.Router) {
	var ae endpoint.AbstractEndpoint
	app.Post("/verify", func(ctx *fiber.Ctx) error {
		return ae.EndpointWhiteList(ctx, credentialsservice.CredentialsService.VerifyService)
	})

	app.Post("/clientcredentials", func(ctx *fiber.Ctx) error {
		return ae.EndpointWhiteList(ctx, credentialsservice.CredentialsService.ClientCredentialsService)
	})

	app.Post("/verify/internal", func(ctx *fiber.Ctx) error {
		return ae.EndpointClientCredentials(ctx, credentialsservice.CredentialsService.VerifyService)
	})
}
