package ordenamiento

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestOrdenamiento(t *testing.T) {
	num := rand.Perm(10)
	sort.Ints(num)

	resultado_esperado := Ordenar(num)

	assert.Equal(t, resultado_esperado, num, "El resultado no es el mismo")

}