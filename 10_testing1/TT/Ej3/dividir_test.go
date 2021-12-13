package dividir_test

import (
	"testing"

	dividir "github.com/extmatperez/meli_bootcamp2/10_testing1/TT/Ej3"
	"github.com/stretchr/testify/assert"
)

func TestDividir1(t *testing.T) {
	a := 5
	b := 1

	div, err := dividir.Dividir(a, b)
	if err != nil {
		assert.NotNil(t, div)
	}
	resultadoEsperado := 5

	assert.Equal(t, resultadoEsperado, div, "Deberian ser 5")
}
func TestDividir2(t *testing.T) {
	a := 5
	b := 0

	_, err := dividir.Dividir(a, b)

	assert.NotNil(t, err)

}
