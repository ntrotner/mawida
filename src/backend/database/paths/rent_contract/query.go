package database_rent_contract

import (
	"context"
	database_common "template_backend/database/common"

	"github.com/go-kivik/kivik"
	"github.com/rs/zerolog/log"
)

type Database interface {
	Put(ctx context.Context, id string, doc interface{}, options ...kivik.Options) (string, error)
	Get(ctx context.Context, id string, doc interface{}, options ...kivik.Options) error
	Delete(ctx context.Context, id string, rev string, options ...kivik.Options) (string, error)
	Find(ctx context.Context, query interface{}) (*kivik.Rows, error)
	CreateIndex(ctx context.Context, ddoc, name string, index interface{}) error
	Name() string
}

func createFindRentContractQuery(id *string, fields []interface{}) database_common.Query {
	query := database_common.Query{
		Selector: map[string]interface{}{},
		Fields: []interface{}{
			"_id",
			"_rev",
			"productId",
			"userId",
			"rentalStartDate",
			"rentalEndDate",
			"status",
			"price",
			"deposit",
			"totalAmount",
			"paymentMethodId",
			"pickupLocationId",
			"returnLocationId",
			"paymentTransactionId",
			"additionalNotes",
			"dynamicAttributes",
			"createdAt",
			"updatedAt",
			"paymentIdentifier",
			"paymentInstruction",
		},
		Limit: 1,
	}

	if id != nil {
		query.Selector["_id"] = map[string]interface{}{"$eq": *id}
	}

	return query
}

func FindRentContractById(ctx context.Context, id string) *RentContract {
	if DatabaseRentContract == nil {
		log.Error().Msg("Database connection not initialized")
		return nil
	}

	query := createFindRentContractQuery(&id, nil)
	rows, err := DatabaseRentContract.Find(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve rent contract")
		return nil
	}
	defer rows.Close()

	var contract RentContract
	if rows.Next() {
		if err := rows.ScanDoc(&contract); err != nil {
			log.Error().Err(err).Msg("Failed to scan rent contract document")
			return nil
		}
		return &contract
	}

	return nil
}

func FindRentContractByCheckoutSessionId(ctx context.Context, checkoutSessionId string) *RentContract {
	if DatabaseRentContract == nil {
		log.Error().Msg("Database connection not initialized")
		return nil
	}

	query := createFindRentContractQuery(nil, []interface{}{
		"paymentIdentifier",
	})
	query.Selector["paymentIdentifier"] = map[string]interface{}{"$eq": checkoutSessionId}

	rows, err := DatabaseRentContract.Find(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve rent contract")
		return nil
	}
	defer rows.Close()

	var contract RentContract
	if rows.Next() {
		if err := rows.ScanDoc(&contract); err != nil {
			log.Error().Err(err).Msg("Failed to scan rent contract document")
			return nil
		}
		return &contract
	}

	return nil
}

func FindRentContractsByUserId(ctx context.Context, userId string) []*RentContract {
	if DatabaseRentContract == nil {
		log.Error().Msg("Database connection not initialized")
		return nil
	}

	query := database_common.Query{
		Selector: map[string]interface{}{
			"userId": map[string]interface{}{"$eq": userId},
		},
		Fields: []interface{}{
			"_id",
			"_rev",
			"productId",
			"userId",
			"rentalStartDate",
			"rentalEndDate",
			"status",
			"price",
			"deposit",
			"totalAmount",
			"paymentMethodId",
			"pickupLocationId",
			"returnLocationId",
			"additionalNotes",
			"dynamicAttributes",
			"createdAt",
			"updatedAt",
			"paymentIdentifier",
			"paymentInstruction",
		},
	}

	rows, err := DatabaseRentContract.Find(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve rent contracts")
		return nil
	}
	defer rows.Close()

	var contracts []*RentContract
	for rows.Next() {
		var contract RentContract
		if err := rows.ScanDoc(&contract); err != nil {
			log.Error().Err(err).Msg("Failed to scan rent contract document")
			continue
		}
		contracts = append(contracts, &contract)
	}

	return contracts
}

