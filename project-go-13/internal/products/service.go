package products

import (
	"encoding/json"
	"os"
	"practices/project-go-13/internal/models"
	"time"

	"github.com/go-playground/validator/v10"
)

var Products []models.Product
var LastId int

func ReadProducts(url string) error {

	file, err := os.ReadFile(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &Products)
	if err != nil {
		return err
	}
	LastId = 500
	return nil
}

func GetById(id int) *models.Product {
	for _, p := range Products {
		if p.Id == id {
			return &p
		}
	}
	return nil
}

func GetByPrice(price float64) (prods []models.Product) {
	for _, p := range Products {
		if p.Price > price {
			prods = append(prods, p)
		}
	}
	return
}

func ValidateProduct(req models.Request) (err error) {
	validate := validator.New()
	err = validate.Struct(&req)
	return
}

func DuplicateProduct(code string) (result bool) {
	for _, p := range Products {
		if p.CodeValue == code {
			result = true
		}
	}
	return
}

func ValExpirationProd(expiration string) (err error) {
	_, err = time.Parse("02/01/2006", expiration)
	return
}

func CreateNewProduct(req models.Request) (prod models.Product) {
	prod = models.Product{
		Id:          LastId + 1,
		Name:        req.Name,
		Quantity:    req.Quantity,
		CodeValue:   req.CodeValue,
		IsPublished: req.IsPublished,
		Expiration:  req.Expiration,
		Price:       req.Price,
	}
	Products = append(Products, prod)
	LastId++
	return
}
