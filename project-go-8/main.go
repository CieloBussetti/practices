package main

import (
	"errors"
	"fmt"
	"os"
)

type Customer struct {
	Legajo    int
	Nombre    string
	DNI       int
	Telefono  int
	Domicilio string
}

var Customers = []Customer{
	{18372, "Maria Perez", 20394859, 1148929839, "Av. Cordoba 345"},
	{18372, "Carlos Gonzalez", 40586938, 113847583, "Av. Santa fe 7653"},
}

func main() {
	file := "./customer.txt"
	c := Customer{18372, "Clara Perez", 40586938, 1148929839, "Av. Cordoba 345"}

	data := readFile(file)
	fmt.Printf("%s \n", data)
	registerCustomer(c)
	fmt.Println(Customers)

	defer func() {
		fmt.Println("Ejecución finalizada")
		fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	}()
}

func readFile(file string) []byte {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	data, err := os.ReadFile(file)
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
	return data
}

func validateCustomer(c Customer) (exist bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	for _, customer := range Customers {
		if customer.DNI == c.DNI {
			exist = true
			panic("Error: el cliente ya existe")
		}
	}
	return
}

func validateNulls(c Customer) (bool, error) {
	if c.Legajo == 0 || c.Nombre == "" || c.DNI == 0 || c.Telefono == 0 || c.Domicilio == "" {
		return false, errors.New("Error")
	}
	return true, nil
}

func registerCustomer(c Customer) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	exist := validateCustomer(c)
	if !exist {
		_, err := validateNulls(c)
		if err != nil {
			panic(err)
		} else {
			Customers = append(Customers, c)
		}
	}
}
