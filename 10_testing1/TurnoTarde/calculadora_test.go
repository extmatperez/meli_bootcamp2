package turnotarde

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestRestar(t *testing.T){
	a := 10
	b := 5
	resultadoEsperado := 5

	resultado := Restar(a,b)

	assert.Equal(t,resultadoEsperado,resultado, "Error al restar")

}

func TestRestarNotEqual(t *testing.T){
	a := 10
	b := 5
	resultadoEsperado := 2

	resultado := Restar(a,b)

	assert.NotEqual(t,resultadoEsperado,resultado, "Error al restar")

}


func TestOrdenar(t *testing.T){

	resultadoEsperado:= []int{1,2,3,4,5,6,7,8,9,10}

	resultado := OrdernarSilice(10,4,5,9,7,6,1,2,3,8)

	assert.Equal(t,resultadoEsperado,resultado, "Error al ordenar")
}


func TestOrdenarNotEqual(t *testing.T){

	resultadoEsperado:= []int{10,4,5,9,7,6,1,2,3,8}

	resultado := OrdernarSilice(10,4,5,9,7,6,1,2,3,8)

	assert.NotEqual(t,resultadoEsperado,resultado, "No deberia ser iguales")
}


func TestDividrError(t *testing.T){

	a := 10
	b := 0

	_,err := Dividir(a,b)

	assert.NotNil(t,err,"No se deberia permtir dividir por 0")
}


func TestDividrCorrect(t *testing.T){

	a := 10
	b := 10
	resultadoEsperado := 1

	resultado,_ := Dividir(a,b)

	assert.Equal(t,resultado,resultadoEsperado,"Error")
}
