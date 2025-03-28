package database_product

type Product struct {
	ID                string                 `json:"_id,omitempty"`
	Rev               string                 `json:"_rev,omitempty"`
	Name              string                 `json:"name"`
	Description       string                 `json:"description"`
	Images            []ProductImage         `json:"images,omitempty"`
	Documents         []ProductDocument      `json:"documents,omitempty"`
	Location          string                 `json:"location"`
	Pricing           ProductPricing         `json:"pricing"`
	DynamicAttributes map[string]interface{} `json:"dynamicAttributes,omitempty"`
	IsRented          bool                   `json:"isRented,omitempty"`
	RenterInfo        *RenterInfo            `json:"renterInfo,omitempty"`
}

type ProductImage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Data string `json:"data"`
}

type ProductDocument struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Data string `json:"data"`
}

type ProductPricing struct {
	Price   float32 `json:"price"`
	Deposit float32 `json:"deposit,omitempty"`
}

type RenterInfo struct {
	UserID          string `json:"userId"`
	RentalStartDate string `json:"rentalStartDate"`
	RentalEndDate   string `json:"rentalEndDate"`
	Status          string `json:"status"`
}
