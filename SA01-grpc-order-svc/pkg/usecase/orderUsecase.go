package usecase

import (
	repository "github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/repository/interface"
	interfaces "github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/usecase/interface"
)

type orderUseCase struct {
	orderRepo repository.OrderRepository
}

func NewOrderUseCase(repo repository.OrderRepository) interfaces.OrderUseCase {
	return &orderUseCase{
		orderRepo: repo,
	}
}
