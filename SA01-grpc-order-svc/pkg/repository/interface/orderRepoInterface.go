package interfaces

import (
	"context"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order domain.Order) (int, error)
}
