package handlers

import (
	"net/http"
	"practices/project-go-13/internal/models"
	"practices/project-go-13/internal/products"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AllProducts(ctx *gin.Context) {
	// request
	// process
	// response
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": products.Products})
}

func IdProduct(ctx *gin.Context) {
	// request
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error: invalid id", "data": nil})
		return
	}
	// process
	prod := products.GetById(id)
	// response
	if prod == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "error: product not found", "data": nil})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": &prod})
}

func PriceProducts(ctx *gin.Context) {
	// request
	price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error: invalid price", "data": nil})
		return
	}
	// process
	prods := products.GetByPrice(price)
	// response
	if prods == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "error: products not found", "data": nil})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": prods})
}

func CreateProduct(ctx *gin.Context) {
	// request
	var req models.Request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error: invalid product", "data": nil})
		return
	}
	// process
	if err := products.ValidateProduct(req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error: invalid product data", "data": nil})
		return
	}
	if products.DuplicateProduct(req.CodeValue) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error: already exit", "data": nil})
		return
	}
	if err := products.ValExpirationProd(req.Expiration); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error: expiration is invalidate", "data": nil})
		return
	}
	newProd := products.CreateNewProduct(req)
	// response
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": newProd})
}
