package interfaces

import (
	"context"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order domain.Order) (string, error)
	CreateItem(ctx context.Context, item domain.Item) (string, error)
	UpdateOrder(ctx context.Context, orderid, status string) (string, error)
	FetchOrder(ctx context.Context, userid int) ([]domain.Order, error)
	FetchItem(ctx context.Context, itemid string,filter domain.Filter) (domain.Item, error)
}
