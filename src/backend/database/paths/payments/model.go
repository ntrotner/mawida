package database_payments

// PaymentTransaction represents a stored record of a payment event, typically from a webhook.
type PaymentTransaction struct {
	ID             string `json:"_id,omitempty"`
	Rev            string `json:"_rev,omitempty"`
	CustomerID     string `json:"customerId,omitempty"`
	ProductID      string `json:"productId,omitempty"`
	ProductPrice   int64  `json:"productPrice,omitempty"`
	ProductDeposit int64  `json:"productDeposit,omitempty"`
	ProductMode    string `json:"productMode,omitempty"`
	Paid           bool   `json:"paid,omitempty"`
}
