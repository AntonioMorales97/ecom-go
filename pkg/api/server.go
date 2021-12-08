package api

import (
	db "github.com/AntonioMorales97/ecom-go/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("product_type_id", validProductTypeID)
	}

	router.POST("/order", server.createProductOrder)
	router.GET("/order/:id", server.getProductOrder)

	router.GET("/products", server.listProducts)

	router.POST("/product", server.createProduct)

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
