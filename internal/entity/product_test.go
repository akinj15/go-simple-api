package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	// Create a new Product
	product, err := NewProduct("Notebook", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Notebook", product.Name)
	assert.NotEmpty(t, product.Name)
	assert.Equal(t, 10.0, product.Price)
	assert.NotEmpty(t, product.Price)
}
