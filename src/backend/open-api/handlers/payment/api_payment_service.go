package openapi

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"template_backend/core/config"
	database_payments "template_backend/database/paths/payments"
	database_rent_contract "template_backend/database/paths/rent_contract"
	models "template_backend/open-api/models"

	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/webhook"
)

// PaymentAPIService is a service that implements the logic for the PaymentAPIServicer
// This service should implement the business logic for every endpoint for the PaymentAPI API.
// Include any external packages or services that will be required by this service.
type PaymentAPIService struct {
}

// NewPaymentAPIService creates a default api service
func NewPaymentAPIService() PaymentAPIServicer {
	return &PaymentAPIService{}
}

// WebhookEvent represents a Stripe webhook event with parsed data
type WebhookEvent struct {
	Type        stripe.EventType
	ID          string
	Data        map[string]interface{}
	Created     int64
	StripeEvent *stripe.Event
}

// HandleWebhook processes a Stripe webhook event
func (s *PaymentAPIService) Webhook(ctx context.Context, r *http.Request) (models.ImplResponse, error) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read webhook request body")
		return models.Response(400, models.Error{ErrorMessages: []models.Message{{Code: "400", Message: "Failed to read webhook request body"}}}), err
	}

	// Get the webhook signing secret
	webhookSecret := config.GlobalConfig.Stripe.WebhookSecret
	if webhookSecret == "" {
		log.Warn().Msg("STRIPE_WEBHOOK_SECRET is not set, webhook signature verification will be skipped")
	}

	// Verify signature
	var event stripe.Event
	if webhookSecret != "" {
		event, err = webhook.ConstructEvent(body, r.Header.Get("Stripe-Signature"), webhookSecret)
		if err != nil {
			log.Error().Err(err).Msg("Failed to verify webhook signature")
			return models.Response(400, models.Error{ErrorMessages: []models.Message{{Code: "400", Message: "Failed to verify webhook signature"}}}), err
		}
	} else {
		// If webhook secret is not set, just parse the event
		if err := json.Unmarshal(body, &event); err != nil {
			log.Error().Err(err).Msg("Failed to parse webhook event")
			return models.Response(400, models.Error{ErrorMessages: []models.Message{{Code: "400", Message: "Failed to parse webhook event"}}}), err
		}
	}

	// Extract data from the event
	var data stripe.CheckoutSession
	if err := json.Unmarshal(event.Data.Raw, &data); err != nil {
		log.Error().Err(err).Msg("Failed to parse webhook event data")
		return models.Response(400, models.Error{ErrorMessages: []models.Message{{Code: "400", Message: "Failed to parse webhook event data"}}}), err
	}

	if event.Type == stripe.EventTypeCheckoutSessionCompleted {
		paymentTransactionId, ok := data.Metadata["paymentTransactionID"]
		if !ok {
			log.Error().Msg("Payment transaction ID not found in metadata")
			return models.Response(400, models.Error{ErrorMessages: []models.Message{{Code: "400", Message: "Payment transaction ID not found in metadata"}}}), nil
		}

		// update rent contract status
		contract := database_rent_contract.UpdateRentContractStatusFromPaymentIdentifier(ctx, data.ID, paymentTransactionId, data.PaymentIntent.ID, database_rent_contract.RentContractStatusPickupPending)
		if contract == nil {
			log.Error().Msg("Failed to update rent contract status")
		}

		_, err = database_payments.CompletePaymentTransaction(ctx, paymentTransactionId)
		if err != nil {
			log.Error().Err(err).Msg("Failed to complete payment transaction")
		}
	}

	return models.Response(200, models.Success{}), nil
}
