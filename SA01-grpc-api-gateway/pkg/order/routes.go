package order

import (
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/auth"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/config"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/order/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	auth := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	order := r.Group("order")

	order.Use(auth.AuthRequired)
	order.POST("/", svc.CreateOrder)
	order.PUT("/", svc.UpdateOrder)
	order.GET("/", svc.FetchOrder)

	return svc
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
func (svc *ServiceClient) UpdateOrder(ctx *gin.Context) {
	routes.UpdateOrder(ctx, svc.Client)
}
func (svc *ServiceClient) FetchOrder(ctx *gin.Context) {
	routes.FetchOrder(ctx, svc.Client)
}
