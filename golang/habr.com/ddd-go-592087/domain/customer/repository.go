// Package Customer holds all the domain logic for the customer domain.
package customer

import (
	"errors"
	"localmachine/habr-com/ddd-go-592087/aggregate"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	ErrUpdateCustomer      = errors.New("failed to update the customer in the repository")
)

// CustomerRepository is a interface that defines the rules around what a customer repository
type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
