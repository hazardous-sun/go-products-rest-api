package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type ProductUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		productRepository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.productRepository.GetProducts()
}

func (pu *ProductUsecase) GetProductById(id int) (*model.Product, error) {
	product, err := pu.productRepository.GetProductByID(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	id, err := pu.productRepository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	return model.Product{
		ID:    id,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
