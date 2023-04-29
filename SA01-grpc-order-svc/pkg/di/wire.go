//go:build wireinject
// +build wireinject

package di

import (
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/api/services"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/config"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/db"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/repository"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/usecase"

	http "github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/api"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewOrderRepository,
		usecase.NewOrderUseCase,
		services.NewOrderService,
		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
