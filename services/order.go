package services

import (
	"context"
	"log"

	"github.com/codeninepoint/ecommerce-platform/aggregate"
	"github.com/codeninepoint/ecommerce-platform/domain/customer"
	"github.com/codeninepoint/ecommerce-platform/domain/customer/memory"
	"github.com/codeninepoint/ecommerce-platform/domain/customer/mongo"
	"github.com/codeninepoint/ecommerce-platform/domain/product"
	prodmem "github.com/codeninepoint/ecommerce-platform/domain/product/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error){
	os := &OrderService{}

	//Loop throuhg all the Cfgs and apply them
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}
	return os, nil

}

// WithCutomerRepository applies customer repository to OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// Returns the function that matches orderconfiguration alias

	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)

}

func WithMongoCustomerRepository(ctx context.Context, connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
	
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}

		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(CustomerId uuid.UUID, ProductId []uuid.UUID ) (float64, error) {
	//Fetch the customer
	c, err := o.customers.Get(CustomerId)
	if err != nil {
		return 0, err
	}

	//Get each product from product repository
	var products []aggregate.Product
	var total float64

	for _, id := range ProductId {
		p, err := o.products.GetById(id)

		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))
	return total, nil
}