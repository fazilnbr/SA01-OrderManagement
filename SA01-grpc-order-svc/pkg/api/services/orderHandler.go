package services

import usecase "github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/usecase/interface"

type OrderService struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderService(usecase usecase.OrderUseCase) *OrderService {
	return &OrderService{
		orderUseCase: usecase,
	}
}
