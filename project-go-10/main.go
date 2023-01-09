package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func main() {

	//Creación de servidor basico
	/*pingHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	router := http.NewServeMux()
	router.Handle("/ping", pingHandler)
	server := &http.Server{
		Addr: ":8080",
		WriteTimeout: 10 * time.Second,
		Handler: router,
	}
	server.ListenAndServe()*/

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	router.POST("/saludo", func(ctx *gin.Context) {

		var person1 Person

		if err := ctx.Bind(&person1); err != nil {
			ctx.String(http.StatusBadRequest, "Error")
			return
		}
		mensaje := fmt.Sprintf("Hola %s %s", person1.Nombre, person1.Apellido)
		ctx.String(http.StatusOK, mensaje)
	})

	router.GET("/next/:number", func(ctx *gin.Context) {
		// request
		num, err4 := strconv.Atoi(ctx.Param("number"))
		if err4 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "faild to parse number",
				"data":    nil,
			})
			return
		}
		// process
		num++
		// response
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success to increment number",
			"data":    num,
		})
	})

	router.Run(":8080")
}
