package api

import (
	"database/sql"
	"net/http"

	db "github.com/AntonioMorales97/ecom-go/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createProductOrderRequest struct {
	Quantity  int32 `json:"quantity" binding:"required,min=1"`
	ProductID int64 `json:"product_id" binding:"required,min=0"`
}

func (server *Server) createProductOrder(ctx *gin.Context) {
	var req createProductOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProductOrderParams{
		Quantity:  req.Quantity,
		ProductID: req.ProductID,
	}

	productOrder, err := server.store.CreateProductOrder(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, productOrder)
}

type getProductOrderRequest struct {
	ID int64 `uri:"id" binding:"required,min=0"`
}

func (server *Server) getProductOrder(ctx *gin.Context) {
	var req getProductOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	productOrder, err := server.store.GetProductOrder(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, productOrder)
}
