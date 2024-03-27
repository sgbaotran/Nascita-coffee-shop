package handlers

import (
	"github.com/sgbaotran/Nascita-coffee-shop/product-api/data"
	"net/http"
)

func (p *Product) DeleteProduct(rw http.ResponseWriter, r *http.Request) {

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)

	p.l.Println("Added products: ", prod)

}
