package routes

import (
	"context"
	"net/http"

	"github.com/fazilnbr/SA01-OrderManagement/pb"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/domain"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @title Go + Gin Seller App Order Management Sytem
// @version 1.0
// @description This is a sample Order Management Sytem. You can visit the GitHub repository at https://github.com/fazilnbr/SA01-OrderManagement

// @contact.name API Support
// @contact.url https://fazilnbr.github.io/mypeosolal.web.portfolio/
// @contact.email fazilkp2000@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /
// @query.collection.format multi

// @Summary SignUp for users
// @ID SignUp authentication
// @Tags Authentication
// @Produce json
// @Param WorkerLogin body domain.User{username=string,password=string} true "Worker Login"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /auth/register [post]
func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := domain.User{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		responses := response.ErrorResponse("Failed to create user", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		response.ResponseJSON(*ctx, responses)
		return
	}

	responses := response.SuccessResponse(true, "SUCCESS", res)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusBadRequest)
	response.ResponseJSON(*ctx, responses)
	return

}
