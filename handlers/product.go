package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (h *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("In Product")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Something went wrong :(", http.StatusBadRequest)
	}
	fmt.Fprintf(rw, "Product, %s", d)
}
