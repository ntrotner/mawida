package openapi

import (
	"context"
	"net/http"
	"strings"
	models "template_backend/open-api/models"
	runtime "template_backend/open-api/runtime"
)

// RentalAPIRouter defines the required methods for binding the api requests to a responses for the RentalAPI
// The RentalAPIRouter implementation should parse necessary information from the http request,
// pass the data to a RentalAPIServicer to perform the required actions, then write the service results to the http response.
type PaymentAPIRouter interface {
	Webhook(http.ResponseWriter, *http.Request)
}

// PaymentAPIServicer defines the api actions for the PaymentAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type PaymentAPIServicer interface {
	Webhook(context.Context, *http.Request) (models.ImplResponse, error)
}

// RentalAPIController binds http requests to an api service and writes the service results to the http response
type PaymentAPIController struct {
	service      PaymentAPIServicer
	errorHandler runtime.ErrorHandler
}

// PaymentAPIOption for how the controller is set up.
type PaymentAPIOption func(*PaymentAPIController)

// WithPaymentAPIErrorHandler inject ErrorHandler into controller
func WithPaymentAPIErrorHandler(h runtime.ErrorHandler) PaymentAPIOption {
	return func(c *PaymentAPIController) {
		c.errorHandler = h
	}
}

// NewRentalAPIController creates a default api controller
func NewPaymentAPIController(s PaymentAPIServicer, opts ...PaymentAPIOption) runtime.Router {
	controller := &PaymentAPIController{
		service:      s,
		errorHandler: runtime.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the RentalAPIController
func (c *PaymentAPIController) Routes() runtime.Routes {
	return runtime.Routes{
		"Webhook": runtime.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "/payment/webhook",
			HandlerFunc: c.Webhook,
		},
	}
}

// Webhook - Retrieve all rent contracts
func (c *PaymentAPIController) Webhook(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.Webhook(r.Context(), r)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	runtime.EncodeJSONResponse(result.Body, &result.Code, w)
}
