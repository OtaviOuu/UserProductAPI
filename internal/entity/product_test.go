package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1. Apenas a criação
func TestNewProduct(t *testing.T) {
	assert := assert.New(t)

	p, err := NewProduct("jorge", 1000)

	assert.Nil(err)
	assert.NotEmpty(p.ID)

	assert.NotNil(p.Name)
	assert.NotNil(p.Price)
}

func TestNewProduct_WhenNameIsRequired(t *testing.T) {
	assert := assert.New(t)

	p, err := NewProduct("", 1000)

	assert.Nil(p)
	assert.Equal(ErrNameIsRequerid, err)

}

func TestNewProduct_WhenPriceIsRequired(t *testing.T) {
	assert := assert.New(t)

	p, err := NewProduct("jorge", 0)

	assert.Nil(p)
	assert.Equal(ErrPriceIsRequerid, err)

}

func TestNewProduct_WhenPriceIsInvalid(t *testing.T) {
	assert := assert.New(t)

	p, err := NewProduct("jorge", 0)

	assert.Nil(p)
	assert.Equal(ErrPriceIsRequerid, err)

}
