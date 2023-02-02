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
	mux.HandleFunc("/empty/average", h.GetAverageEmptyEstates)
	mux.HandleFunc("/", h.DefaultHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		l.WithError(err).Fatal("Error while starting server")
	}
}
