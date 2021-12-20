package memory_test

import (
	"localmachine/habr-com/ddd-go-592087/aggregate"
	"localmachine/habr-com/ddd-go-592087/domain/product"
	"localmachine/habr-com/ddd-go-592087/domain/product/memory"
	"testing"

	"github.com/google/uuid"
)

func TestMemory_ProductRepositoryAdd(t *testing.T) {
	repo := memory.New()
	product, err := aggregate.NewProduct("Beef", "Good for your health", 1.99)
	if err != nil {
		t.Error(err)
	}
	repo.Add(product)
	products, err := repo.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(products))
	}
}

func TestMemory_ProductRepositoryGet(t *testing.T) {
	repo := memory.New()
	existingProd, err := aggregate.NewProduct("Beef", "Good for your health", 1.99)
	if err != nil {
		t.Error(err)
	}
	repo.Add(existingProd)
	products, err := repo.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(products))
	}

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get product by id",
			id:          existingProd.GetID(),
			expectedErr: nil,
		}, {
			name:        "Get non-existing product by id",
			id:          uuid.New(),
			expectedErr: product.ErrProductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetByID(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemory_ProductRepositoryDelete(t *testing.T) {
	repo := memory.New()
	existingProd, err := aggregate.NewProduct("Beef", "Good for your health", 1.99)
	if err != nil {
		t.Error(err)
	}
	repo.Add(existingProd)
	products, err := repo.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(products))
	}

	err = repo.Delete(existingProd.GetID())
	if err != nil {
		t.Error(err)
	}

	products, err = repo.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(products) != 0 {
		t.Errorf("Expected 0 products, got %d", len(products))
	}
}
