package internal

import (
	"encoding/json"
	"errors"
	"fmt"

	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/18_storage2/TurnoManana/internal/transaccion/models"
	"github.com/stretchr/testify/assert"
)

type mockStore struct {
	transactionBeforeUpdate Transaction
}

func (m *mockStore) Read(tran Transaction) bool {
	m.transactionBeforeUpdate = tran
	return true
}

var Datos string = `[{
	"id": 2,
	"codigo": "24safdsadfasdf385",
	"moneda": "Peso Colombiano",
	"monto": "$8228845654645678",
	"emisor": "Luis",
	"receptor": "Perez",
	"fecha": "01/01/2001"
   },
   {
	"id": 3,
	"codigo": "11673-417",
	"moneda": "Franc",
	"monto": "$2.76",
	"emisor": "minstone2",
	"receptor": "sinnott2",
	"fecha": "1/4/2021"
   }]`

type StubStore struct {
	useMethodRead bool
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}
func (s *StubStore) Read(data interface{}) error {
	s.useMethodRead = true
	return json.Unmarshal([]byte(Datos), &data)
}

type StubStoreError struct {
	useMethodRead bool
}

func (s *StubStoreError) Write(data interface{}) error {
	return errors.New("Error al cargar transaccion")
}
func (s *StubStoreError) Read(data interface{}) error {
	s.useMethodRead = true
	return errors.New("No hay un archivo con trasnacciones")
}

func TestGetAll(t *testing.T) {
	stubStore := &StubStore{}
	repo := NewRepository(stubStore)

	var excepted []Transaction

	err := json.Unmarshal([]byte(Datos), &excepted)

	assert.Nil(t, err)

	tran, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, excepted, tran)
}

func TestGetAllError(t *testing.T) {
	stubStored := &StubStoreError{}
	repod := NewRepository(stubStored)

	tran, err := repod.GetAll()

	assert.Nil(t, tran)
	assert.True(t, stubStored.useMethodRead)
	assert.Error(t, err)
}

func TestUpdateCodigo(t *testing.T) {

	stubStore := &StubStore{false}
	repo := NewRepository(stubStore)
	tran2, _ := repo.GetTransactionById(2)
	codgUpdate := "AfterUpdatecod-123"

	tranUpdate, err := repo.Update(2, codgUpdate, "Peso", "55.6", "pepe", "luis", "13/12/2021")

	assert.True(t, stubStore.useMethodRead)
	assert.Equal(t, tran2.ID, tranUpdate.ID)
	assert.Equal(t, codgUpdate, codgUpdate)
	assert.Nil(t, err)
}

func TestUpdateCodigoError(t *testing.T) {
	stubStore := &StubStoreError{false}
	repo := NewRepository(stubStore)
	codgUpdate := "AfterUpdatecod-123"

	tranUpdate, err := repo.Update(2, codgUpdate, "Peso", "55.6", "pepe", "luis", "13/12/2021")

	assert.True(t, stubStore.useMethodRead)
	assert.Equal(t, Transaction{}, tranUpdate)
	assert.Error(t, err)
}

func TestUpdateCodigoAndMonto(t *testing.T) {
	stubStore := &StubStore{false}
	repo := NewRepository(stubStore)
	codgUpdate := "AfterUpdatecod-123"
	monto := "88.5"
	transactionTest := Transaction{2, codgUpdate, "Peso Colombiano", monto, "Luis", "Perez", "01/01/2001"}
	mock := &mockStore{transactionTest}
	isRead := mock.Read(transactionTest)

	tranUpdate, err := repo.UpdateCodigoAndMonto(2, codgUpdate, monto)

	assert.True(t, stubStore.useMethodRead)
	assert.Nil(t, err)
	assert.Equal(t, tranUpdate, mock.transactionBeforeUpdate)
	assert.True(t, isRead)
}

func TestInsert(t *testing.T) {

	transaction := models.Transaction{
		Codigo:   "24safdsadfasdf385",
		Moneda:   "Peso Colombiano",
		Monto:    "$8228845654645678",
		Emisor:   "Luis",
		Receptor: "Perez",
		Fecha:    "01/01/2001",
	}

	repo := NewRepositorySQL()
	tranUpdate, err := repo.Store(transaction)

	assert.Equal(t, transaction.Codigo, tranUpdate.Codigo)
	assert.Equal(t, transaction.Moneda, tranUpdate.Moneda)
	assert.Equal(t, transaction.Moneda, tranUpdate.Moneda)
	assert.Equal(t, transaction.Emisor, tranUpdate.Emisor)
	assert.Equal(t, transaction.Emisor, tranUpdate.Emisor)
	assert.Equal(t, transaction.Receptor, tranUpdate.Receptor)
	assert.Nil(t, err)
}

func TestGetById(t *testing.T) {

	transaction := models.Transaction{
		Codigo:   "24safdsadfasdf385",
		Moneda:   "Peso Colombiano",
		Monto:    "$8228845654645678",
		Emisor:   "Luis",
		Receptor: "Perez",
		Fecha:    "01/01/2001",
	}

	repo := NewRepositorySQL()
	tranUpdate, err := repo.GetById(1)

	assert.Equal(t, transaction.Codigo, tranUpdate.Codigo)
	assert.Equal(t, transaction.Moneda, tranUpdate.Moneda)
	assert.Equal(t, transaction.Moneda, tranUpdate.Moneda)
	assert.Equal(t, transaction.Emisor, tranUpdate.Emisor)
	assert.Equal(t, transaction.Emisor, tranUpdate.Emisor)
	assert.Equal(t, transaction.Receptor, tranUpdate.Receptor)
	assert.Nil(t, err)
}

func TestGetAllSql(t *testing.T) {

	repo := NewRepositorySQL()
	transUpdate, err := repo.GetAll()
	fmt.Println(transUpdate)
	assert.NotNil(t, transUpdate)
	assert.True(t, len(transUpdate) >= 0)
	assert.Nil(t, err)
}
func TestDeletelSql(t *testing.T) {

	repo := NewRepositorySQL()

	transaction := models.Transaction{
		Codigo:   "24safdsadfasdf385",
		Moneda:   "Peso Colombiano",
		Monto:    "$8228845654645678",
		Emisor:   "Luis",
		Receptor: "Perez",
		Fecha:    "01/01/2001",
	}
	tranUpdate, err := repo.Store(transaction)

	assert.Nil(t, err)
	assert.NotNil(t, tranUpdate)
	assert.Equal(t, transaction.Codigo, tranUpdate.Codigo)

	err = repo.Delete(tranUpdate.ID)
	assert.Nil(t, err)
}
