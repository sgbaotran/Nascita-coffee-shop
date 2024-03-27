package handlers

import (
	"microservice/data"
	"net/http"
)

func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)

	p.l.Println("Added products: ", prod)

}
