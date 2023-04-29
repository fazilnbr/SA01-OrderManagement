package order

import (
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	// auth := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	cart := r.Group("cart")

	// cart.Use(auth.AuthRequired)
	cart.POST("/order", svc.AddOrder)

	return svc
}

func (svc *ServiceClient) AddOrder(ctx *gin.Context) {

}
