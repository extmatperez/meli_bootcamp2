package main

import (
	"fmt"
)

const(
	CatA = "A"
	CatB = "B"
	CatC = "C"
)

func main() {


	promedioA :=getSalaryFromCategoryAndHours(CatA,2,15)
	promedioB :=getSalaryFromCategoryAndHours(CatB,5,50)
	promedioC :=getSalaryFromCategoryAndHours(CatC,10,5)

	fmt.Println(promedioA, promedioB,promedioC)
 }


func getSalaryFromCategoryAndHours (categoria string, horas int,salryMontlhy float64)  (float64){
	
	horasfloat := float64(horas)

	switch categoria{
	case CatC:
		return (1000.0 * horasfloat) 
	case CatB:
		return (1500.0 * horasfloat) + (salryMontlhy * (1-0.2) )
	case CatA:
		return (3000 * horasfloat) + (salryMontlhy * (1-0.5) )
	}
	return 0.0
	

}
		

	



