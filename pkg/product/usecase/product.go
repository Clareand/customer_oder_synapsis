package usecase

import "github.com/Clareand/customer_oder_synapsis/pkg/product/model"

type ProductUsecase interface {
	GetProductList(category string) ([]model.Products, error)
}
