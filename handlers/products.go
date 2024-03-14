package handlers

import (
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

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "GET Something went wrong :(", http.StatusBadRequest)
	}
	p.l.Println("Handling adding products: ", lp)

}

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	var prod data.Product

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "POST Something went wrong :(", http.StatusBadRequest)
	}
	data.AddProduct(&prod)
	p.l.Println("Handling adding products: ", prod, err)

}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("In Product")
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)

	} else if r.Method == http.MethodPost {
		p.addProduct(rw, r)
	} else if r.Method == http.MethodPut {
		p.addProduct(rw, r)
	}
}
