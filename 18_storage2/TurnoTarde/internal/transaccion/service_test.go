package internal

import (
	"encoding/json"
	"errors"

	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/18_storage2/TurnoTarde/pkg/store"
	"github.com/stretchr/testify/assert"
)

var transactions string = `[{
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

func TestUpdate(t *testing.T) {
	transacionTest := Transaction{
		Codigo:   "New001",
		Moneda:   "ARS",
		Monto:    "850",
		Emisor:   "Luis",
		Receptor: "Ppepe",
		Fecha:    "13/12/2021",
	}

	mock := store.Mock{Data: []byte(transactions), IsStoreRead: false}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService := NewService(myRepo)

	newTransaction, err := myService.Update(2, transacionTest.Codigo, transacionTest.Moneda, transacionTest.Monto,
		transacionTest.Emisor, transacionTest.Receptor, transacionTest.Fecha)

	assert.Nil(t, err)
	assert.True(t, mock.IsStoreRead)
	assert.True(t, mock.IsStoreWrite)
	assert.Equal(t, transacionTest.Codigo, newTransaction.Codigo)
	assert.Equal(t, transacionTest.Emisor, newTransaction.Emisor)
	assert.Equal(t, transacionTest.Fecha, newTransaction.Fecha)
	assert.Equal(t, transacionTest.Moneda, newTransaction.Moneda)
	assert.Equal(t, transacionTest.Monto, newTransaction.Monto)
	assert.Equal(t, transacionTest.Receptor, newTransaction.Receptor)
	assert.Equal(t, 2, newTransaction.ID)
}

func TestUpdateError(t *testing.T) {
	transacionTest := Transaction{
		Codigo:   "New001",
		Moneda:   "ARS",
		Monto:    "850",
		Emisor:   "Luis",
		Receptor: "Ppepe",
		Fecha:    "13/12/2021",
	}
	erCreated := errors.New("Error al updtear transaction")
	mock := store.Mock{Data: []byte(transactions), Err: erCreated}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService := NewService(myRepo)

	newTransaction, err := myService.Update(2, transacionTest.Codigo, transacionTest.Moneda, transacionTest.Monto,
		transacionTest.Emisor, transacionTest.Receptor, transacionTest.Fecha)

	assert.True(t, !mock.IsStoreRead)
	assert.True(t, !mock.IsStoreWrite)
	assert.NotNil(t, err, erCreated)
	assert.Equal(t, Transaction{}, newTransaction)

}

func TestDelete(t *testing.T) {
	mock := store.Mock{Data: []byte(transactions)}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService := NewService(myRepo)

	err := myService.Delete(2)
	tranAfterDelete, _ := myService.GetTransactionById(2)

	assert.Nil(t, err)
	assert.Equal(t, Transaction{}, tranAfterDelete)

}

func TestDeleteError(t *testing.T) {
	mock := store.Mock{Data: []byte(transactions)}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService := NewService(myRepo)

	err := myService.Delete(100)

	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestStore(t *testing.T) {
	transacionTest := Transaction{
		Codigo:   "New001",
		Moneda:   "ARS",
		Monto:    "850",
		Emisor:   "Luis",
		Receptor: "Ppepe",
		Fecha:    "13/12/2021",
	}

	mock := store.Mock{Data: []byte(transactions), IsStoreRead: false}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService := NewService(myRepo)

	newTransaction, err := myService.Store(transacionTest.Codigo, transacionTest.Moneda, transacionTest.Monto,
		transacionTest.Emisor, transacionTest.Receptor, transacionTest.Fecha)

	assert.Nil(t, err)
	assert.True(t, mock.IsStoreRead)
	assert.True(t, mock.IsStoreWrite)
	assert.Equal(t, transacionTest.Codigo, newTransaction.Codigo)
	assert.Equal(t, transacionTest.Emisor, newTransaction.Emisor)
	assert.Equal(t, transacionTest.Fecha, newTransaction.Fecha)
	assert.Equal(t, transacionTest.Moneda, newTransaction.Moneda)
	assert.Equal(t, transacionTest.Monto, newTransaction.Monto)
	assert.Equal(t, transacionTest.Receptor, newTransaction.Receptor)
}

func TestStoreError(t *testing.T) {
	transacionTest := Transaction{
		Codigo:   "New001",
		Moneda:   "ARS",
		Monto:    "850",
		Emisor:   "Luis",
		Receptor: "Ppepe",
		Fecha:    "13/12/2021",
	}
	erCreated := errors.New("Error al store transaction")
	mock := store.Mock{Data: []byte(transactions), Err: erCreated,ReadAlways: true}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService := NewService(myRepo)

	newTransaction, err := myService.Store(transacionTest.Codigo, transacionTest.Moneda, transacionTest.Monto,
		transacionTest.Emisor, transacionTest.Receptor, transacionTest.Fecha)
		
	assert.Empty(t,newTransaction)
	assert.Error(t,err)
	assert.True(t, mock.IsStoreRead)
	assert.True(t, !mock.IsStoreWrite)
	assert.NotNil(t, err, erCreated)
	assert.Equal(t, Transaction{}, newTransaction)

}

func TestStoreError2(t *testing.T) {
	transacionTest := Transaction{
		Codigo:   "New001",
		Moneda:   "ARS",
		Monto:    "850",
		Emisor:   "Luis",
		Receptor: "Ppepe",
		Fecha:    "13/12/2021",
	}
	erCreated := errors.New("Error al store transaction")
	mock := store.Mock{Data: []byte(transactions), Err: erCreated}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService := NewService(myRepo)

	newTransaction, err := myService.Store(transacionTest.Codigo, transacionTest.Moneda, transacionTest.Monto,
		transacionTest.Emisor, transacionTest.Receptor, transacionTest.Fecha)
		
	assert.Empty(t,newTransaction)
	assert.Error(t,err)
	assert.True(t, !mock.IsStoreRead)
	assert.True(t, !mock.IsStoreWrite)
	assert.NotNil(t, err, erCreated)
	assert.Equal(t, Transaction{}, newTransaction)

}

func TestGetAll1(t *testing.T) {
	
	mock := store.Mock{Data: []byte(transactions)}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService := NewService(myRepo)

	var excepted []Transaction

	err := json.Unmarshal([]byte(transactions), &excepted)
	assert.Nil(t, err)
	
	newTransactions, err := myService.GetAll()

	assert.Equal(t,len(excepted),len(newTransactions))
	assert.True(t, mock.IsStoreRead)
	assert.Nil(t, err)
}

func TestGetAll1Error(t *testing.T) {

	erCreated := errors.New("Error al obtener transaction")
	mock := store.Mock{Data: []byte(transactions), Err: erCreated}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService := NewService(myRepo)

	_, err := myService.GetAll()

	assert.Error(t,err)
	assert.Equal(t, err, erCreated)	

}

func TestGetOneError(t *testing.T) {

	erCreated := errors.New("Error al obtener transaction")
	mock := store.Mock{Data: []byte(transactions), Err: erCreated}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService := NewService(myRepo)

	_, err := myService.GetTransactionById(1)

	assert.Error(t,err)
	assert.Equal(t, err, erCreated)	

}


func TestGetOne(t *testing.T) {
	
	mock := store.Mock{Data: []byte(transactions)}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService := NewService(myRepo)

	var excepted []Transaction

	err := json.Unmarshal([]byte(transactions), &excepted)
	assert.Nil(t, err)
	
	newTransaction, err := myService.GetTransactionById(2)

	assert.Equal(t,excepted[0],newTransaction)
	assert.True(t, mock.IsStoreRead)
	assert.Nil(t, err)
}


func TestUpdateCodigoAndMonto1(t *testing.T) {
	codigo := "Cod123"
	monto := "8500"
	mock := store.Mock{Data: []byte(transactions)}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService := NewService(myRepo)


	newTransaction, err := myService.UpdateCodigoAndMonto(2,codigo,monto)
	assert.Nil(t, err)

	exceptedTransaction,err := myService.GetTransactionById(2)

	assert.Nil(t, err)
	assert.Equal(t,exceptedTransaction,newTransaction)
	assert.True(t, mock.IsStoreRead)
	assert.True(t, mock.IsStoreWrite)
}

func TestUpdateCodigoAndMonto1Error(t *testing.T) {
	codigo := "Cod123"
	monto := "8500"
	erCreated := errors.New("Error al updtear transaction")
	mock := store.Mock{Data: []byte(transactions), Err: erCreated}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService := NewService(myRepo)


	newTransaction, err := myService.UpdateCodigoAndMonto(2,codigo,monto)
	assert.NotNil(t, err)
	assert.Error(t,err)
	assert.Empty(t,newTransaction)
	
}