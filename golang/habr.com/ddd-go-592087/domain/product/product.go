// Product is an aggregate that represents a product.
package product

import (
	"errors"
	"localmachine/habr-com/ddd-go-592087/models"

	"github.com/google/uuid"
)

var ErrMissingValues = errors.New("missing values")

// Product is a aggregate that combines item with a price and quantity
type Product struct {
	// item is the root entity which is an item
	item  *models.Item
	price float64
	// Quantity is the number of products in stock
	quantity int
}

func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}
	return Product{
		item: &models.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *models.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
