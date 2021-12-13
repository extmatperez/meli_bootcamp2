package ordenamiento

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	input := []int{1, 7, 3, 4}

	resultado := Ordenar(input)
	fmt.Println(input)
	fmt.Println(resultado)

	assert.Equal(t, []int{1, 3, 4, 7}, resultado)
	assert.IsIncreasing(t, resultado)

}
