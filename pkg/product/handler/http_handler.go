package handler

import (
	"github.com/Clareand/customer_oder_synapsis/libs/models"
	"github.com/labstack/echo/v4"
)

func (h *HttpHandler) GetProductList(c echo.Context) error {
	category := c.QueryParam("category")
	data, err := h.usecase.GetProductList(category)
	if err != nil {
		return models.ToJSON(c).InternalServerError(err.Error())
	}
	return models.ToJSON(c).Ok(data, "Successfully")
}
