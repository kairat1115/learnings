package services_test

import (
	"localmachine/habr-com/ddd-go-592087/aggregate"
	"localmachine/habr-com/ddd-go-592087/services"
	"testing"

	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	beef, err := aggregate.NewProduct("Beef", "Healthy meat", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := aggregate.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	carrot, err := aggregate.NewProduct("Carrot", "Healthy vegetable", 0.30)
	if err != nil {
		t.Error(err)
	}
	products := []aggregate.Product{beef, peenuts, carrot}
	return products
}

func TestOrder_NewOrderService(t *testing.T) {
	// Create a few products to insert into in memory repo
	products := init_products(t)

	os, err := services.NewOrderService(
		services.WithMemoryCustomerRepository(),
		services.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	// Add Customer
	cust, err := aggregate.NewCustomer("Will Smith")
	if err != nil {
		t.Error(err)
	}

	err = os.AddCustomer(&cust)
	if err != nil {
		t.Error(err)
	}

	// Perform Order for one beef
	order := []uuid.UUID{
		products[0].GetID(),
	}
	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
