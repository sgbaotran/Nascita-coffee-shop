package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product represents a product with basic information.
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

var products = []*Product{
	&Product{
		ID:          1,
		Name:        "Laptop",
		Description: "Powerful laptop with high-performance specifications",
		Price:       999.99,
		SKU:         "ABC123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Smartphone",
		Description: "Feature-rich smartphone with a high-quality camera",
		Price:       599.99,
		SKU:         "DEF456",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

func GetProducts() Products {
	return products
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)

	return d.Decode(p)
}

func AddProduct(p *Product) {
	p.ID = products[len(products)-1].ID + 1
	products = append(products, p)
}
