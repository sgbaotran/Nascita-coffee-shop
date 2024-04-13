package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/sgbaotran/Nascita-coffee-shop/currency/data"
	protos "github.com/sgbaotran/Nascita-coffee-shop/currency/protos/currency"
)

type Currency struct {
	rates *data.ExchangeRate
	log   hclog.Logger
}

func NewServer(log hclog.Logger, rates *data.ExchangeRate) *Currency {
	return &Currency{log: log, rates: rates}
}

func (c *Currency) GetRate(ctx context.Context, req *protos.RateRequest) (*protos.RateResponse, error) {
	rate, err := c.rates.GetRates(req.Base.String(), req.Destination.String())
	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Rate: rate}, nil
}

func (c *Currency) mustEmbedUnimplementedCurrencyServer() {}
