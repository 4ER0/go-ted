package handler

import (
	"encoding/json"
	"fmt"
	"github.com/4ER0/go-ted/estatestics/model"
	"github.com/4ER0/go-ted/estatestics/port"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

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
	if key == "" {
		http.Error(w, "Missing key parameter", http.StatusBadRequest)
		h.l.Warnf("Missing key parameter")
		return
	}

	es, err := h.p.GetRealEstateInfo()
	if err != nil {
		http.Error(w, "Error while getting real estate info", http.StatusInternalServerError)
		h.l.WithError(err).Errorf("Error while getting real estate info")
		return
	}

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

	res := model.AverageResponse{
		AveragePercentage: percentage,
		RecordCount:       int(count),
		TotalEstates:      int(sum),
		TotalEmpty:        int(empty),
		Key:               key,
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

func (h *RealEstateInfoHandler) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	http.Error(w, "Unknown URL", http.StatusNotFound)
	h.l.WithField("url", r.URL).Warnf("Unknown URL invoked")
	return
}