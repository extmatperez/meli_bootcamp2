package main

import (
	"fmt"
	"strings"
)

func main(){
	nombre := "Juan Pablo"



	s := strings.ReplaceAll(nombre, " ", "")
	st := strings.Split(s, "")

	println("La cantidad de letras de este texto es", len(st))

	for i := 0 ; i < len(s) ; i++ {
		fmt.Println("Letra",i,":",st[i])
	}
	

}