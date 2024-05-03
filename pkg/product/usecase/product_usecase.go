package usecase

import (
	"github.com/Clareand/customer_oder_synapsis/pkg/product/model"
	"github.com/Clareand/customer_oder_synapsis/pkg/product/repository"
)

type productUsecase struct {
	repo repository.ProductRepo
}

func NewProductUsecase(repo repository.ProductRepo) ProductUsecase {
	return &productUsecase{repo: repo}
}

func (u *productUsecase) GetProductList(category string) ([]model.Products, error) {
	return u.repo.GetProduct(category)
}
