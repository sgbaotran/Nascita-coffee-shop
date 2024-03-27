package testing

import (
	"github.com/sgbaotran/Nascita-coffee-shop/product-api/data"
	"testing"
)

func TestValidationFunctionality(t *testing.T) {
	p := data.Product{
		Name: "Bao Cao Su durex",
		Price: 22.0,
		SKU: "a-a-a",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}


