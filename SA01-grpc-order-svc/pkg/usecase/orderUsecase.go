package usecase

import (
	"context"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/domain"
	repository "github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/repository/interface"
	interfaces "github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/usecase/interface"
)

type orderUseCase struct {
	orderRepo repository.OrderRepository
}

// CreateOrder implements interfaces.OrderUseCase
func (o *orderUseCase) CreateOrder(ctx context.Context, order domain.RecOrder) (int, error) {
	od:=domain.Order{}
	id, err := o.orderRepo.CreateOrder(ctx, od)
	return id, err
}

func NewOrderUseCase(repo repository.OrderRepository) interfaces.OrderUseCase {
	return &orderUseCase{
		orderRepo: repo,
	}
}
