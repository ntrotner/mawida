package openapi

import (
	"net/http"
	config "template_backend/core/config"
	administration "template_backend/open-api/handlers/administration"
	authentication "template_backend/open-api/handlers/authentication"
	core "template_backend/open-api/handlers/core"
	location "template_backend/open-api/handlers/location"
	payment "template_backend/open-api/handlers/payment"
	product "template_backend/open-api/handlers/product"
	rental "template_backend/open-api/handlers/rental"
	user "template_backend/open-api/handlers/user"
	runtime "template_backend/open-api/runtime"

	"github.com/rs/zerolog/log"
)

func SetupHttp() {
	AdministrationAPIService := administration.NewAdministrationAPIService()
	AdministrationAPIController := administration.NewAdministrationAPIController(AdministrationAPIService)

	AuthenticationAPIService := authentication.NewAuthenticationAPIService()
	AuthenticationAPIController := authentication.NewAuthenticationAPIController(AuthenticationAPIService)

	StatusAPIService := core.NewStatusAPIService()
	StatusAPIController := core.NewStatusAPIController(StatusAPIService)

	UserAPIService := user.NewUserAPIService()
	UserAPIController := user.NewUserAPIController(UserAPIService)

	LocationAPIService := location.NewLocationAPIService()
	LocationAPIController := location.NewLocationAPIController(LocationAPIService)

	ProductAPIService := product.NewProductAPIService()
	ProductAPIController := product.NewProductAPIController(ProductAPIService)

	RentalAPIService := rental.NewRentalAPIService()
	RentalAPIController := rental.NewRentalAPIController(RentalAPIService)

	PaymentAPIService := payment.NewPaymentAPIService()
	PaymentAPIController := payment.NewPaymentAPIController(PaymentAPIService)

	router := runtime.NewRouter(
		AdministrationAPIController,
		AuthenticationAPIController,
		StatusAPIController,
		UserAPIController,
		LocationAPIController,
		ProductAPIController,
		RentalAPIController,
		PaymentAPIController,
	)
	core.SetupPerformanceLogger(router)
	core.SetupSwaggerUi(router)

	log.Info().Msg("Listening on: " + config.GlobalConfig.Server.Address)
	log.Fatal().Msg(http.ListenAndServe(config.GlobalConfig.Server.Address, router).Error())
}
