package repository

import (
	"context"

	"github.com/Clareand/customer_oder_synapsis/config/postgresql"
	"github.com/Clareand/customer_oder_synapsis/pkg/product/model"
	"github.com/labstack/echo/v4"
)

type productRepo struct {
	dbConn  *postgresql.DbConnection
	ctx     context.Context
	echoCtx echo.Context
}

func NewProductRepo(dbConn *postgresql.DbConnection) ProductRepo {
	return &productRepo{dbConn: dbConn}
}

func (r *productRepo) GetProduct(category string) ([]model.Products, error) {
	var products []model.Products

	params := make([]interface{}, 0)
	if category != "" {
		params = append(params, category)
	} else {
		params = append(params, nil)
	}

	sql := `select * from public.f_get_all_product(?)`
	err := r.dbConn.Db.Raw(sql, params...).Scan(&products).Error

	if err != nil {
		return nil, err
	}
	return products, nil
}
