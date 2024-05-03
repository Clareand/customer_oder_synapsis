package repository

import "github.com/Clareand/customer_oder_synapsis/pkg/product/model"

type ProductRepo interface {
	GetProduct(category string) ([]model.Products, error)
}
