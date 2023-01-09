package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(ImpuestosSalario(200000))
	fmt.Println(CalcularPromedio(9, 8, 10, 7))
	fmt.Println(CalcularSalario(30, "Categoría A"))

	operacion, err := CalcularEstadisticas("mínimo")
	if err != nil {
		fmt.Println(err)
	} else {
		resultado := operacion(9, 8, 10, 7)
		fmt.Println(resultado)
	}

	operacion2, err := Animal("Gato")
	if err != nil {
		fmt.Println(err)
	} else {
		resultado2 := operacion2(10)
		fmt.Println(resultado2)
	}
}

// Ejercicio 1
func ImpuestosSalario(salario float64) (impuesto float64) {
	impuesto = 0
	if salario > 150000 {
		impuesto = salario * 0.27
	} else if salario > 50000 {
		impuesto = salario * 0.17
	}
	return
}

// Ejercicio 2
func CalcularPromedio(notas ...uint) float64 {
	var suma uint
	for _, nota := range notas {
		suma += nota
	}
	return float64(suma) / float64(len(notas))
}

// Ejercicio 3
func CalcularSalario(minutos float64, categoria string) (salario float64) {
	salario = 0
	if categoria == "Categoría C" {
		salario = (minutos / 60) * 1000
	} else if categoria == "Categoría B" {
		subtotal := (minutos / 60) * 1500
		salario = subtotal + subtotal*0.2
	} else if categoria == "Categoría A" {
		subtotal := (minutos / 60) * 3000
		salario = subtotal + subtotal*0.5
	}
	return salario
}

// Ejercicio 4
type Operaciones = func(notas ...uint) float64

func CalcularEstadisticas(calculo string) (Operaciones, error) {
	switch calculo {
	case "mínimo":
		return opMinimo, nil
	case "máximo":
		return opMaximo, nil
	case "promedio":
		return opPromedio, nil
	default:
		return nil, errors.New("no es un calculo correcto")
	}
}

func opMinimo(notas ...uint) float64 {
	min := notas[0]
	for _, nota := range notas {
		if min > nota {
			min = nota
		}
	}
	return float64(min)
}
func opMaximo(notas ...uint) float64 {
	max := notas[0]
	for _, nota := range notas {
		if max < nota {
			max = nota
		}
	}
	return float64(max)
}
func opPromedio(notas ...uint) float64 {
	var suma uint
	for _, nota := range notas {
		suma += nota
	}
	return float64(suma) / float64(len(notas))
}

// Ejercicio 5
type Animales func(cantidad int) float64

func Animal(animal string) (Animales, error) {
	switch animal {
	case "Perro":
		return animalPerro, nil
	case "Gato":
		return animalGato, nil
	case "Hamster":
		return animalHamster, nil
	case "Tarántula":
		return animalTarantula, nil
	default:
		return nil, errors.New("no es un animal correcto")
	}
}

func animalPerro(cantidad int) float64 {
	return float64(cantidad) * 10.0
}
func animalGato(cantidad int) float64 {
	return float64(cantidad) * 5.0
}
func animalHamster(cantidad int) float64 {
	return float64(cantidad) * 0.25
}
func animalTarantula(cantidad int) float64 {
	return float64(cantidad) * 0.15
}
