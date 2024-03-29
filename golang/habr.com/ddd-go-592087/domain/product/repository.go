// Package product holds the repository and the implementations for a ProductRepository
package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound     = errors.New("the product was not found")
	ErrProductAlreadyExist = errors.New("the product already exists")
)

// ProductRepository is the repository interface to fulfill to use the product aggregate
type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(uuid.UUID) (Product, error)
	Add(Product) error
	Update(Product) error
	Delete(uuid.UUID) error
}
