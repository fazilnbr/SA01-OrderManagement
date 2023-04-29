package routes

import (
	"net/http"
	"strconv"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/domain"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/order/pb"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Create Order
// @ID createorder
// @Tags Cart
// @Produce json
// @Security BearerAuth
// @Param orderdetials body domain.Order{} true "Order Detials"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /order/items [post]
func CreateOrder(ctx *gin.Context, c pb.CartServiceClient) {
	body := domain.Order{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, _ := strconv.Atoi(ctx.Writer.Header().Get("userId"))
	res, err := c.CreateOrder(ctx, &pb.AddOrderRequest{
		UserId: int64(id),
	})

	if err != nil {
		responses := response.ErrorResponse("Failed to Add Product", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		response.ResponseJSON(*ctx, responses)
		return
	}

	responses := response.SuccessResponse(true, "SUCCESS", res)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusOK)
	response.ResponseJSON(*ctx, responses)
	return

}
