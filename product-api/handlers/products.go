// Package classification of Product API
//
// Documentation for product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"microservice/data"
	"net/http"
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

/*
func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	// if r.Method == http.MethodGet {
	// 	p.GetProducts(rw, r)

	// } else if r.Method == http.MethodPost {
	// 	p.addProduct(rw, r)

	// } else if r.Method == http.MethodPut {
	// 	reg := regexp.MustCompile("/([0-9]+)")

	// 	g := reg.FindAllStringSubmatch(r.URL.Path, -1)

	// 	if len(g) != 1 {
	// 		http.Error(rw, "PUT: Invalid ID (having more than 1) :(", http.StatusBadRequest)
	// 		return
	// 	}

	// 	if len(g[0]) != 2 {
	// 		http.Error(rw, "PUT: Invalid ID (having more than 1) :(", http.StatusBadRequest)
	// 		return
	// 	}

	// 	id, err := strconv.Atoi(g[0][1])

	// 	if err != nil {
	// 		http.Error(rw, "PUT: Cannot parse your id :(", http.StatusBadRequest)
	// 		return
	// 	}
	// 	p.updateProduct(id, rw, r)
	// }
}
*/
