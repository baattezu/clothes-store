package cart

import (
	"context"
	"sync"
)

type MockDB struct {
	cartItems map[string][]CartItem
	mu        sync.Mutex
}

func NewMockDB() *MockDB {
	return &MockDB{
		cartItems: make(map[string][]CartItem),
	}
}

func (db *MockDB) GetCartItems(ctx context.Context, userID string) ([]CartItem, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.cartItems[userID], nil
}

func (db *MockDB) AddToCart(ctx context.Context, userID string, item CartItem) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.cartItems[userID] = append(db.cartItems[userID], item)
	return nil
}
