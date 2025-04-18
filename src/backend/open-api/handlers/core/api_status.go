/*
 * Swagger - OpenAPI 3.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Contact: nikita@ttnr.me
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
	"strings"
	runtime "template_backend/open-api/runtime"
)

// StatusAPIController binds http requests to an api service and writes the service results to the http response
type StatusAPIController struct {
	service      StatusAPIServicer
	errorHandler runtime.ErrorHandler
}

// StatusAPIOption for how the controller is set up.
type StatusAPIOption func(*StatusAPIController)

// WithStatusAPIErrorHandler inject ErrorHandler into controller
func WithStatusAPIErrorHandler(h runtime.ErrorHandler) StatusAPIOption {
	return func(c *StatusAPIController) {
		c.errorHandler = h
	}
}

// NewStatusAPIController creates a default api controller
func NewStatusAPIController(s StatusAPIServicer, opts ...StatusAPIOption) runtime.Router {
	controller := &StatusAPIController{
		service:      s,
		errorHandler: runtime.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the StatusAPIController
func (c *StatusAPIController) Routes() runtime.Routes {
	return runtime.Routes{
		"Health": runtime.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/health",
			HandlerFunc: c.Health,
		},
	}
}

// Health - Check health of server
func (c *StatusAPIController) Health(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.Health(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	runtime.EncodeJSONResponse(result.Body, &result.Code, w)
}
