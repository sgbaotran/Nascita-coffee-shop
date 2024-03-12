package data

import "time"

// Product represents a product with basic information.
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
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
