package database_product

func SanitizeProduct(product *Product) *Product {
	product.RenterInfo = RenterInfo{}
	return product
}
