package data

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Product represents a product with basic information.
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
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
		SKU:         "ABC-ABC-ABC",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Smartphone",
		Description: "Feature-rich smartphone with a high-quality camera",
		Price:       599.99,
		SKU:         "ABC-ABC-ABC",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

func GetProducts() Products {
	return products
}

func GetProduct(id int) (*Product, error) {

	for _, product := range products {
		if id == product.ID {
			return product, nil
		}
	}
	return nil, fmt.Errorf("Product cannot be found")
}

type Products []*Product

func validateSKU(fl validator.FieldLevel) bool {

	reg := regexp.MustCompile("[A-Z]+-[A-Z]+-[A-Z]+")
	matches := reg.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}

	return true
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func AddProduct(p *Product) {
	p.ID = products[len(products)-1].ID + 1
	products = append(products, p)
}

func findProduct(id int) (*Product, int, error) {
	for ind, prod := range products {
		if prod.ID == id {
			return prod, ind, nil
		}
	}
	return nil, -1, fmt.Errorf("can not find such product")
}

func UpdateProduct(id int, p *Product) error {
	old_prod, ind, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = old_prod.ID
	products[ind] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")
