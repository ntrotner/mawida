package database_product

// CalculateTotalAmount calculates the total amount for a rental based on:
// - Product's daily price
// - Number of rental days
// - Product's deposit amount (if any)
func CalculateTotalAmount(product *Product, days int) int64 {
	totalAmount := product.Pricing.Price * int64(days)
	if product.Pricing.Deposit > 0 {
		totalAmount += product.Pricing.Deposit
	}
	return totalAmount
}
