package handler

import (
	"github.com/Clareand/customer_oder_synapsis/config/postgresql"
	"github.com/Clareand/customer_oder_synapsis/pkg/cart/usecase"
	"github.com/labstack/echo/v4"
)

type HttpHandler struct {
	usecase usecase.CartUsecase
}

func NewHTTPHandler(usecase usecase.CartUsecase) *HttpHandler {
	return &HttpHandler{usecase: usecase}
}

func (h *HttpHandler) Mount(g *echo.Group, auth echo.MiddlewareFunc, dbConn *postgresql.DbConnection) {
	g.GET("/cart", h.GetCart, auth)
	g.POST("/cart/add", h.AddToCart, auth)
}
