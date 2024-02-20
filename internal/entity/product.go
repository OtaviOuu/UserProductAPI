package entity

import (
	"errors"
	"time"

	"github.com/crudGolangAPI/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrInvalidId       = errors.New("invalid id")
	ErrNameIsRequerid  = errors.New("name is required")
	ErrPriceIsRequerid = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if p.Name == "" {
		return ErrNameIsRequerid
	}

	if p.Price == 0 {
		return ErrPriceIsRequerid
	}

	// Validação separada
	if p.Price <= 0 {
		return ErrInvalidPrice
	}

	return nil

}
