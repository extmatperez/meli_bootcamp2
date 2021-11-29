package main

import (
	"fmt"
	"net/http"
)

func handlerSaludar(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hola Mundo!")
}

func main() {
	http.HandleFunc("/Saludar", handlerSaludar)
	http.ListenAndServe(":8080", nil)
}
