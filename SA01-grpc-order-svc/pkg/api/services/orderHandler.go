package services

import (
	"context"
	"net/http"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/domain"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/pb"
	usecase "github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/usecase/interface"
)

type OrderService struct {
	orderUseCase usecase.OrderUseCase
}

func (c *OrderService) GetCart(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	order := domain.RecOrder{
		ID:           req.UserId,
		Status:       req.Status,
		Total:        float64(req.Total),
		CurrencyUnit: req.CurrencyUnit,
	}

	id, err := c.orderUseCase.CreateOrder(ctx, order)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusOK,
		Id:     int64(id),
	}, nil

}

func NewOrderService(usecase usecase.OrderUseCase) *OrderService {
	return &OrderService{
		orderUseCase: usecase,
	}
}
