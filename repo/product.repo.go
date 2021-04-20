package repo

import (
	"github.com/adhaufa/nfs/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProduct(product entity.Product) (entity.Product, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
	FindOneProductByID(ID string) (entity.Product, error)
	FindAllProduct(userID string) ([]entity.Product, error)
}

type productRepo struct {
	connection *gorm.DB
}

func NewProductRepo(connection *gorm.DB) ProductRepository {
	return &productRepo{
		connection: connection,
	}
}

func (c *productRepo) InsertProduct(product entity.Product) (entity.Product, error) {
	return entity.Product{}, nil
}

func (c *productRepo) UpdateProduct(product entity.Product) (entity.Product, error) {
	return entity.Product{}, nil
}

func (c *productRepo) FindOneProductByID(ID string) (entity.Product, error) {
	return entity.Product{}, nil
}

func (c *productRepo) FindAllProduct(userID string) ([]entity.Product, error) {
	return []entity.Product{}, nil
}
