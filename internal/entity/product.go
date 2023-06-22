package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrIdIsRequired    = errors.New("id is required")
	ErrInvalidID       = errors.New("invalid id")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at  "`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        uuid.New(),
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

	//add demais validations

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price <= 0 {
		return ErrInvalidPrice
	}

	return nil
}