func FindRentContractsByProductId(ctx context.Context, productId string) []*RentContract {
	if DatabaseRentContract == nil {
		log.Error().Msg("Database connection not initialized")
		return nil
	}

	query := database_common.Query{
		Selector: map[string]interface{}{
			"productId": map[string]interface{}{"$eq": productId},
		},
		Fields: []interface{}{
			"_id",
			"_rev",
			"productId",
			"userId",
			"rentalStartDate",
			"rentalEndDate",
			"status",
			"price",
			"deposit",
			"totalAmount",
			"paymentMethodId",
			"pickupLocationId",
			"returnLocationId",
			"additionalNotes",
			"dynamicAttributes",
			"createdAt",
			"updatedAt",
			"paymentIdentifier",
			"paymentInstruction",
		},
	}

	rows, err := DatabaseRentContract.Find(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve rent contracts")
		return nil
	}
	defer rows.Close()

	var contracts []*RentContract
	for rows.Next() {
		var contract RentContract
		if err := rows.ScanDoc(&contract); err != nil {
			log.Error().Err(err).Msg("Failed to scan rent contract document")
			continue
		}
		contracts = append(contracts, &contract)
	}

	return contracts
}

func FindAllRentContracts(ctx context.Context) []*RentContract {
	if DatabaseRentContract == nil {
		log.Error().Msg("Database connection not initialized")
		return nil
	}

	query := database_common.Query{
		Selector: map[string]interface{}{},
		Fields: []interface{}{
			"_id",
			"_rev",
			"productId",
			"userId",
			"rentalStartDate",
			"rentalEndDate",
			"status",
			"price",
			"deposit",
			"totalAmount",
			"paymentMethodId",
			"pickupLocationId",
			"returnLocationId",
			"additionalNotes",
			"dynamicAttributes",
			"createdAt",
			"updatedAt",
			"paymentIdentifier",
			"paymentInstruction",
		},
	}

	rows, err := DatabaseRentContract.Find(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve rent contracts")
		return nil
	}
	defer rows.Close()

	var contracts []*RentContract
	for rows.Next() {
		var contract RentContract
		if err := rows.ScanDoc(&contract); err != nil {
			log.Error().Err(err).Msg("Failed to scan rent contract document")
			continue
		}
		contracts = append(contracts, &contract)
	}

	return contracts
}

func FindRentContractByPaymentTransactionId(ctx context.Context, paymentTransactionId string) *RentContract {
	if DatabaseRentContract == nil {
		log.Error().Msg("Database connection not initialized")
		return nil
	}

	query := database_common.Query{
		Selector: map[string]interface{}{
			"paymentTransactionId": map[string]interface{}{"$eq": paymentTransactionId},
		},
		Fields: []interface{}{
			"_id",
			"_rev",
			"productId",
			"userId",
			"rentalStartDate",
			"rentalEndDate",
			"status",
			"price",
			"deposit",
			"totalAmount",
			"paymentMethodId",
			"pickupLocationId",
			"returnLocationId",
			"additionalNotes",
			"dynamicAttributes",
			"createdAt",
			"updatedAt",
			"paymentIdentifier",
			"paymentInstruction",
		},
	}

	rows, err := DatabaseRentContract.Find(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve rent contract")
		return nil
	}
	defer rows.Close()

	var contract RentContract
	if rows.Next() {
		if err := rows.ScanDoc(&contract); err != nil {
			log.Error().Err(err).Msg("Failed to scan rent contract document")
			return nil
		}
		return &contract
	}

	return nil
}

func FindRentContractByPaymentIdentifier(ctx context.Context, paymentIdentifier string) *RentContract {
	if DatabaseRentContract == nil {
		log.Error().Msg("Database connection not initialized")
		return nil
	}

	query := database_common.Query{
		Selector: map[string]interface{}{
			"paymentIdentifier.ID": map[string]interface{}{"$eq": paymentIdentifier},
		},
		Fields: []interface{}{
			"_id",
			"_rev",
			"productId",
			"userId",
			"rentalStartDate",
			"rentalEndDate",
			"status",
			"price",
			"deposit",
			"totalAmount",
			"paymentMethodId",
			"pickupLocationId",
			"returnLocationId",
			"additionalNotes",
			"dynamicAttributes",
			"createdAt",
			"updatedAt",
			"paymentIdentifier",
			"paymentInstruction",
		},
	}

	rows, err := DatabaseRentContract.Find(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve rent contract")
		return nil
	}
	defer rows.Close()

	var contract RentContract
	if rows.Next() {
		if err := rows.ScanDoc(&contract); err != nil {
			log.Error().Err(err).Msg("Failed to scan rent contract document")
			return nil
		}
		return &contract
	}

	return nil
}
