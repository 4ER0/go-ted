package handler

import (
	"fmt"
	"github.com/4ER0/go-ted/estatestics/model"
	"strings"
)

func doCalcNoParallel(es []model.RealEstate, key string) (emptyResult, rentedResult) {
	empty := calculateEmptyPercentageNoParallel(es, key)
	rented := calculateRentedPercentageNoParallel(es, key)

	return empty, rented
}

func calculateRentedPercentageNoParallel(es []model.RealEstate, key string) rentedResult {
	var sum float64
	var rented float64
	var count float64
	for _, e := range es {
		if se := fmt.Sprintf("%08d", e.Key); strings.HasPrefix(se, key) {
			sum = sum + float64(e.Total)
			rented = rented + float64(e.Rented)
			count++
		}
	}

	avgSum := sum / count
	avgRented := rented / count
	var percentage float64
	if avgSum > 0 {
		percentage = (avgRented * 100.0) / avgSum
	}

	return rentedResult{
		AverageRentedPercentage: percentage,
		TotalRented:             int(rented),
	}
}

func calculateEmptyPercentageNoParallel(es []model.RealEstate, key string) emptyResult {
	var sum float64
	var empty float64
	var count float64
	for _, e := range es {
		if se := fmt.Sprintf("%08d", e.Key); strings.HasPrefix(se, key) {
			sum = sum + float64(e.Total)
			empty = empty + float64(e.Empty)
			count++
		}
	}

	avgSum := sum / count
	avgEmpty := empty / count
	var percentage float64
	if avgSum > 0 {
		percentage = (avgEmpty * 100.0) / avgSum
	}

	return emptyResult{
		AverageEmptyPercentage: percentage,
		RecordCount:            int(count),
		TotalEstates:           int(sum),
		TotalEmpty:             int(empty),
	}
}
