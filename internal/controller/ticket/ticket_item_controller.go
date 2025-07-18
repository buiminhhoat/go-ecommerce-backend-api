package ticket

import (
	"fmt"

	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/service"
	"github.com/buiminhhoat/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

// manager controller Ticket Item
var TicketItem = new(cTicketItem)

type cTicketItem struct{}

// NewTicketItem creates a new

func (p *cTicketItem) GetTicketItemById(ctx *gin.Context) {
	fmt.Println("Controller: GetTicketItemById")
	// call implementation
	ticketItem, err := service.TicketItem().GetTicketItemById(ctx, 1)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, ticketItem)
}
