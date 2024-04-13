package handlers

import (
	"context"
	"net/http"

	protos "github.com/sgbaotran/Nascita-coffee-shop/currency/protos/currency"

	"github.com/sgbaotran/Nascita-coffee-shop/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_CAD),
		Destination: protos.Currencies(protos.Currencies_USD),
	}
	resp, err := p.cc.GetRate(context.Background(), rr)
	p.l.Println("Handling adding products: ", lp, resp)

	if err != nil {
		http.Error(rw, "GET: Something went wrong :(", http.StatusBadRequest)
	}

	err = data.ToJSON(lp, rw)

	if err != nil {
		http.Error(rw, "GET: Something went wrong :(", http.StatusBadRequest)
	}

}

// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Product) GetProduct(rw http.ResponseWriter, r *http.Request) {

	id := getProductID(r)

	prod, err := data.GetProduct(id)

	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_USD),
		Destination: protos.Currencies(protos.Currencies_CAD),
	}

	resp, err := p.cc.GetRate(context.Background(), rr)
	p.l.Println("Handling adding products: ", prod, resp)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("Unable to fetch product", "error", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		// p.l.Error("Unable to fetching product", "error", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("Unable to serializing product", err)
	}
}
