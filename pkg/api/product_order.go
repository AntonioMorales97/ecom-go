package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/AntonioMorales97/ecom-go/db/sqlc"
	"github.com/AntonioMorales97/ecom-go/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateProductOrderParams{
		Quantity:  req.Quantity,
		ProductID: req.ProductID,
		Owner:     authPayload.Username,
	}

	productOrder, err := server.store.CreateProductOrderTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
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

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if productOrder.Owner != authPayload.Username {
		err := errors.New("product order doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, productOrder)
}

type listProductOrderRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listProductOrders(ctx *gin.Context) {
	var req listProductOrderRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListProductOrdersParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListProductOrders(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
