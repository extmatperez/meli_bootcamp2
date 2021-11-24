package main

import "fmt"

func main() {
	var month_map = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septemba", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}

	fmt.Printf(month_map[11])
}
