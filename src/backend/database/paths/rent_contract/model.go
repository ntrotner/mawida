package database_rent_contract

import (
	paymentTypes "template_backend/infrastructure/payment/types"
)

type RentContractStatus string

const (
	RentContractStatusPickupPending RentContractStatus = "pickupPending"
	RentContractStatusConfirmed     RentContractStatus = "confirmed"
	RentContractStatusActive        RentContractStatus = "active"
	RentContractStatusCompleted     RentContractStatus = "completed"
	RentContractStatusCancelled     RentContractStatus = "cancelled"
)

type PaymentMethod string

const (
	PaymentMethodCash   PaymentMethod = "cash"
	PaymentMethodStripe PaymentMethod = "stripe"
)

type PaymentInstruction struct {
	PaymentMethodID   PaymentMethod          `json:"paymentMethodId" validate:"required,oneof=cash stripe"`
	DynamicAttributes map[string]interface{} `json:"dynamicAttributes,omitempty"`
}

type RentContract struct {
	ID                 string                                  `json:"_id,omitempty"`
	Rev                string                                  `json:"_rev,omitempty"`
	ProductID          string                                  `json:"productId" validate:"required,uuid"`
	UserID             string                                  `json:"userId" validate:"required,uuid"`
	RentalStartDate    int64                                   `json:"rentalStartDate" validate:"required"`
	RentalEndDate      int64                                   `json:"rentalEndDate" validate:"required"`
	Status             RentContractStatus                      `json:"status" validate:"required,oneof=pending confirmed active completed cancelled overdue disputed"`
	Price              int64                                   `json:"price" validate:"required,gte=0"`
	Deposit            int64                                   `json:"deposit,omitempty" validate:"required,gte=0"`
	TotalAmount        int64                                   `json:"totalAmount" validate:"required,gte=0"`
	PaymentMethodID    string                                  `json:"paymentMethodId" validate:"required,uuid"`
	PickupLocationID   string                                  `json:"pickupLocationId" validate:"required,uuid"`
	ReturnLocationID   string                                  `json:"returnLocationId" validate:"required,uuid"`
	AdditionalNotes    string                                  `json:"additionalNotes,omitempty" validate:"omitempty,max=1000"`
	DynamicAttributes  map[string]interface{}                  `json:"dynamicAttributes,omitempty"`
	CreatedAt          int64                                   `json:"createdAt" validate:"required"`
	UpdatedAt          int64                                   `json:"updatedAt" validate:"required"`
	PaymentIdentifier  *paymentTypes.CheckoutSessionIdentifier `json:"paymentIdentifier,omitempty"`
	PaymentInstruction *PaymentInstruction                     `json:"paymentInstruction,omitempty"`
}
