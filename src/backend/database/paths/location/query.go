package database_location

import (
	"context"
	database_common "template_backend/database/common"

	"github.com/rs/zerolog/log"
)

func createFindLocationQuery(id *string, fields []interface{}) database_common.Query {
	query := database_common.Query{
		Selector: map[string]interface{}{},
		Fields:   fields,
		Limit:    1,
	}

	if id != nil {
		query.Selector["_id"] = map[string]interface{}{"$eq": *id}
	}

	return query
}

func FindLocationById(ctx context.Context, id string) *Location {
	if DatabaseLocation == nil {
		return nil
	}

	query := createFindLocationQuery(&id, []interface{}{"_id", "city", "street", "postalCode", "buildingName", "longitude", "latitude", "notes"})
	rows, err := DatabaseLocation.Find(ctx, query)
	if err != nil {
		return nil
	}

	var location Location
	if rows.Next() {
		if err := rows.ScanDoc(&location); err != nil {
			return nil
		}
		return &location
	}

	return nil
}

func GetAllLocations(ctx context.Context) []Location {
	if DatabaseLocation == nil {
		return nil
	}

	query := database_common.Query{
		Fields:   []interface{}{"_id", "city", "street", "postalCode", "buildingName", "longitude", "latitude", "notes"},
		Selector: map[string]interface{}{},
		Limit:    100, // Reasonable limit for locations
	}

	rows, err := DatabaseLocation.Find(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to scan location document")
		return nil
	}
	defer rows.Close()

	var locations []Location
	for rows.Next() {
		var location Location
		if err := rows.ScanDoc(&location); err != nil {
			log.Error().Err(err).Msg("Failed to scan location document")
			continue
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error while iterating through locations")
		return nil
	}

	return locations
}
