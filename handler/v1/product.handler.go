package v1

import (
	"github.com/adhaufa/nfs/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	CreateProduct(ctx *gin.Context)
	FindOneProductByID(ctx *gin.Context)
}

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return &productHandler{
		productService: productService,
	}
}

func (c *productHandler) CreateProduct(ctx *gin.Context) {

}

func (c *productHandler) FindOneProductByID(ctx *gin.Context) {

}
