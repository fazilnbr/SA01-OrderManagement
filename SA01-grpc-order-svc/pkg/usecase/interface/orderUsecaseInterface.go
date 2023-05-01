package interfaces

import (
	"context"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/domain"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/utils"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, order domain.RecOrder) (string, error)
	UpdateOrder(ctx context.Context, orderid, status string) (string, error)
	FetchOrder(ctx context.Context, userid int, filter domain.Filter, pagenation utils.Filter) ([]domain.RecOrder, utils.Metadata, error)
}
