package tavern_test

import (
	"localmachine/habr-com/ddd-go-592087/domain/customer"
	"localmachine/habr-com/ddd-go-592087/domain/product"
	"localmachine/habr-com/ddd-go-592087/services/order"
	"localmachine/habr-com/ddd-go-592087/services/tavern"
	"testing"

	"github.com/google/uuid"
)

func init_products(t *testing.T) []product.Product {
	beef, err := product.NewProduct("Beef", "Healthy meat", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	carrot, err := product.NewProduct("Carrot", "Healthy vegetable", 0.30)
	if err != nil {
		t.Error(err)
	}
	products := []product.Product{beef, peenuts, carrot}
	return products
}

func TestTavern_Memory_Order(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := tavern.NewTavern(tavern.WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := customer.NewCustomer("Will Smith")
	if err != nil {
		t.Error(err)
	}

	err = os.AddCustomer(&cust)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}

func TestTavern_Mongo_Order(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository("mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := tavern.NewTavern(tavern.WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := customer.NewCustomer("Will Smith")
	if err != nil {
		t.Error(err)
	}

	err = os.AddCustomer(&cust)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
