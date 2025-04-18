package payment

type ProductIdentifier struct {
	ID        string `json:"id"`
	Mode      string `json:"mode"`
	PriceID   string `json:"priceId"`
	DepositID string `json:"depositId"`
}
