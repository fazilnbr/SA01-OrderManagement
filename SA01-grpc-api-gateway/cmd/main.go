package main

import (
	"log"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/config"
	_"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/cmd/api/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(cfg.Port)

}
