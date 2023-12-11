package cmd

import (
	"admin-panel/v1/app"
	"admin-panel/v1/pkg/config"
	"admin-panel/v1/pkg/middleware"
	"admin-panel/v1/pkg/route"
)

var (
	coreAPI *app.App
)

func StartServer() {
	coreAPI = app.New()

	coreAPI.Worker(Setting(config.BrokerUrl, config.ResultBackend))

	coreAPI.Middleware(
		middleware.GinMiddleware,
	)
	coreAPI.Route(
		route.ApiRoute,
		route.ChatRoute,
	)
	coreAPI.Run(":8000")
}
