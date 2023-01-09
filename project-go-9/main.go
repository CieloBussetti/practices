package main

import (
	"fmt"
	"os"
)

func main() {

	lec, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Print((string(lec)))

	path := os.Getenv("PATH")
	fmt.Print(path)

	err2 := os.Setenv("COSO", "VALOR")
	if err2 != nil {
		panic(err)
	}
	fmt.Print(os.Getenv("COSO"))

	err3 := os.WriteFile("test.txt", []byte("Hola mundo"), 0644)
	if err3 != nil {
		panic(err3)
	}

	file, err4 := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	file.Write([]byte("\nHola mundo 2"))
	if err4 != nil {
		panic(err4)
	}

}
