package database_rent_contract

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func CreateRentContract(ctx context.Context, productId string, userId string, rentalStartDate int64, rentalEndDate int64, price float32, deposit float32, totalAmount float32, paymentMethodId string, pickupLocationId string, returnLocationId string, additionalNotes string, dynamicAttributes map[string]interface{}) *RentContract {
	if DatabaseRentContract == nil {
		return nil
	}

	contract := &RentContract{
		ID:                uuid.NewString(),
		ProductID:         productId,
		UserID:            userId,
		RentalStartDate:   rentalStartDate,
		RentalEndDate:     rentalEndDate,
		Status:            RentContractStatusPending,
		Price:             price,
		Deposit:           deposit,
		TotalAmount:       totalAmount,
		PaymentMethodID:   paymentMethodId,
		PickupLocationID:  pickupLocationId,
		ReturnLocationID:  returnLocationId,
		AdditionalNotes:   additionalNotes,
		DynamicAttributes: dynamicAttributes,
		CreatedAt:         time.Now().Unix(),
		UpdatedAt:         time.Now().Unix(),
	}

	_, err := DatabaseRentContract.Put(ctx, contract.ID, contract)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create rent contract")
		return nil
	}

	log.Info().
		Str("id", contract.ID).
		Str("productId", contract.ProductID).
		Str("userId", contract.UserID).
		Msg("Created rent contract")

	return contract
}

func UpdateRentContract(ctx context.Context, contract *RentContract) *RentContract {
	if DatabaseRentContract == nil {
		return nil
	}

	contract.UpdatedAt = time.Now().Unix()
	_, err := DatabaseRentContract.Put(ctx, contract.ID, contract)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update rent contract")
		return nil
	}

	log.Info().
		Str("id", contract.ID).
		Str("productId", contract.ProductID).
		Str("userId", contract.UserID).
		Msg("Updated rent contract")

	return contract
}

func UpdateRentContractStatus(ctx context.Context, id string, newStatus RentContractStatus) *RentContract {
	if DatabaseRentContract == nil {
		return nil
	}

	contract := FindRentContractById(ctx, id)
	if contract == nil {
		log.Error().Str("id", id).Msg("Rent contract not found")
		return nil
	}

	if !contract.Status.CanTransitionTo(newStatus) {
		log.Error().
			Str("id", id).
			Str("currentStatus", string(contract.Status)).
			Str("newStatus", string(newStatus)).
			Msg("Invalid status transition")
		return nil
	}

	contract.Status = newStatus
	contract.UpdatedAt = time.Now().Unix()

	_, err := DatabaseRentContract.Put(ctx, contract.ID, contract)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update rent contract status")
		return nil
	}

	log.Info().
		Str("id", contract.ID).
		Str("productId", contract.ProductID).
		Str("userId", contract.UserID).
		Str("status", string(contract.Status)).
		Msg("Updated rent contract status")

	return contract
}

func DeleteRentContract(ctx context.Context, id string, rev string) bool {
	if DatabaseRentContract == nil {
		return false
	}

	_, err := DatabaseRentContract.Delete(ctx, id, rev)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete rent contract")
		return false
	}

	log.Info().
		Str("id", id).
		Msg("Deleted rent contract")

	return true
}
