package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 10)

	assert.Nil(t, err)
	assert.NotNil(t, p)
}

func TestNewProductNameEmpty(t *testing.T) {
	p, err := NewProduct("", 10)

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestNewProductInvalidPrice(t *testing.T) {
	p, err := NewProduct("Product 1", -10)

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
}
