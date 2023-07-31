package repository

import "orders/entity"

// Persists Order to db, filesystem, etc.
type OrderRepository interface {
	Save(order *entity.Order) error
	FindById(orderId string) *entity.Order
}

// InMemoryOrderRepo.
type InMemoryOrderRepo struct {
	orders map[string]*entity.Order
}

// Constructor.
func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		orders: make(map[string]*entity.Order),
	}
}

func (repo *InMemoryOrderRepo) Save(order *entity.Order) error {
	repo.orders[order.Uuid] = order
	return nil
}

func (repo *InMemoryOrderRepo) FindById(orderId string) *entity.Order {
	for _, order := range repo.orders {
		if order.Uuid == orderId {
			return order
		}
	}
	return nil
}
