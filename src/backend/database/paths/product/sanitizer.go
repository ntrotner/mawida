package database_product

func SanitizeProduct(product *Product) *Product {
	product.DynamicAttributes = nil
	product.RenterInfo = RenterInfo{}
	return product
}
