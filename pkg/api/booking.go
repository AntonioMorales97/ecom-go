package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createBookingRequest struct {
	Name string `json:"name" binding:"required"`
}

type Booking struct {
	Name string `json:"name"`
}

func (server *Server) createBooking(ctx *gin.Context) {
	var req createBookingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	booking, err := book(req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	ctx.JSON(http.StatusOK, booking)
}

func book(name string) (Booking, error) {

	return Booking{Name: name}, nil
}
