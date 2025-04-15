package payment

import (
	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/checkout/session"
)

// CheckoutItem represents an item in the checkout session
type CheckoutItem struct {
	PriceID  string
	Quantity int64
}

// CheckoutSessionConfig represents the configuration for a checkout session
type CheckoutSessionConfig struct {
	CustomerID string
	SuccessURL string
	CancelURL  string
	Items      []CheckoutItem
	Mode       string // 'payment', 'subscription', or 'setup'
	Metadata   map[string]string
}

// CreateCheckoutSession creates a new checkout session in Stripe
func CreateCheckoutSession(config CheckoutSessionConfig) (*stripe.CheckoutSession, error) {
	// Initialize Stripe if not already initialized
	InitStripe()

	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String(config.SuccessURL),
		CancelURL:  stripe.String(config.CancelURL),
		Mode:       stripe.String(config.Mode),
	}

	if config.CustomerID != "" {
		params.Customer = stripe.String(config.CustomerID)
	}

	// Add line items
	lineItems := []*stripe.CheckoutSessionLineItemParams{}
	for _, item := range config.Items {
		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(item.PriceID),
			Quantity: stripe.Int64(item.Quantity),
		})
	}
	params.LineItems = lineItems

	if config.Metadata != nil {
		params.Metadata = make(map[string]string)
		for k, v := range config.Metadata {
			params.Metadata[k] = v
		}
	}

	log.Info().Str("mode", config.Mode).Int("items", len(config.Items)).Msg("Creating checkout session in Stripe")
	return session.New(params)
}

// RetrieveCheckoutSession retrieves a checkout session from Stripe by ID
func RetrieveCheckoutSession(sessionID string) (*stripe.CheckoutSession, error) {
	// Initialize Stripe if not already initialized
	InitStripe()

	log.Info().Str("sessionID", sessionID).Msg("Retrieving checkout session from Stripe")
	return session.Get(sessionID, nil)
}

func DeleteCheckoutSession(sessionID string) (*stripe.CheckoutSession, error) {
	// Initialize Stripe if not already initialized
	InitStripe()

	log.Info().Str("sessionID", sessionID).Msg("Deleting checkout session from Stripe")
	return session.Expire(sessionID, nil)
}
