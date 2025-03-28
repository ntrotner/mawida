package database_rent_contract

import (
	"time"
)

type RentContractStatus string

const (
	RentContractStatusPending   RentContractStatus = "pending"
	RentContractStatusConfirmed RentContractStatus = "confirmed"
	RentContractStatusActive    RentContractStatus = "active"
	RentContractStatusCompleted RentContractStatus = "completed"
	RentContractStatusCancelled RentContractStatus = "cancelled"
	RentContractStatusOverdue   RentContractStatus = "overdue"
	RentContractStatusDisputed  RentContractStatus = "disputed"
)

type RentContract struct {
	ID                string                 `json:"_id,omitempty"`
	Rev               string                 `json:"_rev,omitempty"`
	ProductID         string                 `json:"productId" validate:"required,uuid"`
	UserID            string                 `json:"userId" validate:"required,uuid"`
	RentalStartDate   string                 `json:"rentalStartDate" validate:"required"`
	RentalEndDate     string                 `json:"rentalEndDate" validate:"required,"`
	Status            RentContractStatus     `json:"status" validate:"required,oneof=pending confirmed active completed cancelled overdue disputed"`
	Price             float32                `json:"price" validate:"required,gte=0"`
	Deposit           float32                `json:"deposit,omitempty" validate:"omitempty,gte=0"`
	TotalAmount       float32                `json:"totalAmount" validate:"required,gte=0"`
	PaymentMethodID   string                 `json:"paymentMethodId" validate:"required,uuid"`
	PickupLocationID  string                 `json:"pickupLocationId" validate:"required,uuid"`
	ReturnLocationID  string                 `json:"returnLocationId" validate:"required,uuid"`
	AdditionalNotes   string                 `json:"additionalNotes,omitempty" validate:"omitempty,max=1000"`
	DynamicAttributes map[string]interface{} `json:"dynamicAttributes,omitempty"`
	CreatedAt         time.Time              `json:"createdAt"`
	UpdatedAt         time.Time              `json:"updatedAt"`
}

func (s RentContractStatus) IsValid() bool {
	switch s {
	case RentContractStatusPending,
		RentContractStatusConfirmed,
		RentContractStatusActive,
		RentContractStatusCompleted,
		RentContractStatusCancelled,
		RentContractStatusOverdue,
		RentContractStatusDisputed:
		return true
	default:
		return false
	}
}

func (s RentContractStatus) CanTransitionTo(newStatus RentContractStatus) bool {
	switch s {
	case RentContractStatusPending:
		return newStatus == RentContractStatusConfirmed || newStatus == RentContractStatusCancelled
	case RentContractStatusConfirmed:
		return newStatus == RentContractStatusActive || newStatus == RentContractStatusCancelled
	case RentContractStatusActive:
		return newStatus == RentContractStatusCompleted || newStatus == RentContractStatusOverdue || newStatus == RentContractStatusDisputed
	case RentContractStatusOverdue:
		return newStatus == RentContractStatusCompleted || newStatus == RentContractStatusDisputed
	case RentContractStatusDisputed:
		return newStatus == RentContractStatusCompleted || newStatus == RentContractStatusCancelled
	default:
		return false
	}
}
