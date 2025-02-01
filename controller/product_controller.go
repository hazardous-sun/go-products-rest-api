package model

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/model"
	"net/http"
)

type productController struct {
	// Use case
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata frita",
			Price: 20,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
