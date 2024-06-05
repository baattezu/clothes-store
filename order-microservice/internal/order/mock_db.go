package order

import (
	"context"
	"sync"
)

type MockDB struct {
	orders map[string]Order
	mu     sync.Mutex
}

func NewMockDB() *MockDB {
	return &MockDB{
		orders: make(map[string]Order),
	}
}

func (db *MockDB) CreateOrder(ctx context.Context, order Order) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.orders[order.OrderID] = order
	return nil
}

func (db *MockDB) GetOrder(ctx context.Context, orderID string) (Order, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.orders[orderID], nil
}
