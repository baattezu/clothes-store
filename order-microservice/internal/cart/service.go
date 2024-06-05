package cart

import (
	"context"
)

type Service struct {
	repo *MockDB
}

func NewService(repo *MockDB) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetCartItems(ctx context.Context, userID string) ([]CartItem, error) {
	return s.repo.GetCartItems(ctx, userID)
}

func (s *Service) AddToCart(ctx context.Context, userID string, dto AddToCartDTO) error {
	item := CartItem{
		ProductID: dto.ProductID,
		Quantity:  dto.Quantity,
	}
	return s.repo.AddToCart(ctx, userID, item)
}
