package internal

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/17_storage1/PracticaTT/internal/models"
	"github.com/stretchr/testify/assert"
)

var transac string = `[
	{
	"id": 1,
	"cod_transaccion": "123asd456",
	"moneda": "dolar",
	"monto": 20.55,
	"emisor": "Facundo",
	"receptor": "Ezequiel",
	"fecha_trans": "21/01/2021"
   },
   {
	"id": 2,
	"cod_transaccion": "BeforeUpdate",
	"moneda": "dolar",
	"monto": 999.999,
	"emisor": "Facundo",
	"receptor": "Ezequiel",
	"fecha_trans": "21/01/2021"
   },
   {
	"id": 3,
	"cod_transaccion": "123asd456",
	"moneda": "dolar",
	"monto": 20.55,
	"emisor": "Facundo",
	"receptor": "Ezequiel",
	"fecha_trans": "21/01/2021"
   }
   ]
`

type StubStore struct{}

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(transac), &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

type MockStore struct {
	trans Transaccion
}

func (m *MockStore) Read(data interface{}) error {
	var transacciones []Transaccion
	err := json.Unmarshal([]byte(transac), &transacciones)
	if err != nil {
		return err
	}
	m.trans = transacciones[1]

	return json.Unmarshal([]byte(transac), &data)
}

func (m *MockStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	store := StubStore{}
	repo := NewRepository(&store)
	var transacExpected []Transaccion

	json.Unmarshal([]byte(transac), &transacExpected)

	myTransac, err := repo.getAll()

	assert.Equal(t, transacExpected, myTransac, "Los datos no son iguales")
	assert.Nil(t, err)

}

func TestUpdateCodigoYMonto(t *testing.T) {
	//Arrange
	store := MockStore{}
	repo := NewRepository(&store)
	newCodigo := "123asd456"
	newMonto := 20.55

	//Act
	transacResultado, _ := repo.UpdateCodigoYMonto(2, newCodigo, newMonto)

	//Assert
	assert.NotEmpty(t, store.trans)

	assert.Equal(t, "BeforeUpdate", store.trans.CodTransaccion)
	assert.Equal(t, newCodigo, transacResultado.CodTransaccion)
	assert.Equal(t, newMonto, transacResultado.Monto)

}

func TestGetAllRepoSQL(t *testing.T) {
	//Arrange
	repo := NewRepositorySQL()

	//Act
	transacResultado, err := repo.GetAll()

	//Assert
	assert.Nil(t, err)
	assert.True(t, len(transacResultado) > 0)
}

func TestUpdateWithContextSQL(t *testing.T) {
	//Arrange
	repo := NewRepositorySQL()

	transaccionNueva := models.Transaccion{
		CodTransaccion: "7",
		Moneda:         "Dolar",
		Emisor:         "Pepe",
		Receptor:       "Marcos",
		Monto:          7623.33,
	}

	//Act
	transacActualizada, err := repo.UpdateWithContext(context.Background(), transaccionNueva)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, transaccionNueva.Moneda, transacActualizada.Moneda)
	assert.Equal(t, transaccionNueva.Monto, transacActualizada.Monto)
	assert.Equal(t, transaccionNueva.Emisor, transacActualizada.Emisor)
}
