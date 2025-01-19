package memory

import (
	"errors"
	"testing"

	"github.com/codeninepoint/ecommerce-platform/aggregate"
	"github.com/codeninepoint/ecommerce-platform/domain/customer"
	"github.com/google/uuid"
)

func TestMemory_GetCustom(t *testing.T){
	type testCase struct{
		name string
		id	uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("rupesh")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository {
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name: "no customer by id",
			id: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		}, {
			name: "customer by id",
			id: id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases{
		t.Run(tc.name, func (t *testing.T)  {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
	
}