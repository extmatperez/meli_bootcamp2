// Por ejemplo, cuando trabajemos con REST nos puede hacer falta hacer etiquetas.

package main

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
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
func (persona *Persona) calcular_imc() float64 {
	return persona.Peso / math.Pow(2, float64(persona.Altura))
}

func (persona *Persona) setEdad(age int) {
	persona.Edad = age
}

func (persona Persona) getDetalle() {
	fmt.Printf("nombre_persona: \t%s\n, apellido_persona: \t%s\n, edad_persona: \t%d\n, altura_persona: \t%f\n, peso_persona: \t%f\n, imc_persona: \t%2.2f\n", persona.Nombre, persona.Apellido, persona.Edad, persona.Altura, persona.Peso, persona.IMC)
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
	fmt.Println("\n", string(miJSON))
	p0.IMC = p0.calcular_imc()
	fmt.Println(p0.Edad)
	p0.setEdad(27)
	fmt.Println(p0.Edad)
	fmt.Println(p0)
	p := reflect.TypeOf(p0)
	fmt.Println("Type:", p.Name())
	fmt.Println("Kind: ", p.Kind())
	fmt.Println("Cantidad de campos: ", p.NumField())
	fmt.Println("Contenido del primer campo: ", p.Field(0))
	p0.getDetalle()
}
