package routes

import (
	"fmt"
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
// @Param        status   query      string  false  "Status : "
// @Param        mintotal   query      string  false  "Min Total : "
// @Param        maxtolat   query      string  false  "Max Total : "
// @Param        sortby   query      string  false  "Sort By : "
// @Param        sortorder   query      string  false  "Sort Order : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /order [get]
func FetchOrder(ctx *gin.Context, c pb.OrderServiceClient) {

	id, _ := strconv.Atoi(ctx.Writer.Header().Get("userId"))

	status := ctx.Query("status")
	mintotal := ctx.Query("mintotal")
	mintotl, _ := strconv.Atoi(mintotal)
	maxtotal := ctx.Query("maxtolat")
	maxtotl, _ := strconv.Atoi(maxtotal)
	sortby := ctx.Query("sortby")
	sortorder := ctx.Query("sortorder")
	fmt.Println("here", sortorder)

	res, err := c.FetchOrder(ctx, &pb.FetchOrderRequest{
		UserId:    int64(id),
		Status:    status,
		MinTotal:  float32(mintotl),
		MaxTotal:  float32(maxtotl),
		SortBy:    sortby,
		SortOrder: sortorder,
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
