package database_location

import (
	"context"
)

func CreateLocation(ctx context.Context, id string, city string, street string, postalCode string, buildingName string, latitude float64, longitude float64, notes string) error {
	if DatabaseLocation == nil {
		return nil
	}

	location := &Location{
		ID:           id,
		City:         city,
		Street:       street,
		PostalCode:   postalCode,
		BuildingName: buildingName,
		Latitude:     latitude,
		Longitude:    longitude,
		Notes:        notes,
	}

	_, err := DatabaseLocation.Put(ctx, location.ID, location)
	return err
}

func UpdateLocation(ctx context.Context, location *Location) error {
	if DatabaseLocation == nil {
		return nil
	}

	_, err := DatabaseLocation.Put(ctx, location.ID, location)
	return err
}

func DeleteLocation(ctx context.Context, id string, rev string) error {
	if DatabaseLocation == nil {
		return nil
	}

	_, err := DatabaseLocation.Delete(ctx, id, rev)
	return err
}
