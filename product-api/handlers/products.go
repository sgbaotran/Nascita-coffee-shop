package handlers

import (
	"fmt"
	"log"

	protos "github.com/sgbaotran/Nascita-coffee-shop/currency/protos/currency"
)

// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body int
}

// Products handler for getting and updating products
type Product struct {
	l  *log.Logger
	cc protos.CurrencyClient
}

// NewProducts returns a new products handler with the given logger
func NewProduct(l *log.Logger, cc protos.CurrencyClient) *Product {
	return &Product{l, cc}
}

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

var ErrProductNotFound = fmt.Errorf("Product not found")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}
