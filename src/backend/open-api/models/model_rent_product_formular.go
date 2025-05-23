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

import "template_backend/infrastructure/validators"

type RentProductFormular struct {

	// The start date of the rental period
	RentalStartDate int64 `json:"rentalStartDate" validate:"required,gte=0"`

	// The end date of the rental period
	RentalEndDate int64 `json:"rentalEndDate" validate:"required,gtefield=RentalStartDate"`

	// Free text for additional notes or special instructions
	AdditionalNotes string `json:"additionalNotes,omitempty" validate:"min=0,max=1000"`

	// Identifier for the location where the product will be picked up
	LocationId string `json:"locationId" validate:"required"`

	// Identifier for the payment method
	PaymentMethodId PaymentMethod `json:"paymentMethodId" validate:"required,oneof=cash stripe"`

	// Dynamic attributes for the rental request
	DynamicAttributes map[string]interface{} `json:"dynamicAttributes,omitempty" validate:"dynamicAttributes"`
}

// AssertRentProductFormularRequired checks if the required fields are not zero-ed
func AssertRentProductFormularRequired(obj RentProductFormular) error {
	elements := map[string]interface{}{
		"rentalStartDate": obj.RentalStartDate,
		"rentalEndDate":   obj.RentalEndDate,
		"locationId":      obj.LocationId,
		"paymentMethodId": obj.PaymentMethodId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRentProductFormularConstraints checks if the values respects the defined constraints
func AssertRentProductFormularConstraints(obj RentProductFormular) error {
	return validators.Validate.Struct(obj)
}
