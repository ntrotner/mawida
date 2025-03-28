package database_location

type Location struct {
	ID           string  `json:"_id,omitempty"`
	Rev          string  `json:"_rev,omitempty"`
	City         string  `json:"city,omitempty"`
	Street       string  `json:"street,omitempty"`
	PostalCode   string  `json:"postalCode,omitempty"`
	BuildingName string  `json:"buildingName,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Notes        string  `json:"notes,omitempty"`
}

type PublicLocation struct {
	ID           string  `json:"_id,omitempty"`
	City         string  `json:"city,omitempty"`
	Street       string  `json:"street,omitempty"`
	PostalCode   string  `json:"postalCode,omitempty"`
	BuildingName string  `json:"buildingName,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Notes        string  `json:"notes,omitempty"`
}
