package main

import (
	"errors"
	"fmt"
)

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Printf("Alumno %s %s, DNI %d, ingresó el %s \n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

const (
	pequenio = "Pequeño"
	mediano  = "Mediano"
	grande   = "Grande"
)

type Item struct {
	tipo   string
	precio float64
}

type Producto interface {
	Precio() (float64, error)
}

func Factory(tipo string, precio float64) Producto {
	return Item{tipo: tipo, precio: precio}
}

func (p Item) Precio() (float64, error) {
	switch p.tipo {
	case pequenio:
		return p.precio, nil
	case mediano:
		total := p.precio + p.precio*0.3
		return total, nil
	case grande:
		total := p.precio + p.precio*0.6 + 2500.0
		return total, nil
	}
	return 0, errors.New("Error")
}

func main() {
	a := Alumno{"Maria", "Perez", 40980399, "20/03/2021"}
	a.detalle()

	p := Factory(pequenio, 405.90)
	precio, err := p.Precio()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(precio)
}
