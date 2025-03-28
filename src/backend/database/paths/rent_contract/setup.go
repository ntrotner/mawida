package database_rent_contract

import (
	"context"

	_ "github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
	"github.com/rs/zerolog/log"
)

const RENT_CONTRACT_DB = "rent_contracts"

var DatabaseRentContract *kivik.DB

func createIndex(ctx context.Context) error {
	indexId := map[string]interface{}{
		"fields": []string{"_id"},
	}

	indexProductId := map[string]interface{}{
		"fields": []string{"productId"},
	}

	indexUserId := map[string]interface{}{
		"fields": []string{"userId"},
	}

	indexStatus := map[string]interface{}{
		"fields": []string{"status"},
	}

	err := DatabaseRentContract.CreateIndex(ctx, "id-index", "json", indexId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create rent contract id index")
		return err
	}

	err = DatabaseRentContract.CreateIndex(ctx, "product-id-index", "json", indexProductId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create rent contract product id index")
		return err
	}

	err = DatabaseRentContract.CreateIndex(ctx, "user-id-index", "json", indexUserId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create rent contract user id index")
		return err
	}

	err = DatabaseRentContract.CreateIndex(ctx, "status-index", "json", indexStatus)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create rent contract status index")
		return err
	}

	log.Info().
		Str("db", DatabaseRentContract.Name()).
		Msg("Rent contract database indexes created")
	return nil
}

func SetupRentContract(ctx context.Context, databases map[string]*kivik.DB) {
	DatabaseRentContract = databases[RENT_CONTRACT_DB]
	createIndex(ctx)
}
