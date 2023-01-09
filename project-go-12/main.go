package main

import (
	"encoding/json"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type Request struct {
	Name        string  `json:"name" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	CodeValue   string  `json:"code_value" validate:"required"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

var Products []Product
var LastId int = 500

func main() {

	readProducts("./products.json", &Products)

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	products := router.Group("/products")
	{
		products.GET("", AllProducts())
		products.GET("/:id", IdProduct())
		products.GET("/search", PriceProducts())
		products.POST("", CreateProduct())
	}

	router.Run(":8080")
}

func readProducts(url string, products *[]Product) {

	file, err := os.ReadFile(url)
	if err != nil {
		panic(err)
	}

	err2 := json.Unmarshal(file, &products)
	if err != nil {
		panic(err2)
	}
}

func AllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": Products})
	}
}

func IdProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "error: invalid id", "data": nil})
			return
		}
		for _, p := range Products {
			if p.Id == id {
				ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": p})
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{"message": "error: product not found", "data": nil})
	}
}

func PriceProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "error: invalid price", "data": nil})
			return
		}
		var filteredProducts []Product
		for _, p := range Products {
			if p.Price > price {
				filteredProducts = append(filteredProducts, p)
			}
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": filteredProducts})
	}
}

func CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Request
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "error: invalid product", "data": nil})
			return
		}
		validate := validator.New()
		if err := validate.Struct(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error: invalid product data", "data": nil})
			return
		}
		if ExistProduct(req) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error: already exit", "data": nil})
			return
		}
		regexDate := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
		if !regexDate.MatchString(req.Expiration) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error: expiration is invalidate", "data": nil})
			return
		}
		LastId++
		Products = append(Products, Product{LastId, req.Name, req.Quantity, req.CodeValue, req.IsPublished, req.Expiration, req.Price})

		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": Products})
	}
}

func ExistProduct(req Request) (result bool) {
	for _, p := range Products {
		if p.CodeValue == req.CodeValue {
			result = true
		}
	}
	return
}
