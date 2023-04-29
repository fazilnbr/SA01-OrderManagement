package http

import (
	"fmt"
	"log"
	"net"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/api/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewGRPCServer(orderService *services.OrderService, grpcPort string) *grpc.Server {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	fmt.Println("grpcPort/////", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// pb.RegisterCartServiceServer(grpcServer, cartService)

	if err := grpcServer.Serve(lis); err != nil {
		// log.Fatalf("failed to serve: %v", err).
	}
	return grpcServer
}
func NewServerHTTP(orderService *services.OrderService) *ServerHTTP {
	engine := gin.New()
	go NewGRPCServer(orderService, "50083")
	// Use logger from Gin
	engine.Use(gin.Logger())

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":9090")
}
