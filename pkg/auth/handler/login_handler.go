package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Clareand/customer_oder_synapsis/libs/models"
	"github.com/Clareand/customer_oder_synapsis/pkg/auth/model"

	"github.com/labstack/echo/v4"
)

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05.000"
)

func (h *HTTPHandler) NewLoginUser(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Reques 1")
	}
	t := model.ReqNewLogin{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request 2")
	}
	model := model.ReqNewLogin{
		Username:   t.Username,
		Password:   t.Password,
		RememberMe: t.RememberMe,
	}

	ipnumber := c.RealIP()
	result := <-h.usecase.WithContext(c).NewLoginUser(model, ipnumber)

	if result.Error != nil {
		resp := &models.Response{Code: 400, MessageCode: 13, Message: result.Error.Error()}
		return c.JSON(http.StatusBadRequest, resp)

	}
	return c.JSON(http.StatusOK, result.Data)
}

func (h *HTTPHandler) RefreshToken(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}
	t := model.AccessToken{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}
	model := model.AccessToken{
		RefreshToken: t.RefreshToken,
	}
	result := <-h.usecase.WithContext(c).RefreshToken(model)
	if result.Error != nil {
		resp := &models.Response{Code: 400, MessageCode: 13, Message: result.Error.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}
	return c.JSON(http.StatusOK, result.Data)
}

func (h *HTTPHandler) Logout(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}
	t := model.AccessToken{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}
	model := model.AccessToken{
		RefreshToken: t.RefreshToken,
	}
	result := <-h.usecase.WithContext(c).Logout(model)
	if result.Error != nil {
		resp := &models.Response{Code: 400, MessageCode: 13, Message: result.Error.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}
	return c.JSON(http.StatusOK, result.Data)
}
