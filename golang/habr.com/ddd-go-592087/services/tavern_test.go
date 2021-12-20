package services_test

import (
	"localmachine/habr-com/ddd-go-592087/aggregate"
	"localmachine/habr-com/ddd-go-592087/services"
	"testing"

	"github.com/google/uuid"
)

func TestTavern_Memory_Order(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := services.NewOrderService(
		services.WithMemoryCustomerRepository(),
		services.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := services.NewTavern(services.WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer("Will Smith")
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

	os, err := services.NewOrderService(
		services.WithMongoCustomerRepository("mongodb://localhost:27017"),
		services.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := services.NewTavern(services.WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer("Will Smith")
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
