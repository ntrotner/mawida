package openapi

import (
	"net/http"
	"template_backend/common"
	openapi "template_backend/open-api/go"

	"github.com/rs/zerolog/log"
)

func SetupHttp() {
	AdminAPIService := openapi.NewAdminAPIService()
	AdminAPIController := openapi.NewAdminAPIController(AdminAPIService)

	AuthenticationAPIService := openapi.NewAuthenticationAPIService()
	AuthenticationAPIController := openapi.NewAuthenticationAPIController(AuthenticationAPIService)

	StatusAPIService := openapi.NewStatusAPIService()
	StatusAPIController := openapi.NewStatusAPIController(StatusAPIService)

	UserAPIService := openapi.NewUserAPIService()
	UserAPIController := openapi.NewUserAPIController(UserAPIService)

	router := openapi.NewRouter(
		AuthenticationAPIController,
		AdminAPIController,
		StatusAPIController,
		UserAPIController,
	)
	common.SetupPerformanceLogger(router)
	common.SetupSwaggerUi(router)

	log.Info().Msg("Listening on: " + common.EnvironmentConfig.Host)
	log.Fatal().Msg(http.ListenAndServe(common.EnvironmentConfig.Host, router).Error())
}
