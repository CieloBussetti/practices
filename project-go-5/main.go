package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person
}

func (e *Employee) PrintEmployee() {
	fmt.Printf("El empleado %s, es %s y cumple años el %s \n", e.Person.Name, e.Position, e.Person.DateOfBirth)
}

func main() {
	e := Employee{1, "repartidor", Person{1, "Carlos", "22/09/2021"}}
	e.PrintEmployee()
}
