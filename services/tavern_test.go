package services

import (
	"context"
	"testing"

	"github.com/codeninepoint/ecommerce-platform/aggregate"
	"github.com/google/uuid"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		// use WithMemoryCustomerRepository to use only memory
		// WithMemoryCustomerRepository(),

		WithMongoCustomerRepository(context.Background(), "mongodb+srv://rupeshthakur:rNK3FFaEPiPy4BSo@cluster0-dev.x9kv6pm.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0-Dev"),

		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("neha")
	if err != nil {
		t.Fatal(err)
	}

	if err = os.customers.Add(cust); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID {
		products[0].GetID(),
	}

	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Fatal(err)
	}
}