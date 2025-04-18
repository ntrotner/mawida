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
	"context"
	"errors"
	"net/http"
	database_user "template_backend/database/paths/user"
	authentication "template_backend/infrastructure/authentication"
	api_authentication "template_backend/open-api/authentication"
	models "template_backend/open-api/models"

	"github.com/rs/zerolog/log"
)

// UserAPIRouter defines the required methods for binding the api requests to a responses for the UserAPI
// The UserAPIRouter implementation should parse necessary information from the http request,
// pass the data to a UserAPIServicer to perform the required actions, then write the service results to the http response.
type UserAPIRouter interface {
	ChangeEmailPost(http.ResponseWriter, *http.Request)
	ChangePasswordPost(http.ResponseWriter, *http.Request)
	PasswordResetPost(http.ResponseWriter, *http.Request)
	ProfileGet(http.ResponseWriter, *http.Request)
}

// UserAPIServicer defines the api actions for the UserAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type UserAPIServicer interface {
	ChangeEmailPost(context.Context, models.ChangeEmail, *http.Request) (models.ImplResponse, error)
	ChangePasswordPost(context.Context, models.ChangePassword, *http.Request) (models.ImplResponse, error)
	PasswordResetPost(context.Context, models.PasswordReset) (models.ImplResponse, error)
	ProfileGet(context.Context, *http.Request) (models.ImplResponse, error)
}

// UserAPIService is a service that implements the logic for the UserAPIServicer
// This service should implement the business logic for every endpoint for the UserAPI API.
// Include any external packages or services that will be required by this service.
type UserAPIService struct {
}

// NewUserAPIService creates a default api service
func NewUserAPIService() UserAPIServicer {
	return &UserAPIService{}
}

// ChangeEmailPost - Change user email
func (s *UserAPIService) ChangeEmailPost(ctx context.Context, changeEmail models.ChangeEmail, r *http.Request) (models.ImplResponse, error) {
	token, found := api_authentication.ReadTokenFromHeader(r)
	if !found {
		log.Error().Msg("Bearer format invalid")
		return models.Response(401, models.Error{ErrorMessages: []models.Message{{Code: "100", Message: "Unauthorized. Please check your credentials."}}}), nil
	}

	_, content, err := authentication.VerifyJWT(&token)
	if err != nil {
		log.Error().Msg("Couldn't verify token to refresh")
		return models.Response(401, models.Error{ErrorMessages: []models.Message{{Code: "100", Message: "Unauthorized. Please check your credentials."}}}), nil
	}

	_, err = database_user.ChangeUserEmail(ctx, &content.ID, &changeEmail.NewEmail)
	if err != nil {
		log.Error().Str("id", content.ID).Msg("Couldn't update email for user")
		return models.Response(401, models.Error{ErrorMessages: []models.Message{{Code: "100", Message: "Unauthorized. Please check your credentials."}}}), nil
	}

	return models.Response(200, models.Success{}), nil
}

// ChangePasswordPost - Change user password
func (s *UserAPIService) ChangePasswordPost(ctx context.Context, changePassword models.ChangePassword, r *http.Request) (models.ImplResponse, error) {
	token, found := api_authentication.ReadTokenFromHeader(r)
	if !found {
		log.Error().Msg("Bearer format invalid")
		return models.Response(401, models.Error{ErrorMessages: []models.Message{{Code: "100", Message: "Unauthorized. Please check your credentials."}}}), nil
	}

	_, content, err := authentication.VerifyJWT(&token)
	if err != nil {
		log.Error().Msg("Couldn't verify token to refresh")
		return models.Response(401, models.Error{ErrorMessages: []models.Message{{Code: "100", Message: "Unauthorized. Please check your credentials."}}}), nil
	}

	user := database_user.AuthenticateUserById(ctx, content.ID, changePassword.CurrentPassword)
	if user == nil {
		log.Error().Str("id", content.ID).Msg("Couldn't find user by id")
		return models.Response(401, models.Error{ErrorMessages: []models.Message{{Code: "100", Message: "Unauthorized. Please check your credentials."}}}), nil
	}

	_, err = database_user.ChangeUserPassword(ctx, &content.ID, &changePassword.NewPassword)
	if err != nil {
		log.Error().Str("id", content.ID).Msg("Couldn't update email for user")
		return models.Response(401, models.Error{ErrorMessages: []models.Message{{Code: "100", Message: "Unauthorized. Please check your credentials."}}}), nil
	}

	return models.Response(200, models.Success{}), nil
}

// PasswordResetPost - Initiate password reset
func (s *UserAPIService) PasswordResetPost(ctx context.Context, passwordReset models.PasswordReset) (models.ImplResponse, error) {
	// TODO - update PasswordResetPost with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Success{}) or use other options such as http.Ok ...
	// return Response(200, Success{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	return models.Response(http.StatusNotImplemented, nil), errors.New("PasswordResetPost method not implemented")
}

// ProfileGet - Get user profile
func (s *UserAPIService) ProfileGet(ctx context.Context, r *http.Request) (models.ImplResponse, error) {
	user, err := api_authentication.IsUserAuthorized(ctx, r)
	if err != nil {
		log.Error().Msg(err.Error())
		return models.Response(401, models.Error{ErrorMessages: []models.Message{{Code: "100", Message: "Unauthorized."}}}), nil
	}
	sanitized := database_user.SanitizeUserProfile(user)
	return models.Response(200, sanitized), nil
}
