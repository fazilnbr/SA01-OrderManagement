package routes

import (
	"net/http"
	"strconv"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-api-gateway/pkg/utils/response"
	"github.com/fazilnbr/SA01-OrderManagement/pb"
	"github.com/gin-gonic/gin"
)

// @Summary Fetch Order
// @ID Fetchorder
// @Tags Order
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /order [get]
func FetchOrder(ctx *gin.Context, c pb.OrderServiceClient) {

	id,_ := strconv.Atoi(ctx.Writer.Header().Get("userId"))

	res, err := c.FetchOrder(ctx, &pb.FetchOrderRequest{
		UserId: int64(id),
	})

	if err != nil {
		responses := response.ErrorResponse("Failed to Fetch Order", err.Error(), nil)
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
