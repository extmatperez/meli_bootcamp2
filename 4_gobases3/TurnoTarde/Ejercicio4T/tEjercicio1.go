package main

import (
	"fmt"
	"math/rand"
)


func OrdenamientoInsercion(numeros []int) []int{
	var clave,j int
	for i :=1;i<=len(numeros)-1;i++ {
		clave = numeros[i]
		j = i-1
		fmt.Println(i,clave)
		fmt.Println(j,	numeros[j])
	
		for (j >= 0 && clave < numeros[j]){
					fmt.Println("Clave",clave,"<",numeros[j])
					numeros[j+1] = numeros[j] 
					numeros[j] = clave				
					j--
		}
		
		fmt.Println(numeros)
	}
	return numeros
}

func main() {
	variable1 := rand.Perm(100)

	
	fmt.Println(variable1)
	fmt.Println(OrdenamientoInsercion(variable1))

}

