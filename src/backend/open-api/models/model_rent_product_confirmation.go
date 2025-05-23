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

type RentProductConfirmation struct {

	// Unique identifier for the rent contract
	RentContractId string `json:"rentContractId,omitempty"`
}

// AssertRentProductConfirmationRequired checks if the required fields are not zero-ed
func AssertRentProductConfirmationRequired(obj RentProductConfirmation) error {
	return nil
}

// AssertRentProductConfirmationConstraints checks if the values respects the defined constraints
func AssertRentProductConfirmationConstraints(obj RentProductConfirmation) error {
	return nil
}
