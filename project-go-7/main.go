package main

import (
	"errors"
	"fmt"
)

var salary int = 8000

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return e.msg
}

var err2 = MyError{msg: "Error: el salario es menor a 10.000"}
var err3 = errors.New("Error: el salario es menor a 10.000")
var err4 = fmt.Errorf("Error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)

func main() {

	if salary < 150000 {
		err := MyError{msg: "Error: el salario ingresado no alcanza el mínimo imponible"}
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}

	if err := validar2(salary); err != nil {
		fmt.Println(errors.Is(err, err2)) //compara tipo y valor
	}

	if err := validar3(salary); err != nil {
		fmt.Println(errors.Is(err, err3))
	}

	if err := validar4(salary); err != nil {
		fmt.Println(errors.Is(err, err4))
	}

	salario, err := calcularSalario(40, 4000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario es: %.0f ", salario)
	}

}

func validar2(salary int) error {
	if salary < 10000 {
		return err2
	}
	return nil
}

func validar3(salary int) error {
	if salary < 10000 {
		return err3
	}
	return nil
}

func validar4(salary int) error {
	if salary < 10000 {
		return err4
	}
	return nil
}

func calcularSalario(horas float64, valor float64) (salario float64, err error) {
	if horas*valor < 80000 {
		err = errors.New("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	} else {
		salario = horas * valor
	}
	if horas*valor >= 150000 {
		salario = salario - salario*0.10
	}
	return
}
