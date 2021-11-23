package main

import "fmt"

func main(){
	var temp float32 = 25.3
	var hum int = 50
	var pres float32 = 992.5

	fmt.Printf("La temperatura es de %v grados\n", temp)
	fmt.Printf("La humedad es del %v%%\n" , hum)
	fmt.Printf("La presion es de %vhPa \n" , pres)
}
