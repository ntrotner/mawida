package database_payments

import (
	"context"

	_ "github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
	"github.com/rs/zerolog/log"
)

const PAYMENTS_DB = "payments"

var DatabasePayments *kivik.DB

func createIndex(ctx context.Context) error {
	indexId := map[string]interface{}{
		"fields": []string{"_id"},
	}

	indexCustomer := map[string]interface{}{
		"fields": []string{"customerId"},
	}

	indexProduct := map[string]interface{}{
		"fields": []string{"productId"},
	}

	err := DatabasePayments.CreateIndex(ctx, "id-index", "json", indexId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create payment id index")
		return err
	}

	err = DatabasePayments.CreateIndex(ctx, "customer-index", "json", indexCustomer)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create payment customer index")
		return err
	}

	err = DatabasePayments.CreateIndex(ctx, "product-index", "json", indexProduct)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create payment product index")
		return err
	}

	log.Info().
		Str("db", DatabasePayments.Name()).
		Msg("Payments database indexes created")
	return nil
}

func SetupPayments(ctx context.Context, databases map[string]*kivik.DB) {
	DatabasePayments = databases[PAYMENTS_DB]
	createIndex(ctx)
}
