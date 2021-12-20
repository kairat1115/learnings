// Package memory is a in-memory implementation of the customer repository
package memory

import (
	"fmt"
	"localmachine/habr-com/ddd-go-592087/domain/customer"
	"sync"

	"github.com/google/uuid"
)

// MemoryRepository fulfills the CustomerRepository interface
type MemoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

// Get finds a customer by ID
func (mr *MemoryRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return customer.Customer{}, customer.ErrCustomerNotFound
}

// Add will add a new customer to the repository
func (mr *MemoryRepository) Add(c customer.Customer) error {
	if mr.customers == nil {
		// Safety check if customers is not create, shouldn't happen if using the Factory, but you never know
		func() {
			mr.Lock()
			defer mr.Unlock()
			mr.customers = make(map[uuid.UUID]customer.Customer)
		}()
	}
	// Make sure Customer isn't already in the repository
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}
	func() {
		mr.Lock()
		defer mr.Unlock()
		mr.customers[c.GetID()] = c
	}()
	return nil
}

// Update will replace an existing customer information with the new customer information
func (mr *MemoryRepository) Update(c customer.Customer) error {
	// Make sure Customer is in the repository
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}
	func() {
		mr.Lock()
		defer mr.Unlock()
		mr.customers[c.GetID()] = c
	}()
	return nil
}
