package database_location

import (
	"context"

	_ "github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
	"github.com/rs/zerolog/log"
)

const LOCATION_DB = "locations"

var DatabaseLocation *kivik.DB

func createIndex(ctx context.Context) {
	indexId := map[string]interface{}{
		"fields": []string{"_id"},
	}

	err := DatabaseLocation.CreateIndex(ctx, "id-index", "json", indexId)
	if err != nil {
		log.Error().Msg(err.Error())
	} else {
		log.Info().
			Str("index", "id-index").
			Str("db", DatabaseLocation.Name()).
			Msg("Index created")
	}
}

func SetupLocation(ctx context.Context, databases map[string]*kivik.DB) {
	DatabaseLocation = databases[LOCATION_DB]
	createIndex(ctx)
}
