package _product

import (
	"github.com/adhaufa/nfs/entity"
	_user "github.com/adhaufa/nfs/service/user"
)

type ProductResponse struct {
	ID          int64              `json:"id"`
	ProductName string             `json:"product_name"`
	Price       uint64             `json:"price"`
	User        _user.UserResponse `json:"user,omitempty"`
}

func NewProductResponse(product entity.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		ProductName: product.Name,
		Price:       product.Price,
		User:        _user.NewUserResponse(product.User),
	}
}

func NewProductArrayResponse(products []entity.Product) []ProductResponse {
	productRes := []ProductResponse{}
	for _, v := range products {
		p := ProductResponse{
			ID:          v.ID,
			ProductName: v.Name,
			Price:       v.Price,
			User:        _user.NewUserResponse(v.User),
		}
		productRes = append(productRes, p)
	}
	return productRes
}
