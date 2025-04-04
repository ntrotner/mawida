package database_product

import (
	"context"

	_ "github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
	"github.com/rs/zerolog/log"
)

const PRODUCT_DB = "products"

var DatabaseProduct *kivik.DB

func createIndex(ctx context.Context) error {
	indexId := map[string]interface{}{
		"fields": []string{"_id"},
	}

	indexLocation := map[string]interface{}{
		"fields": []string{"location"},
	}

	indexRented := map[string]interface{}{
		"fields": []string{"isRented"},
	}

	err := DatabaseProduct.CreateIndex(ctx, "id-index", "json", indexId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create product id index")
		return err
	}

	err = DatabaseProduct.CreateIndex(ctx, "location-index", "json", indexLocation)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create product location index")
		return err
	}

	err = DatabaseProduct.CreateIndex(ctx, "rented-index", "json", indexRented)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create product rented index")
		return err
	}

	log.Info().
		Str("db", DatabaseProduct.Name()).
		Msg("Product database indexes created")
	return nil
}

func SetupProduct(ctx context.Context, databases map[string]*kivik.DB) {
	DatabaseProduct = databases[PRODUCT_DB]
	createIndex(ctx)
}
