package fino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestBase0(t *testing.T){
	numero := 0
	excpected := 0

	resultado := Finobacci(numero)

	assert.Equal(t,excpected,resultado)
}

func TestBase1(t *testing.T){
	numero := 1
	excpected := 1

	resultado := Finobacci(numero)

	assert.Equal(t,excpected,resultado)
}


func TestGetSerie(t *testing.T){
	numero := 10
	excpected := []int{0,1,1,2,3,5,8,13,21,34}

	resultado := GetSerie(numero)

	assert.Equal(t,excpected,resultado)
}

func TestGetSerie0(t *testing.T){
	numero := 0

	resultado := GetSerie(numero)

	assert.Nil(t,resultado)
}

func TestGetSerie1(t *testing.T){
	numero := 1
	excpected := []int{0}

	resultado := GetSerie(numero)

	assert.Equal(t,excpected,resultado)
}