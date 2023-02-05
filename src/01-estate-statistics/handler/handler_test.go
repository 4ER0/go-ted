package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/4ER0/go-ted/estatestics/model"
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

func setup(fileName string, status int, error error) (*RealEstateInfoHandler, *MockRoundTripper) {
	var mr *MockRoundTripper
	c := &http.Client{}
	if fileName != "" {
		file, _ := os.ReadFile(fileName)
		mr = &MockRoundTripper{
			response: &http.Response{
				StatusCode: status,
				Body:       io.NopCloser(bytes.NewReader(file)),
			},
			error: error,
		}
		c = &http.Client{
			Transport: mr,
		}
	}
	l := logrus.New()
	l.SetLevel(logrus.DebugLevel)
	p := port.NewRealEstatePort(l, c)
	h := NewRealEstateInfoHandler(l, p)
	return h, mr
}

func doRequest(t *testing.T, hf func(w http.ResponseWriter, r *http.Request), method string, url string) *httptest.ResponseRecorder {

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hf)
	handler.ServeHTTP(rr, req)
	return rr
}

// Test default handler should return error
func TestDefaultHandler(t *testing.T) {
	h, _ := setup("", 0, nil)
	rr := doRequest(t, h.DefaultHandler, "GET", "/")

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
	if "Unknown URL\n" != (rr.Body.String()) {
		t.Errorf("handler returned unexpected body: got %q want %q",
			rr.Body.String(), "Unknown URL\n")
	}
}

// Test average handler should return correct average
func TestAverageHandler_ShouldReturnGoodResponse(t *testing.T) {
	response, _ := os.ReadFile("../data/average-response.json")
	h, mr := setup("../data/real-estate-mock-data.json", http.StatusOK, nil)
	rr := doRequest(t, h.GetAverageEmptyEstates, "GET", "/average?key=0001")

	var got model.AverageResponse
	_ = json.Unmarshal(rr.Body.Bytes(), &got)
	var expected model.AverageResponse
	_ = json.Unmarshal(response, &expected)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if expected != got {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, expected)
	}
	if mr.calledUrl != "https://manualmk.blob.core.windows.net/goted/wohnungen-deutschland.json?sp=r&st=2023-02-02T13:19:48Z&se=2023-02-28T21:19:48Z&spr=https&sv=2021-06-08&sr=b&sig=rZRd03YH3AVwJPusaI10nAkAX5FOtSyKlPVd7T9Tqh0%3D" {
		t.Errorf("handler called wrong url: got %q want %q",
			mr.calledUrl, "https://manualmk.blob.core.windows.net/goted/wohnungen-deutschland.json?sp=r&st=2023-02-02T13:19:48Z&se=2023-02-28T21:19:48Z&spr=https&sv=2021-06-08&sr=b&sig=rZRd03YH3AVwJPusaI10nAkAX5FOtSyKlPVd7T9Tqh0%3D")
	}
}

// Test average handler should handle connection error correctly
func TestAverageHandler_ShouldHandleConnectionError(t *testing.T) {
	h, mr := setup("unused", 0, fmt.Errorf("connection error"))
	rr := doRequest(t, h.GetAverageEmptyEstates, "GET", "/average?key=0001")

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
	if "Error while getting real estate info\n" != rr.Body.String() {
		t.Errorf("handler returned unexpected body: got %q want %q",
			rr.Body.String(), "Error while getting real estate info\n")
	}
	if mr.calledUrl != "https://manualmk.blob.core.windows.net/goted/wohnungen-deutschland.json?sp=r&st=2023-02-02T13:19:48Z&se=2023-02-28T21:19:48Z&spr=https&sv=2021-06-08&sr=b&sig=rZRd03YH3AVwJPusaI10nAkAX5FOtSyKlPVd7T9Tqh0%3D" {
		t.Errorf("handler called wrong url: got %q want %q",
			mr.calledUrl, "https://manualmk.blob.core.windows.net/goted/wohnungen-deutschland.json?sp=r&st=2023-02-02T13:19:48Z&se=2023-02-28T21:19:48Z&spr=https&sv=2021-06-08&sr=b&sig=rZRd03YH3AVwJPusaI10nAkAX5FOtSyKlPVd7T9Tqh0%3D")
	}
}
