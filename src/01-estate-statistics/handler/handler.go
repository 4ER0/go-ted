package handler

import (
	"encoding/json"
	"fmt"
	"github.com/4ER0/go-ted/estatestics/model"
	"github.com/4ER0/go-ted/estatestics/port"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type rentedResult struct {
	AverageRentedPercentage float64
	TotalRented             int
}

type emptyResult struct {
	AverageEmptyPercentage float64
	RecordCount            int
	TotalEstates           int
	TotalEmpty             int
}

type RealEstateInfoHandler struct {
	l *logrus.Logger
	p *port.RealEstatePort
}

func NewRealEstateInfoHandler(l *logrus.Logger, p *port.RealEstatePort) *RealEstateInfoHandler {
	return &RealEstateInfoHandler{l, p}
}

// GetAverageEmptyEstates returns the average number of empty estates for given key
func (h *RealEstateInfoHandler) GetAverageEmptyEstates(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	es, err := h.p.GetRealEstateInfo()
	if err != nil {
		http.Error(w, "Error while getting real estate info", http.StatusInternalServerError)
		h.l.WithError(err).Errorf("Error while getting real estate info")
		return
	}

	empty, rented := doCalc(es, key)

	res := model.AverageResponse{
		AverageEmptyPercentage:  empty.AverageEmptyPercentage,
		AverageRentedPercentage: rented.AverageRentedPercentage,
		RecordCount:             empty.RecordCount,
		TotalEstates:            empty.TotalEstates,
		TotalEmpty:              empty.TotalEmpty,
		TotalRented:             rented.TotalRented,
		Key:                     key,
	}

	bs, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		h.l.WithError(err).Errorf("Error marshalling response")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bs)
}

func doCalc(es []model.RealEstate, key string) (emptyResult, rentedResult) {
	cRented := make(chan rentedResult)
	cEmpty := make(chan emptyResult)
	go calculateEmptyPercentage(es, key, cEmpty)
	go calculateRentedPercentage(es, key, cRented)

	empty := <-cEmpty
	rented := <-cRented
	return empty, rented
}

func calculateRentedPercentage(es []model.RealEstate, key string, c chan rentedResult) {
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

	c <- rentedResult{
		AverageRentedPercentage: percentage,
		TotalRented:             int(rented),
	}
}

func calculateEmptyPercentage(es []model.RealEstate, key string, c chan emptyResult) {
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

	c <- emptyResult{
		AverageEmptyPercentage: percentage,
		RecordCount:            int(count),
		TotalEstates:           int(sum),
		TotalEmpty:             int(empty),
	}
}

func (h *RealEstateInfoHandler) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Unknown URL", http.StatusNotFound)
	h.l.WithField("url", r.URL).Warnf("Unknown URL invoked")
	return
}
