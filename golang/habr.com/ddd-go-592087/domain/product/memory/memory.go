// Package memory is a in memory implementation of the ProductRepository interface.
package memory

import (
	"localmachine/habr-com/ddd-go-592087/domain/product"
	"sync"

	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

// GetAll returns all products as a slice
// Yes, it never returns an error, but
// A database implementation could return an error for instance
func (mpr *MemoryProductRepository) GetAll() ([]product.Product, error) {
	// Collect all Products from map
	var products []product.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

// GetByID searches for a product based on it's ID
func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (product.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}
	return product.Product{}, product.ErrProductNotFound
}

// Add will add a new product to the repository
func (mpr *MemoryProductRepository) Add(newprod product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()
	if _, ok := mpr.products[newprod.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}
	mpr.products[newprod.GetID()] = newprod
	return nil
}

// Update will change all values for a product based on it's ID
func (mpr *MemoryProductRepository) Update(updprod product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()
	_, err := mpr.GetByID(updprod.GetID())
	if err != nil {
		return err
	}
	mpr.products[updprod.GetID()] = updprod
	return nil
}

// Delete remove an product from the repository
func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()
	_, err := mpr.GetByID(id)
	if err != nil {
		return err
	}
	delete(mpr.products, id)
	return nil
}
