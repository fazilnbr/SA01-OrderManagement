package order

import (
	"fmt"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/config"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CartServiceClient
}

func InitServiceClient(c *config.Config) pb.CartServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCartServiceClient(cc)
}
