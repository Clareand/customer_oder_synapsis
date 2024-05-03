package repository

import (
	"context"

	"github.com/Clareand/customer_oder_synapsis/config/postgresql"
	"github.com/Clareand/customer_oder_synapsis/pkg/cart/model"
	"github.com/labstack/echo/v4"
)

type cartRepo struct {
	dbConn  *postgresql.DbConnection
	ctx     context.Context
	echoCtx echo.Context
}

func NewCartRepo(dbConn *postgresql.DbConnection) CartRepo {
	return &cartRepo{dbConn: dbConn}
}

func (r *cartRepo) GetCart(customer_id string) ([]model.Carts, error) {
	var carts []model.Carts

	params := make([]interface{}, 0)

	params = append(params, customer_id)

	sql := `select * from public.f_get_cart(?)`
	err := r.dbConn.Db.Raw(sql, params...).Scan(&carts).Error

	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *cartRepo) AddToCart(param model.AddToCart) error {
	db := r.dbConn.Db
	paramData := make([]interface{}, 0)
	paramData = append(paramData, param.CustomerID, param.ProductID, param.Quantity)

	sql := `CALL public.p_add_cart(?,?,?)`

	return db.Exec(sql, paramData...).Error
}
