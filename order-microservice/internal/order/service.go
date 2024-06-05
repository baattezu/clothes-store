package order

import (
	"context"
	"github.com/google/uuid"
)

type Service struct {
	repo *MockDB
}

func NewService(repo *MockDB) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateOrder(ctx context.Context, userID string, dto CreateOrderDTO) (string, error) {
	orderID := uuid.New().String()
	order := Order{
		OrderID: orderID,
		UserID:  userID,
		Items:   dto.Items,
	}
	err := s.repo.CreateOrder(ctx, order)
	if err != nil {
		return "", err
	}
	return orderID, nil
}

func (s *Service) GetOrder(ctx context.Context, orderID string) (Order, error) {
	return s.repo.GetOrder(ctx, orderID)
}
