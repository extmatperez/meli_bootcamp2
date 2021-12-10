package calculador_test

import (
	"testing"

	calculador "github.com/extmatperez/meli_bootcamp2/10_testing1/TT/Ej1"
)

func TestResta(t *testing.T) {
	res := calculador.Restar(1, 2)
	if res != -1 {
		t.Errorf("error: La resta no es correcta")
		t.Fail()
	} else {
		t.Log("success: El test ha pasado correctamente")
	}
}
