package database_payments

// PaymentTransaction represents a stored record of a payment event, typically from a webhook.
type PaymentTransaction struct {
	ID                    string `json:"_id,omitempty"`
	Rev                   string `json:"_rev,omitempty"`
	CustomerID            string `json:"customerId,omitempty"`
	ProductID             string `json:"productId,omitempty"`
	ProductPrice          int64  `json:"productPrice,omitempty"`
	ProductDeposit        int64  `json:"productDeposit,omitempty"`
	ProductMode           string `json:"productMode,omitempty"`
	Paid                  bool   `json:"paid,omitempty"`
	CreatedAt             int64  `json:"createdAt,omitempty"`
	PaidAt                int64  `json:"paidAt,omitempty"`
	ReturnedAt            int64  `json:"returnedAt,omitempty"`
	DepositRefundedAmount int64  `json:"depositRefundedAmount,omitempty"`
	DepositRefundedAt     int64  `json:"depositRefundedAt,omitempty"`
}
