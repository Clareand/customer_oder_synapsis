package handler

import (
	"github.com/Clareand/customer_oder_synapsis/pkg/auth/usecase"

	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	usecase usecase.LoginUsecase
}

func NewHTTPHandler(usecase usecase.LoginUsecase) *HTTPHandler {
	return &HTTPHandler{usecase: usecase}
}

func (h *HTTPHandler) Mount(g *echo.Group) {
	g.POST("/auth/new-login", h.NewLoginUser)
	g.POST("/auth/logout", h.Logout)
	g.POST("/auth/refresh-token", h.RefreshToken)
}
