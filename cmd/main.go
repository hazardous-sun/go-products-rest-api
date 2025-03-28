package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/usecase"
	"log"
)

func main() {
	server := gin.Default()

	// Connecting to the database
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	// Repository layer
	productRepository := repository.NewProductRepository(dbConnection)

	// Use case layer
	productUsecase := usecase.NewProductUsecase(productRepository)

	// Controller layer
	productController := controller.NewProductController(productUsecase)

	// Setting up HTTP calls handling ----------------------------------------------------------------------------------
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Create
	server.POST("/product", productController.CreateProduct)

	// Read
	server.GET("/products", productController.GetProducts)
	server.GET("/product/:productId", productController.GetProductById)

	// Update
	server.POST("/updateProduct", productController.UpdateProduct)
	// -----------------------------------------------------------------------------------------------------------------

	err = server.Run(":8000")

	if err != nil {
		log.Panicln(err)
		return
	}
}
