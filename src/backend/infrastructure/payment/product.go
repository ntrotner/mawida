package payment

import (
	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/price"
	"github.com/stripe/stripe-go/v82/product"
	"github.com/stripe/stripe-go/v82/refund"
)

// ProductData represents the data needed to create a product
type ProductData struct {
	Name        string
	Description string
	Images      []string
	Metadata    map[string]string
}

// PriceData represents the data needed to create a price
type PriceData struct {
	Currency   string
	UnitAmount int64
	Recurring  *stripe.PriceRecurringParams
	Metadata   map[string]string
}

// CreateProduct creates a new product in Stripe
func CreateProduct(data ProductData) (*stripe.Product, error) {
	// Initialize Stripe if not already initialized
	InitStripe()

	params := &stripe.ProductParams{
		Name:        stripe.String(data.Name),
		Description: stripe.String(data.Description),
	}

	if len(data.Images) > 0 {
		params.Images = stripe.StringSlice(data.Images)
	}

	if data.Metadata != nil {
		params.Metadata = make(map[string]string)
		for k, v := range data.Metadata {
			params.Metadata[k] = v
		}
	}

	log.Info().Str("product", data.Name).Msg("Creating product in Stripe")
	return product.New(params)
}

// CreatePrice creates a new price for a product in Stripe
func CreatePrice(productID string, data PriceData) (*stripe.Price, error) {
	// Initialize Stripe if not already initialized
	InitStripe()

	params := &stripe.PriceParams{
		Product:    stripe.String(productID),
		Currency:   stripe.String(data.Currency),
		UnitAmount: stripe.Int64(data.UnitAmount),
	}

	if data.Recurring != nil {
		params.Recurring = data.Recurring
	}

	if data.Metadata != nil {
		params.Metadata = make(map[string]string)
		for k, v := range data.Metadata {
			params.Metadata[k] = v
		}
	}

	log.Info().Str("productID", productID).Str("currency", data.Currency).Msg("Creating price in Stripe")
	return price.New(params)
}

func DeleteProduct(productID string) (*stripe.Product, error) {
	// Initialize Stripe if not already initialized
	InitStripe()

	return product.Del(productID, nil)
}

func RefundPrice(paymentIntentID string, amount int64) (*stripe.Refund, error) {
	// Initialize Stripe if not already initialized
	InitStripe()

	params := &stripe.RefundParams{
		PaymentIntent: stripe.String(paymentIntentID),
		Amount:        stripe.Int64(amount),
	}
	return refund.New(params)
}
