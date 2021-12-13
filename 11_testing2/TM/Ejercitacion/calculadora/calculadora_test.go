package calculadora

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Aqui en el archivo de testing, se crea un struct dummyLogger.
type dummyLogger struct{}

// Se crea tambien un struct para stub, un stubLogger.
type stubLogger struct{}

// Se crea tambien un struct para spy, un spyLogger.
type spyLogger struct {
	spyCalled bool
}

// Se crea tambien un struct para el mock, un mockConfig.
type mockConfig struct {
	clienteUsado string
}

// Se crea un struct fakeConfig que implemente una logica en la que solo habilita la suma al cliente llamado "John Doe".
type fakeConfig struct{}

// Se escriben las funciones necesarias para que dummyLogger cumpla con la interfaz que va a reemplazar (Logger).
func (d *dummyLogger) Log(str string) error {
	return nil
}

// Se escriben las funciones necesarias para que stubLogger retorne exactamente lo que necesitamos.
func (s *stubLogger) Log(str string) error {
	return errors.New("Error desde stub.")
}

// Para espiar creamos un spyLogger que setea en true spyCalled si entra al metodo.
func (s *spyLogger) Log(str string) error {
	s.spyCalled = true
	return nil
}

// Para implementar el test del mock, el mock debe implementar el metodo necesario y comprobar que SumaEnabled sea llamado y que se haga exactamente
// con el mismo cliente que recibio SumaRestricted.
func (m *mockConfig) SumaEnabled(cliente string) bool {
	m.clienteUsado = cliente
	return true
}

// Se sigue en la creacion del fakeConfig.
func (f *fakeConfig) SumaEnabled(cliente string) bool {
	return cliente == "John Doe"
}

func TestSumar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	// Este es el ARRANGE.
	n1 := 3
	n2 := 5
	promisedValue := 8

	// Se genera el objeto dummy a usar para satisfacer la necesidad de la funcion Sumar.
	myDummy := &dummyLogger{}

	// Se ejecuta el test. Este es el ACT.
	res := SumarDummy(n1, n2, myDummy)

	// Se validan los resultados. Este es el ASSERT.
	// Usar assert es lo mismo que validar la correspondencia y responder con el resultado del test.
	/*if res != promisedValue {
		t.Errorf("Funcion Suma() arrojo el resultado = %v, pero el esperado es %v", res, promisedValue)
	}*/
	assert.Equal(t, promisedValue, res, fmt.Sprintf("Funcion Suma() arrojo el resultado %v, pero el esperado es %v", res, promisedValue))
}

func TestSumarError(t *testing.T) {
	// Se inicializan los datos a usar.
	n1 := 3
	n2 := 5
	promisedValue := -99999

	// Se genera el objeto stub a usar para satisfacer la necesidad de la funcion Sumar.
	myStub := &stubLogger{}

	// Se ejecuta el test.
	res := SumarDummy(n1, n2, myStub)

	assert.Equal(t, promisedValue, res, "Deben ser iguales")
}

func TestSumarSpy(t *testing.T) {
	n1 := 3
	n2 := 5
	promisedValue := 8

	// Se genera el objeto spy a usar.
	mySpy := &spyLogger{}

	// Se ejecuta el test y se validan el resultado y que spyCalled sea true para dar el test por válido, todo con los assert.
	res := SumarDummy(n1, n2, mySpy)
	assert.Equal(t, promisedValue, res, "Deben ser iguales.")
	assert.True(t, mySpy.spyCalled)
}

func TestSumarRestricted(t *testing.T) {
	n1 := 3
	n2 := 5
	cliente := "John Doe"
	promisedValue := 8

	// Se genera el objeto mock a usar para satisfacer la necesidad de la funcion Sumar.
	myMock := &mockConfig{}

	// Se ejecuta el test y se valida el resultado y que el mock haya registrado la informacion correcta.
	res := SumaRestricted(n1, n2, myMock, cliente)
	assert.Equal(t, promisedValue, res, "Deben ser iguales.")
	assert.Equal(t, cliente, myMock.clienteUsado)
}

func TestSumaRestrictedFake(t *testing.T) {
	n1 := 3
	n2 := 5
	cliente := "John Doe"
	cliente_bis := "Mister Pmosh"
	promisedValue := 8
	promisedValueError := -99999

	// Se genera el objeto fake a usar.
	myFake := &fakeConfig{}

	// Se ejecuta el test y se valida que para el cliente autorizado devuelva el resultado correcto de la suma y que para el otro cliente devuelve el error.
	res_ok := SumaRestricted(n1, n2, myFake, cliente)
	assert.Equal(t, promisedValue, res_ok, "Deben ser iguales.")

	res_error := SumaRestricted(n1, n2, myFake, cliente_bis)
	assert.Equal(t, promisedValueError, res_error, "Deben ser iguales")
}

func TestRestar(t *testing.T) {
	n1 := 10
	n2 := 6
	promisedValue := 4

	res := Restar(n1, n2)

	/*if res != promisedValue {
		t.Errorf("Funcion Resta() arrojo el resultado = %v, pero el esperado es %v", res, promisedValue)
	}*/
	assert.Equal(t, promisedValue, res, fmt.Sprintf("Funcion Resta() arrojo el resultado %v, pero el esperado es %v", res, promisedValue))
}

func TestMultiplicar(t *testing.T) {
	n1 := 8
	n2 := 5
	promisedValue := 40

	res := Multiplicar(n1, n2)

	/*if res != promisedValue {
		t.Errorf("Funcion Multiplicar() arrojo el resultado = %v, pero el esperado es %v", res, promisedValue)
	}*/
	assert.Equal(t, promisedValue, res, fmt.Sprintf("Funcion Multiplicar() arrojo el resultado %v, pero el esperado es %v", res, promisedValue))
}

func TestDividir(t *testing.T) {
	n1 := 5
	n2 := 2
	promisedValue := 2

	res, err := Dividir(n1, n2)

	/*if res != promisedValue {
		t.Errorf("Funcion Dividir() arrojo el resultado = %v, pero el esperado es %v", res, promisedValue)
	}*/
	if err != nil {
		assert.Fail(t, err.Error())
	} else {
		assert.Equal(t, promisedValue, res, fmt.Sprintf("Funcion Dividir() arrojo el resultado %v, pero el esperado es %v", res, promisedValue))
	}
}

func TestOrdenarAsc(t *testing.T) {
	nums := []int{9, 12, 3, 2, 6, 5}
	promised := []int{2, 3, 5, 6, 9, 12}

	res := OrdenarAsc(nums)

	/*if res != promisedValue {
		t.Errorf("Funcion Dividir() arrojo el resultado = %v, pero el esperado es %v", res, promisedValue)
	}*/
	assert.Equal(t, promised, res, fmt.Sprintf("Funcion OrdenarAsc() arrojo el resultado %v, pero el esperado es %v", res, promised))
}

func TestOrdenarDesc(t *testing.T) {
	nums := []int{9, 12, 3, 2, 6, 5}
	promised := []int{12, 9, 6, 5, 3, 2}

	res := OrdenarDesc(nums)

	/*if res != promisedValue {
		t.Errorf("Funcion Dividir() arrojo el resultado = %v, pero el esperado es %v", res, promisedValue)
	}*/
	assert.Equal(t, promised, res, fmt.Sprintf("Funcion OrdenarDesc() arrojo el resultado %v, pero el esperado es %v", res, promised))
}
