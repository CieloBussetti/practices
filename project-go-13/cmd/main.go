package main

import (
	"net/http"
	"practices/project-go-13/cmd/handlers"
	"practices/project-go-13/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {

	err := products.ReadProducts("./products.json")
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	//test
	router.GET("/ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })

	products := router.Group("/products")
	products.GET("", handlers.AllProducts)
	products.GET("/:id", handlers.IdProduct)
	products.GET("/search", handlers.PriceProducts)
	products.POST("", handlers.CreateProduct)

	router.Run(":8080")
}
