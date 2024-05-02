package models

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type httpContext struct {
	c echo.Context
}

type Response struct {
	Code        int         `json:"code"`
	MessageCode int         `json:"message_code"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

type Result struct {
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}

type CountData struct {
	Count int `json:"count" gorm:"column:count"`
}

func ToJSON(c echo.Context) *httpContext {
	httpContext := &httpContext{c}
	return httpContext
}

func (httpContext httpContext) Ok(data interface{}, msg string) error {
	return httpContext.c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
	})
}

func (httpContext httpContext) InternalServerError(msg string) error {
	return httpContext.c.JSON(http.StatusInternalServerError, &Response{
		Code:    http.StatusBadRequest,
		Message: msg,
		Data:    nil,
	})
}

func (httpContext httpContext) BadRequest(msg string) error {
	return httpContext.c.JSON(http.StatusBadRequest, &Response{
		Code:    http.StatusBadRequest,
		Message: msg,
		Data:    nil,
	})
}

func (httpContext httpContext) StatusNotFound(msg string) error {
	return httpContext.c.JSON(http.StatusNotFound, &Response{
		Code:    http.StatusNotFound,
		Message: msg,
		Data:    nil,
	})
}
