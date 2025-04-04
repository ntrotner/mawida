package database_product

import (
	"context"
	database_common "template_backend/database/common"

	"github.com/rs/zerolog/log"
)

func createFindProductQuery(id *string, fields []interface{}) database_common.Query {
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

func FindProductById(ctx context.Context, id string) *Product {
	if DatabaseProduct == nil {
		return nil
	}

	query := createFindProductQuery(&id, []interface{}{
		"_id",
		"_rev",
		"name",
		"description",
		"images",
		"documents",
		"location",
		"pricing",
		"dynamicAttributes",
		"isRented",
		"renterInfo",
	})
	rows, err := DatabaseProduct.Find(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve product")
		return nil
	}
	defer rows.Close()

	var product Product
	if rows.Next() {
		if err := rows.ScanDoc(&product); err != nil {
			log.Error().Err(err).Msg("Failed to scan product document")
			return nil
		}
		return &product
	}

	return nil
}

func GetAllProducts(ctx context.Context) []Product {
	if DatabaseProduct == nil {
		return nil
	}

	query := database_common.Query{
		Selector: map[string]interface{}{},
		Limit:    100, // Reasonable limit for products,
		Fields: []interface{}{
			"_id",
			"_rev",
			"name",
			"description",
			"images",
			"documents",
			"location",
			"pricing",
			"dynamicAttributes",
			"isRented",
			"renterInfo",
		},
	}

	rows, err := DatabaseProduct.Find(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve products")
		return nil
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.ScanDoc(&product); err != nil {
			log.Error().Err(err).Msg("Failed to scan product document")
			continue
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error while iterating through products")
		return nil
	}

	return products
}

func GetProductsByLocation(ctx context.Context, locationId string) []Product {
	if DatabaseProduct == nil {
		return nil
	}

	query := database_common.Query{
		Selector: map[string]interface{}{
			"location": map[string]interface{}{"$eq": locationId},
		},
		Limit: 100,
	}

	rows, err := DatabaseProduct.Find(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve products by location")
		return nil
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.ScanDoc(&product); err != nil {
			log.Error().Err(err).Msg("Failed to scan product document")
			continue
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error while iterating through products")
		return nil
	}

	return products
}
