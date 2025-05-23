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
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	models "template_backend/open-api/models"
	runtime "template_backend/open-api/runtime"

	"github.com/gorilla/mux"
)

// AdministrationAPIController binds http requests to an api service and writes the service results to the http response
type AdministrationAPIController struct {
	service      AdministrationAPIServicer
	errorHandler runtime.ErrorHandler
}

// AdministrationAPIOption for how the controller is set up.
type AdministrationAPIOption func(*AdministrationAPIController)

// WithAdministrationAPIErrorHandler inject ErrorHandler into controller
func WithAdministrationAPIErrorHandler(h runtime.ErrorHandler) AdministrationAPIOption {
	return func(c *AdministrationAPIController) {
		c.errorHandler = h
	}
}

// NewAdministrationAPIController creates a default api controller
func NewAdministrationAPIController(s AdministrationAPIServicer, opts ...AdministrationAPIOption) runtime.Router {
	controller := &AdministrationAPIController{
		service:      s,
		errorHandler: runtime.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the AdministrationAPIController
func (c *AdministrationAPIController) Routes() runtime.Routes {
	return runtime.Routes{
		"AdministrationChangeRoleUserIdPost": runtime.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "/administration/change-role/{userId}",
			HandlerFunc: c.AdministrationChangeRoleUserIdPost,
		},
		"AdministrationRefundDepositPost": runtime.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "/administration/refund-deposit",
			HandlerFunc: c.AdministrationRefundDepositPost,
		},
		"AdministrationSalesGet": runtime.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/administration/sales",
			HandlerFunc: c.AdministrationSalesGet,
		},
		"AdministrationUsersGet": runtime.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/administration/users",
			HandlerFunc: c.AdministrationUsersGet,
		},
	}
}

// AdministrationChangeRoleUserIdPost - Change role of an user
func (c *AdministrationAPIController) AdministrationChangeRoleUserIdPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIdParam := params["userId"]
	if userIdParam == "" {
		c.errorHandler(w, r, &models.RequiredError{"userId"}, nil)
		return
	}
	changeRoleParam := models.ChangeRole{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&changeRoleParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertChangeRoleRequired(changeRoleParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertChangeRoleConstraints(changeRoleParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AdministrationChangeRoleUserIdPost(r.Context(), userIdParam, changeRoleParam, r)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	runtime.EncodeJSONResponse(result.Body, &result.Code, w)
}

// AdministrationRefundDepositPost - Refund deposit of a payment transaction
func (c *AdministrationAPIController) AdministrationRefundDepositPost(w http.ResponseWriter, r *http.Request) {
	refundDepositParam := models.RefundDeposit{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&refundDepositParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &models.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertRefundDepositRequired(refundDepositParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertRefundDepositConstraints(refundDepositParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AdministrationRefundDepositPost(r.Context(), refundDepositParam, r)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	runtime.EncodeJSONResponse(result.Body, &result.Code, w)
}

// AdministrationSalesGet - Retrieve all sales
func (c *AdministrationAPIController) AdministrationSalesGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.AdministrationSalesGet(r.Context(), r)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	runtime.EncodeJSONResponse(result.Body, &result.Code, w)
}

// AdministrationUsersGet - Retrieve all users
func (c *AdministrationAPIController) AdministrationUsersGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.AdministrationUsersGet(r.Context(), r)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	runtime.EncodeJSONResponse(result.Body, &result.Code, w)
}
