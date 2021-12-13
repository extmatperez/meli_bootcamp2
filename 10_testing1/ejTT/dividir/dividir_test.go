package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	resultado, err := Dividir(8.0, 4)
	_, errorDivision := Dividir(3.0, 0)
	assert.Equal(t, resultado, 2.0, "la division debe ser igual a 2 ")
	assert.Nil(t, err, "no tendria que lanzar ningun error diviendo dos enteros")
	assert.NotNil(t, errorDivision, "tiene que lanzar error cuando se divide por 0")
}
