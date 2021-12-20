package order_test

import (
	"localmachine/habr-com/ddd-go-592087/domain/customer"
	"localmachine/habr-com/ddd-go-592087/domain/product"
	"localmachine/habr-com/ddd-go-592087/services/order"
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

func TestOrder_NewOrderService(t *testing.T) {
	// Create a few products to insert into in memory repo
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	// Add Customer
	cust, err := customer.NewCustomer("Will Smith")
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
