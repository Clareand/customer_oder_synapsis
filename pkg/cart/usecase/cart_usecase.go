package usecase

import (
	"github.com/Clareand/customer_oder_synapsis/pkg/cart/model"
	"github.com/Clareand/customer_oder_synapsis/pkg/cart/repository"
)

type cartUsecase struct {
	repo repository.CartRepo
}

func NewCartUsecase(repo repository.CartRepo) CartUsecase {
	return &cartUsecase{repo: repo}
}

func (u *cartUsecase) GetCart(customer_id string) ([]model.Carts, error) {
	return u.repo.GetCart(customer_id)
}

func (u *cartUsecase) AddToCart(param model.AddToCart) error {
	return u.repo.AddToCart(param)
}
