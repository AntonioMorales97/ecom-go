package api

import (
	"net/http"

	db "github.com/AntonioMorales97/ecom-go/db/sqlc"
	"github.com/gin-gonic/gin"
)

type listProductsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listProducts(ctx *gin.Context) {
	var req listProductsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListProductsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	products, err := server.store.ListProducts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, products)
}

type createProductRequest struct {
	Name              string `json:"name" binding:"required"`
	DescriptionLong   string `json:"description_long" `
	DescriptionShort  string `json:"description_short"`
	Price             int32  `json:"price" binding:"required"`
	ProductTypeID     int64  `json:"product_type_id" binding:"required,product_type_id"`
	ProductCategoryID *int64 `json:"product_category_id"`
	Quantity          int32  `json:"quantity" binding:"required,gte=0"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProductTxParams{
		Name:              req.Name,
		DescriptionLong:   req.DescriptionLong,
		DescriptionShort:  req.DescriptionShort,
		Price:             req.Price,
		ProductTypeID:     req.ProductTypeID,
		ProductCategoryID: req.ProductCategoryID,
		Quantity:          req.Quantity,
	}

	result, err := server.store.CreateProductTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
