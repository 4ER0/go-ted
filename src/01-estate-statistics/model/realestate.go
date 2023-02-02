package model

type RealEstate struct {
	Key         int    `json:"schlüssel"`
	City        string `json:"gemeinde"`
	Total       int    `json:"gesamt"`
	OwnedLiving int    `json:"wohneigentum_bewohnt"`
	Rented      int    `json:"vermeitet"`
	SecondHome  int    `json:"ferienwohnung"`
	Empty       int    `json:"leehrstand"`
}

type AverageResponse struct {
	AveragePercentage float64 `json:"averagePercentage"`
	RecordCount       int     `json:"total"`
	TotalEstates      int     `json:"totalEstates"`
	TotalEmpty        int     `json:"totalEmpty"`
	Key               string  `json:"key"`
}
