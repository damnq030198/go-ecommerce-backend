package ticket

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anonystick/go-ecommerce-backend-api/internal/service"
	"github.com/anonystick/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

// manager controller Ticket Item
var TicketItem = new(cTicketItem)

type cTicketItem struct{}

// NewTicketItem creates a new

func (p *cTicketItem) GetTicketItemById(ctx *gin.Context) {
	// get the ticket item
	ticket_item := ctx.Param("id")
	// Convert the string parameter to an integer.
	idInt, err := strconv.Atoi(ticket_item)
	if err != nil {
		// Handle the conversion error.  This is crucial!
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket item ID format"})
		return
	}
	// call implementation
	ticketItem, err := service.TicketItem().GetTicketItemById(ctx, idInt)
	if err != nil {
		if errors.Is(err, response.CouldNotGetTicketErr) {
			fmt.Println("4004???")
		}

		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())

	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, ticketItem)
}
