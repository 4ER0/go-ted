package handler

import (
	"bytes"
	"github.com/4ER0/go-ted/estatestics/port"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type MockRoundTripper struct {
	calledUrl string
	response  *http.Response
	error     error
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	m.calledUrl = req.URL.String()
	return m.response, m.error
}

// Test default handler should return error
func TestDefaultHandler(t *testing.T) {
	file, _ := os.ReadFile("./data/real-estate-mock-data.json")
	mr := &MockRoundTripper{
		response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(file)),
		},
		error: nil,
	}
	c := &http.Client{
		Transport: mr,
	}
	l := logrus.New()
	l.SetLevel(logrus.DebugLevel)
	p := port.NewRealEstatePort(l, c)
	h := NewRealEstateInfoHandler(l, p)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.DefaultHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
	if "Unknown URL\n" != (rr.Body.String()) {
		t.Errorf("handler returned unexpected body: got %q want %q",
			rr.Body.String(), "Unknown URL\n")
	}
	if mr.calledUrl != "https://manualmk.blob.core.windows.net/goted/wohnungen-deutschland.json?sp=r&st=2023-02-02T13:19:48Z&se=2023-02-28T21:19:48Z&spr=https&sv=2021-06-08&sr=b&sig=rZRd03YH3AVwJPusaI10nAkAX5FOtSyKlPVd7T9Tqh0%3D" {
		t.Errorf("handler called wrong url: got %q want %q",
			mr.calledUrl, "https://manualmk.blob.core.windows.net/goted/wohnungen-deutschland.json?sp=r&st=2023-02-02T13:19:48Z&se=2023-02-28T21:19:48Z&spr=https&sv=2021-06-08&sr=b&sig=rZRd03YH3AVwJPusaI10nAkAX5FOtSyKlPVd7T9Tqh0%3D")
	}
}
