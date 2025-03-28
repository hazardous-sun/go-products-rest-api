package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	id, err := pu.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	return model.Product{
		ID:    id,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
