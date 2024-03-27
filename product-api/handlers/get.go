package handlers

import (
	"microservice/data"
	"net/http"
)

func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "GET: Something went wrong :(", http.StatusBadRequest)
	}
	p.l.Println("Handling adding products: ", lp)

}
