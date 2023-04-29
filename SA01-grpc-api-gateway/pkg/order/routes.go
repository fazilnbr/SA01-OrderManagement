package order

import (
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/config"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/order/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	// auth := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	cart := r.Group("order")

	// cart.Use(auth.AuthRequired)
	cart.POST("/items", svc.CreateOrder)

	return svc
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
