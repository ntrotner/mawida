package database_product

import (
	"context"
	paymentTypes "template_backend/infrastructure/payment/types"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func CreateProduct(ctx context.Context, name string, description string, location string, pricing ProductPricing, images []ProductImage, documents []ProductDocument, productIdentifier *paymentTypes.ProductIdentifier, dynamicAttributes map[string]interface{}) *Product {
	if DatabaseProduct == nil {
		return nil
	}

	product := &Product{
		ID:                uuid.NewString(),
		Name:              name,
		Description:       description,
		Location:          location,
		Pricing:           pricing,
		Images:            images,
		Documents:         documents,
		DynamicAttributes: dynamicAttributes,
		IsRented:          false,
		RenterInfo:        RenterInfo{},
		ProductIdentifier: productIdentifier,
	}

	_, err := DatabaseProduct.Put(ctx, product.ID, product)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create product")
		return nil
	}

	log.Info().
		Str("id", product.ID).
		Str("name", product.Name).
		Msg("Created product")

	return product
}

func UpdateProduct(ctx context.Context, product *Product) *Product {
	if DatabaseProduct == nil {
		return nil
	}

	_, err := DatabaseProduct.Put(ctx, product.ID, product)
	if err != nil {
		return nil
	}

	log.Info().
		Str("id", product.ID).
		Str("name", product.Name).
		Msg("Updated product")

	return product
}

func DeleteProduct(ctx context.Context, id string, rev string) bool {
	if DatabaseProduct == nil {
		return false
	}

	_, err := DatabaseProduct.Delete(ctx, id, rev)
	if err != nil {
		return false
	}

	log.Info().
		Str("id", id).
		Msg("Deleted product")

	return true
}
