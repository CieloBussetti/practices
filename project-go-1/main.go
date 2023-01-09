package main

import (
	"fmt"
)

func main() {

	nombre := "Cielo"
	var direccion = "Av. Santa Fe 4462"
	fmt.Println("Mi nombre es", nombre, "y vivo en", direccion)

	var temperatura int8 = 26
	var humedad uint = 63
	var presion float32 = 1015
	fmt.Println("Temperatura:", temperatura, "Humedad:", humedad, "Presión:", presion)

	num1 := 8
	num2 := 6

	if num1 > num2 {
		fmt.Println("mayor el", num1)
	} else if num1 < num2 {
		fmt.Println("mayor el", num2)
	} else {
		fmt.Println("son iguales")
	}

	if n1, n2 := 7, 6; n1 == n2 {
		fmt.Println("son iguales")
	} else {
		fmt.Println("no son iguales")
	}

	type empleado struct {
		name           string
		edad           int
		estaEmpleada   bool
		diasAntiguedad int
		saldo          float64
	}
	var empleados = []empleado{
		{"Maria", 25, true, 558, 12000.5},
		{"Carlos", 40, false, 0, 0.0},
		{"Pablo", 40, true, 369, 209000.0},
	}
	for _, empleado := range empleados {
		if empleado.edad >= 22 && empleado.estaEmpleada && empleado.diasAntiguedad > 365 {
			if empleado.saldo > 100000 {
				fmt.Println("Se te puede otorgar un prestamo sin interés")
			} else {
				fmt.Println("Se te puede otorgar un prestamo con interés")
			}
		} else {
			fmt.Println("No se te puede otorgar un prestamo")
		}
	}

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var mayores int
	for i, employee := range employees {
		if i == "Benjamin" {
			fmt.Println("la edad de Benjamin es:", employee)
		}
		if employee >= 21 {
			mayores++
		}
	}
	fmt.Println("Hay", mayores, "empleados mayores")
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)

	//array
	a := [3]string{"nombre", "apellido"}
	a[2] = "edad"
	fmt.Println(a)
	fmt.Printf("Letra 4 de la palabra: %c\n", a[1][4])
	for _, value := range a[1] {
		fmt.Printf("%c\n", value)
	}

	//slice
	s := []string{"nombre", "apellido", "edad", "direccion"}
	s = append(s, "telefono")
	s = append(s[:2], s[3:]...)
	fmt.Println(s)

	//map
	m := map[string]int{"Maria": 20, "Carlos": 25}
	m["Clara"] = 60
	delete(m, "Maria")
	fmt.Println(m)
	for key, value := range m {
		fmt.Printf("%s %d\n", key, value)
	}
	if valor, ok := m["Pablo"]; !ok {
		fmt.Println("El valor no existe")
	} else {
		fmt.Println(valor)
	}

	//llamados a funciones
	meses4 := []string{"Diciembre", "Mayo", "Abril", "Septiembre"}
	ObtenerEstacion(meses4)

	mes := 12
	fmt.Println(
		ObtenerMes(mes),
		ObtenerMes2(mes),
		ObtenerDias(mes),
	)

	colores := []string{"Verde", "Azul", "Naranja", "Rojo", "Amarillo", "Blanco"}
	num := 6
	fmt.Println(
		TipoFor(num),
		TipoWhile(num),
	)
	TipoForeach(colores)

	var palabra string = "Hola"
	TipoFor2(palabra)
}

func ObtenerEstacion(meses4 []string) {
	for _, mes := range meses4 {
		switch mes {
		case "Enero", "Febrero", "Marzo":
			fmt.Println("Verano")
		case "Abril", "Mayo", "Junio":
			fmt.Println("Otoño")
		case "Julio", "Agosto", "Septiembre":
			fmt.Println("Invierno")
		case "Octubre", "Noviembre", "Diciembre":
			fmt.Println("Primavera")
		default:
			fmt.Println("No existe este mes")
		}
	}
}

func ObtenerMes(mes int) string {
	var mesTexto string
	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	for i, cadaMes := range meses {
		if mes < 1 || mes > 12 {
			return "No es un mes valido"
		} else if i+1 == mes {
			mesTexto = cadaMes
			break
		}
	}
	return mesTexto
}

func ObtenerMes2(mes int) string {
	var mesTexto string
	switch mes {
	case 1:
		mesTexto = "Enero"
	case 2:
		mesTexto = "Febrero"
	case 3:
		mesTexto = "Marzo"
	case 4:
		mesTexto = "Abril"
	case 5:
		mesTexto = "Mayo"
	case 6:
		mesTexto = "Junio"
	case 7:
		mesTexto = "Julio"
	case 8:
		mesTexto = "Agosto"
	case 9:
		mesTexto = "Septiembre"
	case 10:
		mesTexto = "Octubre"
	case 11:
		mesTexto = "Noviembre"
	case 12:
		mesTexto = "Diciembre"
	default:
		mesTexto = "No es un mes valido"
	}
	return mesTexto
}

func ObtenerDias(mes int) string {
	//swich -> No hace falta break, termina en el primer match que logra
	var mesDias string
	switch mes := 6; mes {
	case 4, 6, 9, 11:
		mesDias = "El mes tiene 30 días"
	case 1, 3, 5, 7, 8, 10, 12:
		mesDias = "El mes tiene 31 días"
	case 2:
		mesDias = "El mes tiene 28 días"
	default:
		mesDias = "No es un mes valido"
	}
	return mesDias
}

// for
func TipoFor(num int) int {
	var sum int
	for i := 0; i < num; i++ {
		sum += i
	}
	return sum
}

func TipoFor2(palabra string) {
	fmt.Println("La cantidad de letras de la palabra es:", len(palabra))
	for i := 0; i < len(palabra); i++ {
		fmt.Printf("%c\n", palabra[i])
	}
}

// while
func TipoWhile(num int) int {
	var sum = 1
	for sum < num {
		sum += sum
	}
	return sum
}

// foreach
func TipoForeach(colores []string) {
	for _, color := range colores {
		if color == "Azul" {
			continue
		}
		if color == "Blanco" {
			break
		}
		fmt.Println(color)
	}
}
