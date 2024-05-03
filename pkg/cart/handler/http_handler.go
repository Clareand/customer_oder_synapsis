package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Clareand/customer_oder_synapsis/libs/models"
	"github.com/Clareand/customer_oder_synapsis/pkg/cart/model"
	"github.com/labstack/echo/v4"
)

func (h *HttpHandler) GetCart(c echo.Context) error {
	customer_id := c.QueryParam("customer_id")
	fmt.Println("Hello", customer_id)
	data, err := h.usecase.GetCart(customer_id)
	if err != nil {
		return models.ToJSON(c).InternalServerError(err.Error())
	}

	return models.ToJSON(c).Ok(data, "Successfully")
}

func (h *HttpHandler) AddToCart(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}

	params := model.AddToCart{}
	err = json.Unmarshal(body, &params)

	err = h.usecase.AddToCart(params)
	if err != nil {
		return models.ToJSON(c).InternalServerError(err.Error())
	}
	return models.ToJSON(c).Ok(err, "Successfully")

}
