// Package product holds the repository and the implementations for a ProductRepository
package product

import (
	"errors"
	"localmachine/habr-com/ddd-go-592087/aggregate"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound     = errors.New("the product was not found")
	ErrProductAlreadyExist = errors.New("the product already exists")
)

// ProductRepository is the repository interface to fulfill to use the product aggregate
type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(uuid.UUID) (aggregate.Product, error)
	Add(aggregate.Product) error
	Update(aggregate.Product) error
	Delete(uuid.UUID) error
}
