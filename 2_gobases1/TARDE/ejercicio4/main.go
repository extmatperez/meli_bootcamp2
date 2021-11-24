package main

import "fmt"

func main(){
	var meses = map[int]string{1: "enero", 2: "febrero", 3: "marzo", 4: "abril", 5: "mayo", 6: "junio",
	7: "julio", 8: "agosto", 9: "septiembre", 10: "octubre", 11: "nobiembre", 12: "diciembre",
}

fmt.Println("el mes 7 es: ", meses[7])

}