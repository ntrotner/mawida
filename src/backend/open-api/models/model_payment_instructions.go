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




type PaymentInstructions struct {

	PaymentMethodId PaymentMethod `json:"paymentMethodId"`

	// Dynamic attributes for the payment instructions
	DynamicAttributes map[string]interface{} `json:"dynamicAttributes,omitempty"`
}

// AssertPaymentInstructionsRequired checks if the required fields are not zero-ed
func AssertPaymentInstructionsRequired(obj PaymentInstructions) error {
	elements := map[string]interface{}{
		"paymentMethodId": obj.PaymentMethodId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPaymentInstructionsConstraints checks if the values respects the defined constraints
func AssertPaymentInstructionsConstraints(obj PaymentInstructions) error {
	return nil
}
