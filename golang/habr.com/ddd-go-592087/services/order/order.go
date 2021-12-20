package order

import (
	"context"
	"localmachine/habr-com/ddd-go-592087/domain/customer"
	customer_memory "localmachine/habr-com/ddd-go-592087/domain/customer/memory"
	"localmachine/habr-com/ddd-go-592087/domain/customer/mongo"
	"localmachine/habr-com/ddd-go-592087/domain/product"
	product_memory "localmachine/habr-com/ddd-go-592087/domain/product/memory"
	"log"

	"github.com/google/uuid"
)

// OrderService is a implementation of the OrderService
type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

// OrderConfiguration is an alias for a function that will take in a pointer to an OrderService and modify it
type OrderConfiguration func(os *OrderService) error

// NewOrderService takes a variable amount of OrderConfiguration functions and returns a new OrderService
// Each OrderConfiguration will be called in the order they are passed in
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository applies a given customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

// WithMemoryCustomerRepository applies a memory customer repository to the OrderService
func WithMemoryCustomerRepository() OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	cr := customer_memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the mongo repo, if we needed parameters, such as connection strings they could be inputted here
		cr, err := mongo.New(context.Background(), connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

// WithMemoryProductRepository adds a in memory product repo and adds all input products
func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		pr := product_memory.New()

		// Add Items to repo
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func (o *OrderService) GetCustomers() customer.CustomerRepository {
	return o.customers
}

func (o *OrderService) AddCustomer(customer *customer.Customer) error {
	err := o.customers.Add(*customer)
	return err
}

func (o *OrderService) GetProducts() product.ProductRepository {
	return o.products
}

// func (o *OrderService) SetProducts()

// CreateOrder will chain together all repositories to create a order for a customer
// will return the collected price of all Products
func (o *OrderService) CreateOrder(cutomerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	// Get the customer
	c, err := o.customers.Get(cutomerID)
	if err != nil {
		return 0, err
	}

	// Get each Product
	var products []product.Product
	var price float64
	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, nil
		}
		products = append(products, p)
		price += p.GetPrice()
	}

	// All Products exists in store, now we can create the order
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))
	return price, nil
}
