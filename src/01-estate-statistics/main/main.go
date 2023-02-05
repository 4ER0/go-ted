package main

import (
	"github.com/4ER0/go-ted/estatestics/handler"
	"github.com/4ER0/go-ted/estatestics/port"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	l := logrus.New()
	l.SetLevel(logrus.DebugLevel)
	c := &http.Client{}
	mux := http.NewServeMux()
	p := port.NewRealEstatePort(l, c)
	h := handler.NewRealEstateInfoHandler(l, p)
	mux.HandleFunc("/average", checkGetRequest(h.GetAverageEmptyEstates))
	mux.HandleFunc("/", h.DefaultHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		l.WithError(err).Fatal("Error while starting server")
	}
}

func checkGetRequest(handler func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}
