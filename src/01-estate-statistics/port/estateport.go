package port

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/4ER0/go-ted/estatestics/model"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

const uRL = "https://manualmk.blob.core.windows.net/goted/wohnungen-deutschland.json?sp=r&st=2023-02-02T13:19:48Z&se=2023-02-28T21:19:48Z&spr=https&sv=2021-06-08&sr=b&sig=rZRd03YH3AVwJPusaI10nAkAX5FOtSyKlPVd7T9Tqh0%3D"

type RealEstatePort struct {
	l *logrus.Logger
	c *http.Client
}

func NewRealEstatePort(l *logrus.Logger, c *http.Client) *RealEstatePort {
	return &RealEstatePort{l, c}
}

func (p *RealEstatePort) GetRealEstateInfo() ([]model.RealEstate, error) {
	res, err := p.c.Get(uRL)
	if err != nil {
		return nil, fmt.Errorf("could not get response: %w", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get real estate info: %s", res.Status)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	var re []model.RealEstate
	err = json.Unmarshal(buf.Bytes(), &re)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %w", err)
	}

	return re, nil
}
