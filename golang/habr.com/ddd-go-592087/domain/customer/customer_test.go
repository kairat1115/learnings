package customer_test

import (
	"localmachine/habr-com/ddd-go-592087/domain/customer"
	"testing"
)

func TestCustomer_NewCustomer(t *testing.T) {
	// Build our needed testcase data struct
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	// Create new test cases
	testCases := []testCase{
		{
			test:        "Empty Name validation",
			name:        "",
			expectedErr: customer.ErrInvalidPerson,
		}, {
			test:        "Valid Name",
			name:        "Will Smith",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		// Run tests
		t.Run(tc.test, func(t *testing.T) {
			// Create a new customer
			_, err := customer.NewCustomer(tc.name)
			// Check if the error matches the expected error
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
