package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/18_storage2/TurnoTarde/db"
	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/18_storage2/TurnoTarde/internal/transaccion/models"
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

func GetColumns() []string {
	return []string{"ID", "Codigo", "Moneda", "Monto", "Emisor", "Receptor", "Fecha"}
}

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
	db := db.StorageDB
	transaction := models.Transaction{
		Codigo:   "24safdsadfasdf385",
		Moneda:   "Peso Colombiano",
		Monto:    "$8228845654645678",
		Emisor:   "Luis",
		Receptor: "Perez",
		Fecha:    "01/01/2001",
	}

	repo := NewRepositorySQL(db)
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
	db := db.StorageDB
	repo := NewRepositorySQL(db)
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
	db := db.StorageDB
	repo := NewRepositorySQL(db)
	transUpdate, err := repo.GetAll()
	fmt.Println(transUpdate)
	assert.NotNil(t, transUpdate)
	assert.True(t, len(transUpdate) >= 0)
	assert.Nil(t, err)
}
func TestDeletelSql(t *testing.T) {

	db := db.StorageDB
	repo := NewRepositorySQL(db)

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

func TestUpdateSql(t *testing.T) {

	transaction := models.Transaction{
		ID:       2,
		Codigo:   "878885",
		Moneda:   "Peso ARGENTINO",
		Monto:    "$55",
		Emisor:   "Luis",
		Receptor: "jose",
		Fecha:    "01/01/2001",
	}

	db := db.StorageDB
	repo := NewRepositorySQL(db)
	tranUpdate, err := repo.Update(transaction, context.Background())

	assert.Equal(t, transaction.Codigo, tranUpdate.Codigo)
	assert.Equal(t, transaction.Moneda, tranUpdate.Moneda)
	assert.Equal(t, transaction.Moneda, tranUpdate.Moneda)
	assert.Equal(t, transaction.Emisor, tranUpdate.Emisor)
	assert.Equal(t, transaction.Emisor, tranUpdate.Emisor)
	assert.Equal(t, transaction.Receptor, tranUpdate.Receptor)
	assert.Nil(t, err)
}

func TestInsertMock(t *testing.T) {

	transaction := models.Transaction{
		Codigo:   "24safdsadfasdf385",
		Moneda:   "Peso Colombiano",
		Monto:    "$8228845654645678",
		Emisor:   "Luis",
		Receptor: "Perez",
		Fecha:    "01/01/2001",
	}

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(InsertOneTest)
	mock.ExpectExec(InsertOneTest).WillReturnResult(sqlmock.NewResult(5, 1))

	repo := NewRepositorySQL(db)
	tranUpdate, err := repo.Store(transaction)

	// rows := sqlmock.NewRows(GetColumns())
	// rows.AddRow(1, "24safdsadfasdf385", "Peso Colombiano", "$8228845654645678", "Luis",
	// 	"Perez", "01/01/2001")

	assert.Equal(t, 5, tranUpdate.ID)
	assert.Equal(t, transaction.Codigo, tranUpdate.Codigo)
	assert.Equal(t, transaction.Moneda, tranUpdate.Moneda)
	assert.Equal(t, transaction.Emisor, tranUpdate.Emisor)
	assert.Equal(t, transaction.Receptor, tranUpdate.Receptor)
	assert.Nil(t, err)
}

func TestGetOneMock(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	columns := GetColumns()
	rows := sqlmock.NewRows(columns)
	rows.AddRow(1, "24safdsadfasdf385", "Peso Colombiano", "$8228845654645678", "Luis",
		"Perez", "01/01/2001")

	mock.ExpectQuery(GetByIdTest).WithArgs(1).WillReturnRows(rows)

	repo := NewRepositorySQL(db)
	tranRecived, err := repo.GetById(1)

	assert.Equal(t, 1, tranRecived.ID)
	assert.Equal(t, "24safdsadfasdf385", tranRecived.Codigo)
	assert.Nil(t, err)
}

func TestUpdateMock(t *testing.T) {
	transactionUpdate := models.Transaction{
		ID:       5,
		Codigo:   "ca455",
		Moneda:   "PesoColombianosadfsad",
		Monto:    "$8",
		Emisor:   "L",
		Receptor: "P",
		Fecha:    "01/01/2001",
	}


	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	 columns := GetColumns()
	 rows := sqlmock.NewRows(columns)

	rows.AddRow(5, "24safdsadfasdf385", "Peso Colombiano", "$8228845654645678", "Luis",
	 	"Perez", "01/01/2001")

	 mock.ExpectQuery(GetByIdTest).WithArgs(5).WillReturnRows(rows)
	
	mock.ExpectPrepare(UpdateTest)
	mock.ExpectExec(UpdateTest).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewRepositorySQL(db)

	//compruebo que existe tran
	tranGet,err := repo.GetById(5)

	assert.NoError(t, err)
	assert.Equal(t,5,tranGet.ID)

	// acutalizo tran
	tranUpdate, err := repo.Update(transactionUpdate, context.Background())

	assert.Equal(t, 5, tranUpdate.ID)
	assert.Equal(t, transactionUpdate.Codigo, tranUpdate.Codigo)
	assert.Equal(t, transactionUpdate.Moneda, tranUpdate.Moneda)
	assert.Equal(t, transactionUpdate.Emisor, tranUpdate.Emisor)
	assert.Equal(t, transactionUpdate.Receptor, tranUpdate.Receptor)
	assert.Nil(t, err)
}

func TestDeleteMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := GetColumns()
	rows := sqlmock.NewRows(columns)

	rows.AddRow(5, "24safdsadfasdf385", "Peso Colombiano", "$8228845654645678", "Luis",
		"Perez", "01/01/2001")

	mock.ExpectQuery(GetByIdTest).WithArgs(5).WillReturnRows(rows)

	prep := mock.ExpectPrepare(DeleteTest)
	prep.ExpectExec().WithArgs(5).WillReturnResult(sqlmock.NewResult(0, 1))
	
	repo := NewRepositorySQL(db)

	//compruebo que existe
	tranGet,err := repo.GetById(5)

	assert.NoError(t, err)
	assert.Equal(t,5,tranGet.ID)

	// elimino tran
	err = repo.Delete(5)
	assert.NoError(t, err)

	// compruebo que no existe
	tranGet,err = repo.GetById(5)
	assert.Error(t, err)
	assert.Empty(t,tranGet)

}

