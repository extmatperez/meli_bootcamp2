package main

import "fmt"

type Matrix struct {
	valores []float64
	dim_al int
	dim_an int
	tipo string
	max float64
}

func (m *Matrix) set_valores(valores []float64){
	m.valores = valores
}

func (m Matrix) print_matrix(){
	cont := 0
	for i := 0; i<m.dim_al; i++{
		for j := 0; j<m.dim_an; j++{
			fmt.Printf(" %v ", m.valores[cont])
			cont += 1
		}
		fmt.Printf("\n")
	}
}
func (m *Matrix) maximo(){
	maximo := 0.0
	for i := 0; i<len(m.valores); i++{
		if m.valores[i] > maximo{
			maximo = m.valores[i]
			m.max = maximo
		}
	}
}
func main(){

	values := make([]float64, 0)
	var add float64

	m := Matrix{
		dim_al: 3,
		dim_an: 3,
		tipo: "Cuadrado",
	}

	for i := 0; i<m.dim_al*m.dim_an;i++{
		fmt.Printf("Ingrese valor:")
		fmt.Scanln(&add)
		if add == 0 {
		break	
		}
		values = append(values, add)
	}
	m.set_valores(values)
	m.maximo()
	m.print_matrix()
	fmt.Println(m)


}