package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sgbaotran/Nascita-coffee-shop/product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *Product) DeleteProduct(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "DELETE: Something went wrong (cannot find product) ("+err.Error()+") :(", http.StatusBadRequest)
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
