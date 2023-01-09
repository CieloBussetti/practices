package main

import (
	"fmt"
)

type Persona struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
	Preferidos
}

func (p *Persona) Detail() {
	fmt.Printf("%s %s tiene %d años de edad, su color preferido es %s y su número preferido es %d \n", p.Nombre, p.Apellido, p.Edad, p.Preferidos.Color, p.Preferidos.Numero)
}

type Preferidos struct {
	Color  string
	Numero int
}

type Cliente struct {
	Persona
	Numero int
}

type Item struct {
	Nombre string
	Precio float64
}

type Shop struct {
	storage       []Item
	limite        int
	transacciones int
}

func (s *Shop) CargaStock(items ...Item) {
	s.storage = append(s.storage, items...)
	s.transacciones++
}

func main() {
	p := Persona{"Carlos", "Perez", 22, Preferidos{"verde", 26}}
	c := Cliente{Persona{"Carlos", "Perez", 22, Preferidos{"verde", 26}}, 3829}
	p.Detail()
	fmt.Printf("%v+", c)

	storage := make([]Item, 0)
	shop := &Shop{
		limite:  3,
		storage: storage,
	}
	shop.CargaStock(Item{"Manzana", 0})
}