func TestInsertMockTxdb(t *testing.T) {

	transaction := models.Transaction{
		Codigo:   "unosss",
		Moneda:   "Peso Colombiaaaano",
		Monto:    "$8228845654645678",
		Emisor:   "Luis",
		Receptor: "Perez",
		Fecha:    "01/01/2001",
	}
	db, err := db.InitDb()
	assert.NoError(t, err)
	repo := NewRepositorySQL(db)
	tranUpdate, err := repo.Store(transaction)
	assert.Equal(t, transaction.Codigo, tranUpdate.Codigo)
	assert.Equal(t, transaction.Moneda, tranUpdate.Moneda)
	assert.Equal(t, transaction.Emisor, tranUpdate.Emisor)
	assert.Equal(t, transaction.Receptor, tranUpdate.Receptor)
	assert.Nil(t, err)
}


func TestGetOneMockTxdb(t *testing.T) {
	transaction := models.Transaction{
		Codigo:   "unosss",
		Moneda:   "Peso Colombiaaaano",
		Monto:    "$8228845654645678",
		Emisor:   "Luis",
		Receptor: "Perez",
		Fecha:    "01/01/2001",
	}
	db, err := db.InitDb()
	assert.NoError(t, err)
	repo := NewRepositorySQL(db)
	tranUpdate, err := repo.Store(transaction)
	assert.Nil(t, err)
	tranGet, err := repo.GetById(tranUpdate.ID)

	assert.Equal(t,tranUpdate.ID,tranGet.ID)
	assert.Equal(t,tranUpdate.Codigo,tranGet.Codigo)
	assert.Equal(t,  tranUpdate.Moneda,tranGet.Moneda)
	assert.Equal(t,  tranUpdate.Emisor,tranGet.Emisor)
	assert.Equal(t, tranUpdate.Receptor,tranGet.Receptor)
	assert.Nil(t, err)
}



func TestUpdateMockTxdb(t *testing.T) {

	transaction := models.Transaction{
		Codigo:   "unosss",
		Moneda:   "Peso Colombiaaaano",
		Monto:    "$8228845654645678",
		Emisor:   "Luis",
		Receptor: "Perez",
		Fecha:    "01/01/2001",
	}

	db, err := db.InitDb()
	assert.NoError(t, err)
	repo := NewRepositorySQL(db)

	// inserto tran
	tranUpdate, err := repo.Store(transaction)
	assert.Nil(t, err)

	// compruebo que existe
	tranGet, err := repo.GetById(tranUpdate.ID)
	assert.Nil(t, err)
	// acutalizo codigo
	tranGet.Codigo = "nuevoCodigoPrueba"
	tranUpdate, err = repo.Update(tranGet,context.Background())

	assert.Equal(t, tranGet.ID, tranUpdate.ID)
	assert.Equal(t, "nuevoCodigoPrueba", tranUpdate.Codigo)
	assert.Nil(t, err)
}

func TestDeleteMockTxdb(t *testing.T) {
	transaction := models.Transaction{
		Codigo:   "unosss",
		Moneda:   "Peso Colombiaaaano",
		Monto:    "$8228845654645678",
		Emisor:   "Luis",
		Receptor: "Perez",
		Fecha:    "01/01/2001",
	}

	db, err := db.InitDb()
	assert.NoError(t, err)
	repo := NewRepositorySQL(db)
	
	// inserto tran
	tranInsert, err := repo.Store(transaction)
	assert.Nil(t, err)
	
	//compruebo que existe
	tranGet,err := repo.GetById(tranInsert.ID)

	assert.NoError(t, err)
	assert.Equal(t,tranInsert.ID,tranGet.ID)

	// elimino tran
	err = repo.Delete(tranInsert.ID)
	assert.NoError(t, err)

	// compruebo que no existe
	tranGet,err = repo.GetById(tranInsert.ID)
	assert.Empty(t,tranGet)

}