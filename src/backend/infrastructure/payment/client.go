package payment

import (
	"sync"
	"template_backend/core/config"

	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v82"
)

var (
	once sync.Once
)

// InitStripe initializes the Stripe configuration
func InitStripe() {
	once.Do(func() {
		apiKey := config.GlobalConfig.Stripe.SecretKey
		if apiKey == "" {
			panic("STRIPE_SECRET_KEY environment variable is not set")
		}

		// Set Stripe API key
		stripe.Key = apiKey

		// Set the version of the Stripe API
		stripe.SetAppInfo(&stripe.AppInfo{
			Name:    "MawidaApp",
			Version: "1.0.0",
		})
		log.Info().Msg("Stripe initialized")
	})
}
