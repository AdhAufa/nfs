package service

import (
	"github.com/adhaufa/nfs/dto"
	"github.com/adhaufa/nfs/repo"

	_product "github.com/adhaufa/nfs/service/product"
)

type ProductService interface {
	CreateProduct(productRequest dto.ProductRequest) (*_product.ProductResponse, error)
}

type productService struct {
	productRepo repo.ProductRepository
}

func NewProductService(productRepo repo.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (c *productService) CreateProduct(productRequest dto.ProductRequest) (*_product.ProductResponse, error) {
	return nil, nil
}
