package handlers

import (
	"github.com/sgbaotran/Nascita-coffee-shop/product-api/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "PUT: Something went wrong (cannot convert ID) ("+err.Error()+") :(", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err != nil {
		http.Error(rw, "PUT: Something went wrong ("+err.Error()+") :(", http.StatusBadRequest)
		return
	}
	p.l.Println("Updated products: ", prod, err)

}
