package payment

import (
	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/customer"
)

// CustomerData represents the data needed to create a customer
type CustomerData struct {
	Email       string
	Name        string
	Description string
	Metadata    map[string]string
}

// CreateCustomer creates a new customer in Stripe
func CreateCustomer(data CustomerData) (*stripe.Customer, error) {
	// Initialize Stripe if not already initialized
	InitStripe()

	params := &stripe.CustomerParams{
		Email:       stripe.String(data.Email),
		Name:        stripe.String(data.Name),
		Description: stripe.String(data.Description),
	}

	if data.Metadata != nil {
		params.Metadata = make(map[string]string)
		for k, v := range data.Metadata {
			params.Metadata[k] = v
		}
	}

	log.Info().Str("email", data.Email).Msg("Creating customer in Stripe")
	return customer.New(params)
}

// GetCustomer retrieves a customer from Stripe by ID
func GetCustomer(customerID string) (*stripe.Customer, error) {
	// Initialize Stripe if not already initialized
	InitStripe()

	log.Info().Str("customerID", customerID).Msg("Retrieving customer from Stripe")
	return customer.Get(customerID, nil)
}

// UpdateCustomer updates a customer in Stripe
func UpdateCustomer(customerID string, data CustomerData) (*stripe.Customer, error) {
	// Initialize Stripe if not already initialized
	InitStripe()

	params := &stripe.CustomerParams{}

	if data.Email != "" {
		params.Email = stripe.String(data.Email)
	}

	if data.Name != "" {
		params.Name = stripe.String(data.Name)
	}

	if data.Description != "" {
		params.Description = stripe.String(data.Description)
	}

	if data.Metadata != nil {
		params.Metadata = make(map[string]string)
		for k, v := range data.Metadata {
			params.Metadata[k] = v
		}
	}

	log.Info().Str("customerID", customerID).Msg("Updating customer in Stripe")
	return customer.Update(customerID, params)
}

func DeleteCustomer(customerID string) (*stripe.Customer, error) {
	// Initialize Stripe if not already initialized
	InitStripe()

	log.Info().Str("customerID", customerID).Msg("Deleting customer from Stripe")
	return customer.Del(customerID, nil)
}
