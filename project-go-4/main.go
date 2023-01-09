package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{
	{1, "Remera", 356.99, "Remera de algodon", "Ropa"},
	{2, "Destornillador", 560.99, "Destornillador philips", "Herramientas"},
}

func (p *Product) Save() {
	Products = append(Products, *p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Println(product)
	}
}

func getById(id int) (p Product) {
	for _, product := range Products {
		if product.ID == id {
			p = product
			break
		}
	}
	return
}

func main() {
	fmt.Println(getById(2))
	p := Product{3, "Harina", 80.50, "Harina 0000", "Alimentos"}
	p.Save()
	p.GetAll()
}
