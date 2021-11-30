package main

import (
	"fmt"
	"net/http"
)

func getAll2(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hola")

}
