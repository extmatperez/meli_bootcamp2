package main

import (
	"fmt"
	"net/http"
)

func funcionHandlerQuePuedeSerAnonima(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola")
}

func main() {

	http.HandleFunc("/hola", funcionHandlerQuePuedeSerAnonima)
	http.ListenAndServe(":8080", nil)
}
