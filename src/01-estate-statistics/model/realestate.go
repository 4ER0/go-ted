package model

type RealEstate struct {
	Key         int    `json:"schlüssel"`
	City        string `json:"gemeinde"`
	Total       int    `json:"gesamt"`
	OwnedLiving int    `json:"wohneigentum_bewohnt"`
	Rented      int    `json:"vermietet"`
	SecondHome  int    `json:"ferienwohnung"`
	Empty       int    `json:"leehrstand"`
}

type AverageResponse struct {
	AverageRentedPercentage float64 `json:"averageRentedPercentage"`
	AverageEmptyPercentage  float64 `json:"averageEmptyPercentage"`
	RecordCount             int     `json:"totalRecords"`
	TotalEstates            int     `json:"totalEstates"`
	TotalEmpty              int     `json:"totalEmpty"`
	TotalRented             int     `json:"totalRented"`
	Key                     string  `json:"key"`
}
