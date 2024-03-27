// Package classification of Product API
// Documentation for product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//	Consumes:
//	- application/json
//	Produces:
//	- application/json
//
// swagger: meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/sgbaotran/Nascita-coffee-shop/product-api/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

type KeyProduct struct{}

func ValidateProductMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var prod data.Product

		err := prod.FromJSON(r.Body)

		if err != nil {
			http.Error(rw, "POST: Something went wrong (failed to serialize json) ("+(err.Error())+" :(", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			http.Error(rw, "POST: Something went wrong (validation fail) ("+(err.Error())+" :(", http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)

		r = r.WithContext(ctx)

		fmt.Println("In Middleware: ")

		next.ServeHTTP(rw, r)
	})

}
