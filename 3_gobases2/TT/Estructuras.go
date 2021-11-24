// Por ejemplo, cuando trabajemos con REST nos puede hacer falta hacer etiquetas.

package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type Persona struct {
	Nombre   string  `json:"nombre_persona"`
	Apellido string  `json:"apellido_persona"`
	Edad     int     `json:"edad_persona"`
	Altura   float64 `json:"altura_persona"`
	Peso     float64 `json:"peso_persona"`
	IMC      float64 `json:"imc_persona"`
}

// Esto es la representacion de los metodos de cada estructura, para simular programacion orientada a objetos.
func (persona Persona) calcular_imc() float64 {
	return persona.Peso / math.Pow(2, float64(persona.Altura))
}

func main() {
	p0 := Persona{
		Nombre:   "Rodrigo",
		Apellido: "Vega Gimenez",
		Edad:     25,
		Altura:   1.82,
		Peso:     92,
	}
	miJSON, _ := json.Marshal(p0)
	fmt.Println("\n", miJSON)
	p0.IMC = p0.calcular_imc()
	fmt.Println(p0)
}
