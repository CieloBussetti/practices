package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImpuestosSalario(t *testing.T) {
	//Arrage
	num := 40000.0
	expectedResult := 0.0
	//Act
	result := ImpuestosSalario(num)
	//Assert
	assert.Equal(t, expectedResult, result)
}
func TestImpuestosSalario_RangoMas50000(t *testing.T) {
	//Arrage
	num := 70000.0
	expectedResult := 11900.0
	//Act
	result := ImpuestosSalario(num)
	//Assert
	assert.Equal(t, expectedResult, result)
}
func TestImpuestosSalario_RangoMas150000(t *testing.T) {
	//Arrage
	num := 200000.0
	expectedResult := 54000.0
	//Act
	result := ImpuestosSalario(num)
	//Assert
	assert.Equal(t, expectedResult, result)
}

func TestCalcularPromedio(t *testing.T) {
	//Arrage
	notas := []uint{9, 8, 5, 10}
	expectedResult := 8.0
	//Act
	result := CalcularPromedio(notas...)
	//Assert
	assert.Equal(t, expectedResult, result)
}

func TestCalcularSalario_CategoriaB(t *testing.T) {
	//Arrage
	minutos := 370.0
	categoria := "Categoría B"
	expectedResult := 11100.0
	//Act
	result := CalcularSalario(minutos, categoria)
	//Assert
	assert.Equal(t, expectedResult, result)
}

func TestCalcularSalario_CategoriaC(t *testing.T) {
	//Arrage
	minutos := 360.0
	categoria := "Categoría C"
	expectedResult := 6000.0
	//Act
	result := CalcularSalario(minutos, categoria)
	//Assert
	assert.Equal(t, expectedResult, result)
}

func TestCalcularSalario_CategoriaA(t *testing.T) {
	//Arrage
	minutos := 370.0
	categoria := "Categoría A"
	expectedResult := 27750.0
	//Act
	result := CalcularSalario(minutos, categoria)
	//Assert
	assert.Equal(t, expectedResult, result)
}

func TestCalcularEstadisticas_Minimo(t *testing.T) {
	//Arrage
	calculo := "mínimo"
	notas := []uint{9, 8, 10, 7}
	expectedResult := 7.0
	//Act
	preResult, _ := CalcularEstadisticas(calculo)
	result := preResult(notas...)
	//Assert
	assert.Equal(t, expectedResult, result)
}

func TestCalcularEstadisticas_Maximo(t *testing.T) {
	//Arrage
	calculo := "máximo"
	notas := []uint{9, 8, 10, 7}
	expectedResult := 10.0
	//Act
	preResult, _ := CalcularEstadisticas(calculo)
	result := preResult(notas...)
	//Assert
	assert.Equal(t, expectedResult, result)
}

func TestCalcularEstadisticas_Promedio(t *testing.T) {
	//Arrage
	calculo := "promedio"
	notas := []uint{9, 8, 10, 7}
	expectedResult := 8.5
	//Act
	preResult, _ := CalcularEstadisticas(calculo)
	result := preResult(notas...)
	//Assert
	assert.Equal(t, expectedResult, result)
}

func TestAnimalPerro(t *testing.T) {
	//Arrage
	cantidad := 12
	expectedResult := 120.0
	//Act
	result := animalPerro(cantidad)
	//Assert
	assert.Equal(t, expectedResult, result)
}

func TestAnimalGato(t *testing.T) {
	//Arrage
	cantidad := 6
	expectedResult := 30.0
	//Act
	result := animalGato(cantidad)
	//Assert
	assert.Equal(t, expectedResult, result)
}

func TestAnimalHamster(t *testing.T) {
	//Arrage
	cantidad := 25
	expectedResult := 6.25
	//Act
	result := animalHamster(cantidad)
	//Assert
	assert.Equal(t, expectedResult, result)
}

func TestAnimalTarantula(t *testing.T) {
	//Arrage
	cantidad := 20
	expectedResult := 3.0
	//Act
	result := animalTarantula(cantidad)
	//Assert
	assert.Equal(t, expectedResult, result)
}
