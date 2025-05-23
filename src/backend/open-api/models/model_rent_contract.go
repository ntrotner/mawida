/*
 * Swagger - OpenAPI 3.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Contact: nikita@ttnr.me
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type RentContract struct {

	// Unique identifier for the rent contract
	Id string `json:"id,omitempty"`

	// Unique identifier for the product
	ProductId string `json:"productId,omitempty"`

	// Unique identifier for the user
	UserId string `json:"userId,omitempty"`

	PaymentInstructions PaymentInstructions `json:"paymentInstructions,omitempty"`

	// Unix timestamp for the start date of the rental period
	RentalStartDate int64 `json:"rentalStartDate,omitempty"`

	// Unix timestamp for the end date of the rental period
	RentalEndDate int64 `json:"rentalEndDate,omitempty"`

	Status RentContractStatus `json:"status,omitempty"`

	// Price of the rent contract
	Price int64 `json:"price,omitempty"`

	// Deposit amount for the rent contract
	Deposit int64 `json:"deposit,omitempty"`

	// Total amount of the rent contract
	TotalAmount int64 `json:"totalAmount,omitempty"`

	// Unique identifier for the pickup location
	PickupLocationId string `json:"pickupLocationId,omitempty"`

	// Unique identifier for the return location
	ReturnLocationId string `json:"returnLocationId,omitempty"`
}

// AssertRentContractRequired checks if the required fields are not zero-ed
func AssertRentContractRequired(obj RentContract) error {
	if err := AssertPaymentInstructionsRequired(obj.PaymentInstructions); err != nil {
		return err
	}
	return nil
}

// AssertRentContractConstraints checks if the values respects the defined constraints
func AssertRentContractConstraints(obj RentContract) error {
	return nil
}
