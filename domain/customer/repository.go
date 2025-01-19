package customer

import (
	"errors"

	"github.com/codeninepoint/ecommerce-platform/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add customer")
	ErrFailedToUpdateCustomer = errors.New("failed to update customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}