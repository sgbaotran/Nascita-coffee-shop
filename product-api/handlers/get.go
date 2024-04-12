package handlers

import (
	"net/http"

	"github.com/sgbaotran/Nascita-coffee-shop/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "GET: Something went wrong :(", http.StatusBadRequest)
	}
	p.l.Println("Handling adding products: ", lp)

}
