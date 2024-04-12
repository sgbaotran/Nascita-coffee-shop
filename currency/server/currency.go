package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/sgbaotran/Nascita-coffee-shop/currency/protos/currency"
)

type Currency struct {
	log hclog.Logger
}

func NewServer(log hclog.Logger) *Currency {
	return &Currency{log: log}
}

func (c *Currency) GetRate(ctx context.Context, req *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("In get currency")
	return &protos.RateResponse{Rate: .5}, nil
}

func (c *Currency) mustEmbedUnimplementedCurrencyServer() {}
