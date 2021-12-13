package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	n1 := 1.0
	n2 := 1.0
	res, _ := Dividir(n1, n2)
	assert.Equal(t, res, 1.0)
	//assert.Equal(t, errors.New("el denomiador no puede ser 0"), err)
}

func TestDividir2(t *testing.T) {
	n1 := 1.0
	n2 := 0.0
	_, err := Dividir(n1, n2)
	erorEsperado := "el denomiador no puede ser 0"
	assert.EqualErrorf(t, err, erorEsperado, "el error deber ser: %v, y tenemos: %v", erorEsperado, err)
}
