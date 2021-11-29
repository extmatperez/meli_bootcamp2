package main

import (
	"fmt"
	"math/rand"
	"time"
)


func OrdenamientoInsercion(numeros []int) []int{
	init := time.Now()
	var clave,j int
	for i :=1;i<=len(numeros)-1;i++ {
		clave = numeros[i]
		j = i-1
		for (j >= 0 && clave < numeros[j]){

					numeros[j+1] = numeros[j] 
					numeros[j] = clave				
					j--
		}
	}
	fin := time.Now()
	tiempo := fin.Sub(init).Seconds()
	fmt.Println("El tiempo para ordenar por insercion: ",tiempo)
	return numeros
}

func OrdenamientoBrubuja(ListaDesordenada []int) []int{
	init := time.Now()
	var auxiliar int
	for i := 0; i < len(ListaDesordenada); i++ {
	 for j := 0; j < len(ListaDesordenada); j++ {
	  if ListaDesordenada[i] > ListaDesordenada[j] {
	   auxiliar = ListaDesordenada[i]
	   ListaDesordenada[i] = ListaDesordenada[j]
	   ListaDesordenada[j] = auxiliar
	  }
	 }
	}
	fin := time.Now()

	tiempo := fin.Sub(init).Seconds()

	fmt.Println("El tiempo para ordenar por burbuja: ",tiempo)
	return ListaDesordenada
}





func main() {
	variable1 := rand.Perm(100)
   variable2 := rand.Perm(1000)
   variable3 := rand.Perm(10000)


 	fmt.Println("Variable 1 Tiempo")
	OrdenamientoInsercion(variable1)
	OrdenamientoBrubuja(variable1)

	fmt.Println("Variable 2 Tiempo")
	OrdenamientoInsercion(variable2)
	OrdenamientoBrubuja(variable2)



	fmt.Println("Variable 3 Tiempo")
	OrdenamientoInsercion(variable3)
	OrdenamientoBrubuja(variable3)

}

