package main

import (
	"errors"
	"fmt"
)

func main() {
	notas := []float64 {}

	promedio,e:=getPromedio(notas)

	fmt.Println(promedio,e)
 }


func getPromedio(notas []float64) (float64,error){
	cantidadNotas := len(notas)

	if(cantidadNotas == 0){
		return 0,errors.New("Tiene que haber al menos una nota para calcular el promedioio.")
		}else{
			sum := 0.0
			for _, nota := range notas {
				sum += nota
			}
			return sum/float64(cantidadNotas),nil
		}
		

	


}
