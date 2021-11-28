package main

import (
	"encoding/json"
	"fmt"
	"sort"

)

type Matrix struct {
	Valores []float64
	Alto int
	Ancho int
	IsCudratica bool
	ValorMaximo float64
} 
func (m *Matrix) Set(valores ... float64) {
	m.Valores=valores
	
	if(m.Alto == m.Ancho){
		m.IsCudratica = true
	}else{
		m.IsCudratica = false
	}
	sort.Float64s(valores)
	m.ValorMaximo = valores[len(valores)-1]


}

func (m Matrix)  Print() {
	contadorColumna := 0


	fmt.Println("Matrix representada")
	for i :=0; i < len(m.Valores); i++{
		contadorColumna++
		fmt.Printf("%v\t", m.Valores[i])		
		if(contadorColumna == m.Ancho){
			fmt.Printf("\n")
			contadorColumna=0
		}
			
	}
	fmt.Printf("\n")
	fmt.Println("Matrix Datos")
	jsonR,_ := json.Marshal(m)
	fmt.Printf("%v\n",string(jsonR))

}

func main() {
	matriz1 := Matrix{Alto: 5,Ancho: 3}
	matriz1.Set(100, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3,1000,12,15,8)
	matriz1.Print()
}

