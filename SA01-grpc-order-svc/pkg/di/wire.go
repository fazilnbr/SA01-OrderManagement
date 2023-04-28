//go:build wireinject
// +build wireinject

package di

import (
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/config"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/db"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) error {
	wire.Build(db.ConnectDatabase)

	return nil
}
