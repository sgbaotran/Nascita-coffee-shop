package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/go-hclog"
)

type ExchangeRate struct {
	l    hclog.Logger
	rate map[string]float64
}

func NewExchangeRate(l hclog.Logger) (*ExchangeRate, error) {
	return &ExchangeRate{l, map[string]float64{}}, nil
}

func (er *ExchangeRate) GetRates(base, destination string) (float64, error) {
	er.getRates()
	fmt.Println(er.rate)
	fmt.Println(base, destination)
	br, ok := er.rate[base]
	fmt.Println(ok)
	if !ok {
		return 0, fmt.Errorf("Rate not found for the currency ", br)
	}

	dr, ok := er.rate[destination]
	fmt.Println(ok)
	if !ok {
		return 0, fmt.Errorf("Rate not found for the currency ", dr)
	}

	return dr / br, nil
}

func (er *ExchangeRate) getRates() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")

	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected error code 200 got %d", resp.StatusCode)
	}
	md := &Cubes{}
	xml.NewDecoder(resp.Body).Decode(&md)

	for _, c := range md.CubeData {
		r, err := strconv.ParseFloat(c.Rate, 64)

		if err != nil {
			return nil
		}
		er.rate[c.Currency] = r
	}
	return nil
}

type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

type Cube struct {
	Rate     string `xml:"rate,attr"`
	Currency string `xml:"currency,attr"`
}
