package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"io"
)

type product struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
	Active bool `json:"active,omitempty"`
}

func main() {

	products := map[string]any {
		"name": "producto1",
		"price": 2.99,
		"active": false,
	}

	mapAsJson, err := json.Marshal(products)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(mapAsJson))

	producto1 := product{"producto1", 2.99, false}

	structAsJson, err := json.Marshal(producto1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(structAsJson))

	productsJson := `{"name":"producto1","price":2.99}`

	var producto2 product

	err2 := json.Unmarshal([]byte(productsJson), &producto2)
	if err != nil {
		panic(err2)
	}
	fmt.Println(producto2)

	encoder := json.NewEncoder(os.Stdout)

	producto3 := product{"producto3", 3.99, true}

	encoder.Encode(producto3)

	producto4 := `
		{"producto4", 4.99, true}
		{"producto5", 4.99, true}
	`

	decoder := json.NewDecoder(strings.NewReader(producto4))

	for {
		var producto5 product
		if err4 := decoder.Decode(&producto5); err == io.EOF {
			break 
		} else if err != nil {
			panic(err4)
		}
		fmt.Println(producto5)
	}

}