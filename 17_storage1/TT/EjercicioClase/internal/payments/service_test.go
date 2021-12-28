package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/12_testing3/TM/Ejercicios/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/17_storage1/TT/EjercicioClase/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/17_storage1/TT/EjercicioClase/pkg/db"
	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
	useGetAll bool
}

var pays_bis string = `[
	{
		"id": 1,
		"codigo": "0000001",
		"moneda": "ARS",
		"monto": 10535.26,
		"emisor": "Rodrigo Vega Gimenez",
		"receptor": "Juan Pablo Nieto",
		"fecha": "2021-12-01"
	   },
	   {
		"id": 2,
		"codigo": "0000002",
		"moneda": "ARS",
		"monto": 9563.45,
		"emisor": "Rodrigo Vega Gimenez",
		"receptor": "Maximiliano Caceres",
		"fecha": "2021-12-01"
	   }]`

var pays_bis_sql string = `[
	{
		"id": 11,
		"codigo": "AAA001",
		"moneda": "ARS",
		"monto": 895.45,
		"emisor": "Rodrigo Vega",
		"receptor": "Cristian Lopez",
		"fecha": "2021-12-28"
	   }]`

var pays_bis_sql_full_data string = `[
		{
			"id": 12,
			"codigo": "AAA001",
			"moneda": "ARS",
			"monto": 956.56,
			"emisor": "Rodrigo Vega",
			"receptor": "Cristian Lopez",
			"fecha": "2021-12-17"
			"box_closing": {
				"id": 1,
				"responsable": "Federico Gutierrez",
				"fecha": "2021-12-27"
			}
		   }]`

func (s *StubRepository) GetAll() ([]Payment, error) {
	var out []Payment
	err := json.Unmarshal([]byte(pays), &out)
	s.useGetAll = true
	return out, err
}

func (s *StubRepository) Filtrar(values ...string) ([]Payment, error) {
	return []Payment{}, nil
}

func (s *StubRepository) Store(codigo, moneda, emisor, receptor, fecha string, monto float64) (Payment, error) {
	return Payment{}, nil
}

func (s *StubRepository) Update(id int, codigo, moneda, emisor, receptor, fecha string, monto float64) (Payment, error) {
	return Payment{}, nil
}

func (s *StubRepository) UpdateCodigo(id int, codigo string) (Payment, error) {
	return Payment{}, nil
}

func (s *StubRepository) UpdateMonto(id int, monto float64) (Payment, error) {
	return Payment{}, nil
}

func (s *StubRepository) Delete(id int) (string, error) {
	return "", nil
}

func (s *StubRepository) LastId() (int, error) {
	return 0, nil
}

// Todos los test se van a hacer con el MOCK, que ya esta definido en File para que siempre sea validado!
func TestGetAllServiceMock(t *testing.T) {
	//Arrange
	dataByte := []byte(pays_bis)
	var expectedPayments []Payment
	json.Unmarshal(dataByte, &expectedPayments)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	myPayments, _ := service.GetAll()

	assert.Equal(t, expectedPayments, myPayments)
}

func TestGetAllServiceMockError(t *testing.T) {
	//Arrange
	expectedError := errors.New("No hay datos en el Mock")

	dbMock := store.Mock{Err: expectedError}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	myPayments, receivedError := service.GetAll()

	assert.Equal(t, expectedError, receivedError)
	assert.Nil(t, myPayments)
}

func TestStoreServiceMock(t *testing.T) {
	//Arrange
	newPayment := Payment{Codigo: "0000003", Moneda: "ARS", Monto: 674.34, Emisor: "Rodrigo Vega Gimenez", Receptor: "Carlos Miño", Fecha: "2021-12-02"}

	dbMock := store.Mock{Data: []byte(`[]`)}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	createdPayment, _ := service.Store(newPayment.Codigo, newPayment.Moneda, newPayment.Emisor, newPayment.Receptor, newPayment.Fecha, newPayment.Monto)

	assert.Equal(t, newPayment.Codigo, createdPayment.Codigo)
	assert.Equal(t, newPayment.Emisor, createdPayment.Emisor)
	assert.Equal(t, newPayment.Receptor, createdPayment.Receptor)
}

func TestStoreServiceMockError(t *testing.T) {
	//Arrange
	newPayment := Payment{Codigo: "0000003", Moneda: "ARS", Monto: 674.34, Emisor: "Rodrigo Vega Gimenez", Receptor: "Carlos Miño", Fecha: "2021-12-02"}

	expectedError := errors.New("No hay datos en el Mock")

	dbMock := store.Mock{Data: []byte(`[]`), Err: expectedError}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	createdPayment, err := service.Store(newPayment.Codigo, newPayment.Moneda, newPayment.Emisor, newPayment.Receptor, newPayment.Fecha, newPayment.Monto)

	assert.Equal(t, expectedError, err)
	assert.Equal(t, Payment{}, createdPayment)
}

func TestUpdateServiceMock(t *testing.T) {
	//Arrange
	newPayment := Payment{Codigo: "0000003", Moneda: "ARS", Monto: 674.34, Emisor: "Rodrigo Vega Gimenez", Receptor: "Carlos Miño", Fecha: "2021-12-02"}

	dataByte := []byte(pays_bis)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	updatedPayment, _ := service.Update(1, newPayment.Codigo, newPayment.Moneda, newPayment.Emisor, newPayment.Receptor, newPayment.Fecha, newPayment.Monto)

	assert.Equal(t, newPayment.Codigo, updatedPayment.Codigo)
	assert.Equal(t, newPayment.Emisor, updatedPayment.Emisor)
	assert.Equal(t, newPayment.Receptor, updatedPayment.Receptor)
	assert.Equal(t, 1, updatedPayment.Id)
}

func TestUpdateNotFoundMock(t *testing.T) {
	//Arrange
	newPayment := Payment{Codigo: "0000003", Moneda: "ARS", Monto: 674.34, Emisor: "Rodrigo Vega Gimenez", Receptor: "Carlos Miño", Fecha: "2021-12-02"}

	dataByte := []byte(pays_bis)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.Update(3, newPayment.Codigo, newPayment.Moneda, newPayment.Emisor, newPayment.Receptor, newPayment.Fecha, newPayment.Monto)
	assert.Error(t, err)
}

// Por si hay un error en la conexion.
func TestUpdateErrorMock(t *testing.T) {
	//Arrange
	newPayment := Payment{Codigo: "0000003", Moneda: "ARS", Monto: 674.34, Emisor: "Rodrigo Vega Gimenez", Receptor: "Carlos Miño", Fecha: "2021-12-02"}

	dbMock := store.Mock{Err: errors.New("Error en la conexión con la base de datos.")}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.Update(1, newPayment.Codigo, newPayment.Moneda, newPayment.Emisor, newPayment.Receptor, newPayment.Fecha, newPayment.Monto)
	assert.NotNil(t, err)
}

func TestUpdateCodigoServiceMock(t *testing.T) {
	//Arrange
	newCode := "0000004"

	dataByte := []byte(pays_bis)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	updatedPayment, _ := service.UpdateCodigo(2, newCode)

	assert.Equal(t, newCode, updatedPayment.Codigo)
	assert.Equal(t, 2, updatedPayment.Id)
}

func TestUpdateCodigoNotFoundMock(t *testing.T) {
	//Arrange
	newCode := "0000004"

	dataByte := []byte(pays_bis)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.UpdateCodigo(2, newCode)
	assert.Error(t, err)
}

// Por si hay un error en la conexion.
func TestUpdateCodigoErrorMock(t *testing.T) {
	//Arrange
	newCode := "0000004"

	dbMock := store.Mock{Err: errors.New("Error en la conexión con la base de datos.")}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.UpdateCodigo(2, newCode)
	assert.NotNil(t, err)
}

func TestUpdateMontoServiceMock(t *testing.T) {
	//Arrange
	newAmount := 1050.55

	dataByte := []byte(pays_bis)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	updatedPayment, _ := service.UpdateMonto(2, newAmount)

	assert.Equal(t, newAmount, updatedPayment.Monto)
	assert.Equal(t, 2, updatedPayment.Id)
}

func TestUpdateMontoNotFoundMock(t *testing.T) {
	//Arrange
	newAmount := 1050.55

	dataByte := []byte(pays_bis)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.UpdateMonto(2, newAmount)
	assert.Error(t, err)
}

// Por si hay un error en la conexion.
func TestUpdateMontoErrorMock(t *testing.T) {
	//Arrange
	newAmount := 1050.55

	dbMock := store.Mock{Err: errors.New("Error en la conexión con la base de datos.")}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.UpdateMonto(2, newAmount)
	assert.NotNil(t, err)
}

func TestDeleteServiceMock(t *testing.T) {
	//Arrange

	dataByte := []byte(pays_bis)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.Delete(2)

	assert.Nil(t, err)

	allPayments, _ := service.GetAll()

	assert.Equal(t, 1, len(allPayments))
}

// ACA FALTAN CONTROL SOBRE DELETE Y UPDATE NOT FOUND Y ERROR.

func TestDeleteNotFoundServiceMock(t *testing.T) {
	//Arrange

	dataByte := []byte(pays_bis)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.Delete(4)

	assert.Error(t, err)
}

func TestDeleteErrorServiceMock(t *testing.T) {
	//Arrange
	dbMock := store.Mock{Err: errors.New("Error en la conexión con la base de datos.")}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.Delete(4)

	assert.Error(t, err)
}

// AQUI COMIENZAN LOS TESTS DEL SERVICIO CON SQL, ES DECIR SE USARA LA ESTRUCTURA DE SERVICESQL.
func TestStoreServiceSql(t *testing.T) {
	newPayment := models.Payment{
		Codigo:   "AAA001",
		Moneda:   "ARS",
		Monto:    956.56,
		Emisor:   "Rodrigo Vega",
		Receptor: "Cristian Lopez",
		Fecha:    "2021-12-17",
	}

	repo := NewRepositorySql()

	service := NewServiceSql(repo)

	createdPayment, err := service.Store(newPayment.Codigo, newPayment.Moneda, newPayment.Emisor, newPayment.Receptor, newPayment.Fecha, newPayment.Monto)

	assert.Equal(t, newPayment.Codigo, createdPayment.Codigo)
	assert.Equal(t, newPayment.Moneda, createdPayment.Moneda)
	assert.Nil(t, err)
}

func TestGetByIdServiceSql(t *testing.T) {
	newPayment := models.Payment{
		Id:       11,
		Codigo:   "AAA001",
		Moneda:   "R$$",
		Monto:    956.56,
		Emisor:   "Rodrigo Vega",
		Receptor: "Cristian Lopez",
		Fecha:    "2021-12-17",
	}

	repo := NewRepositorySql()

	service := NewServiceSql(repo)

	obtainedPayment := service.GetById(11)

	assert.Equal(t, newPayment.Codigo, obtainedPayment.Codigo)
	assert.Equal(t, newPayment.Moneda, obtainedPayment.Moneda)
}

func TestGetByCodeServiceSql(t *testing.T) {
	newPayment := models.Payment{
		Codigo:   "AAA001",
		Moneda:   "ARS",
		Monto:    956.56,
		Emisor:   "Rodrigo Vega",
		Receptor: "Cristian Lopez",
		Fecha:    "2021-12-17",
	}

	repo := NewRepositorySql()

	service := NewServiceSql(repo)

	obtainedPayment := service.GetByCode("AAA001")

	assert.Equal(t, newPayment.Codigo, obtainedPayment.Codigo)
	assert.Equal(t, newPayment.Moneda, obtainedPayment.Moneda)
}

func TestGetAllPaymentsServiceSql(t *testing.T) {
	dataByte := []byte(pays_bis_sql)
	var expectedPayments []models.Payment
	json.Unmarshal(dataByte, &expectedPayments)

	repo := NewRepositorySql()

	service := NewServiceSql(repo)

	obtainedPayments, _ := service.GetAllPayments()

	assert.Equal(t, expectedPayments, obtainedPayments)
	assert.True(t, len(obtainedPayments) > 0)
}

func TestGetAllPaymentsServiceSql_Failed(t *testing.T) {
	dataByte := []byte(pays_bis_sql)
	var expectedPayments []models.Payment
	json.Unmarshal(dataByte, &expectedPayments)

	repo := NewRepositorySql()

	service := NewServiceSql(repo)

	obtainedPayments, _ := service.GetAllPayments()

	assert.NotEqual(t, expectedPayments, obtainedPayments)
}

func TestUpdateServiceSql(t *testing.T) {
	expectedPayment := models.Payment{
		Id:       11,
		Codigo:   "AAA001",
		Moneda:   "ARS",
		Monto:    float64(895.45),
		Emisor:   "Rodrigo Vega",
		Receptor: "Cristian Lopez",
		Fecha:    "2021-12-28",
	}

	repo := NewRepositorySql()

	service := NewServiceSql(repo)

	updatedPayment, err := service.Update(expectedPayment)

	assert.Equal(t, expectedPayment.Moneda, updatedPayment.Moneda)
	assert.Equal(t, expectedPayment.Monto, updatedPayment.Monto)
	assert.Equal(t, expectedPayment.Receptor, updatedPayment.Receptor)
	assert.Nil(t, err)
}

func TestUpdateServiceSql_Failed(t *testing.T) {
	expectedPayment := models.Payment{
		Id:       11,
		Codigo:   "AAA001",
		Moneda:   "R$$",
		Monto:    float64(95.80),
		Emisor:   "Rodrigo Vega",
		Receptor: "Cristiano Lope",
		Fecha:    "2021-12-18",
	}

	repo := NewRepositorySql()

	service := NewServiceSql(repo)

	_, err := service.Update(expectedPayment)

	assert.Equal(t, "No se encontró la transacción.", err.Error())
}

func TestDeleteServiceSql(t *testing.T) {
	newPayment := models.Payment{
		Codigo:   "AAA001",
		Moneda:   "ARS",
		Monto:    956.56,
		Emisor:   "Rodrigo Vega",
		Receptor: "Cristian Lopez",
		Fecha:    "2021-12-17",
	}

	repo := NewRepositorySql()

	service := NewServiceSql(repo)

	_, err := service.Store(newPayment.Codigo, newPayment.Moneda, newPayment.Emisor, newPayment.Receptor, newPayment.Fecha, newPayment.Monto)

	err = service.Delete(newPayment.Id)
	assert.Nil(t, err)
}

func TestDeleteServiceSql_Failed(t *testing.T) {
	repo := NewRepositorySql()

	service := NewServiceSql(repo)

	err := service.Delete(0)
	assert.Equal(t, "No se encontró la transacción.", err.Error())
}

func TestGetByIdWithContextServiceSql(t *testing.T) {
	newPayment := models.Payment{
		Id:       12,
		Codigo:   "AAA001",
		Moneda:   "ARS",
		Monto:    956.56,
		Emisor:   "Rodrigo Vega",
		Receptor: "Cristian Lopez",
		Fecha:    "2021-12-17",
	}

	repo := NewRepositorySql()

	service := NewServiceSql(repo)

	// Definimos el context.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	obtainedPayment, err := service.GetByIdWithContext(ctx, 12)

	fmt.Println(err)

	assert.Equal(t, newPayment.Codigo, obtainedPayment.Codigo)
	assert.Equal(t, newPayment.Moneda, obtainedPayment.Moneda)
}

func TestUpdateWithContextServiceSql(t *testing.T) {
	expectedPayment := models.Payment{
		Id:       12,
		Codigo:   "AAA002",
		Moneda:   "R$",
		Monto:    float64(295.25),
		Emisor:   "Rodrigo Vega Gimenez",
		Receptor: "Cristian Lopez Diaz",
		Fecha:    "2021-12-28",
	}

	repo := NewRepositorySql()

	service := NewServiceSql(repo)

	// Definimos el context.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updatedPayment, err := service.UpdateWithContext(ctx, expectedPayment)

	fmt.Println(err)

	assert.Equal(t, expectedPayment.Moneda, updatedPayment.Moneda)
	assert.Equal(t, expectedPayment.Monto, updatedPayment.Monto)
	assert.Equal(t, expectedPayment.Receptor, updatedPayment.Receptor)
	assert.Nil(t, err)
}

// ==================================================================================
// ACA COMIENZA EL TESTING DE SQL USANDO MOCKS. INCLUYE ACTIVIDAD.
func TestServiceSqlStore_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	newPayment := models.Payment{
		Codigo:   "AAA002",
		Moneda:   "ARS",
		Monto:    float64(856.34),
		Emisor:   "Rodrigo Vega",
		Receptor: "Matias Perez",
		Fecha:    "2021-12-28",
	}

	rows := sqlmock.NewRows([]string{"id", "codigo", "moneda", "monto", "emisor", "receptor", "fecha"})
	rows.AddRow(1, "AAA002", "ARS", 856.34, "Rodrigo Vega", "Matias Perez", "2021-12-28")

	mock.ExpectPrepare("INSERT INTO Payment")
	mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
	//mock.ExpectQuery("SELECT id, codigo, moneda, monto, emisor, receptor, fecha FROM Payments WHERE id = ?").WithArgs(1).WillReturnRows(rows)

	repo := NewRepositorySqlMock(db)
	service := NewServiceSql(repo)

	insertedPayment, err := service.Store(newPayment.Codigo, newPayment.Moneda, newPayment.Emisor, newPayment.Receptor, newPayment.Fecha, newPayment.Monto)

	assert.Nil(t, err)
	assert.Equal(t, newPayment.Codigo, insertedPayment.Codigo)
}

func TestServiceSqlStore_Mock_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	newPayment := models.Payment{
		Codigo:   "AAA003",
		Moneda:   "ARS",
		Monto:    float64(856.34),
		Emisor:   "Rodrigo Vega",
		Receptor: "Matias Perez",
		Fecha:    "2021-12-28",
	}

	rows := sqlmock.NewRows([]string{"id", "codigo", "moneda", "monto", "emisor", "receptor", "fecha"})
	rows.AddRow(1, "AAA001", "ARS", 856.34, "Rodrigo Vega", "Matias Perez", "2021-12-28")
	rows.AddRow(2, "AAA002", "ARS", 1856.34, "Rodrigo Vega", "Julian Mondaca", "2021-12-28")

	mock.ExpectPrepare("INSERT INTO Payment")
	mock.ExpectExec("").WillReturnError(errors.New("No se pudo registrar la transacción."))
	//mock.ExpectQuery("SELECT id, codigo, moneda, monto, emisor, receptor, fecha FROM Payments WHERE id = ?").WithArgs(1).WillReturnRows(rows)

	repo := NewRepositorySqlMock(db)
	service := NewServiceSql(repo)

	_, err0 := service.Store(newPayment.Codigo, newPayment.Moneda, newPayment.Emisor, newPayment.Receptor, newPayment.Fecha, newPayment.Monto)

	assert.Equal(t, "No se pudo registrar la transacción.", err0.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestServiceSqlGetById_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "codigo", "moneda", "monto", "emisor", "receptor", "fecha"})
	rows.AddRow(1, "AAA002", "ARS", 856.34, "Rodrigo Vega", "Matias Perez", "2021-12-28")

	mock.ExpectQuery("SELECT id, codigo, moneda, monto, emisor, receptor, fecha FROM Payments WHERE id = ?").WithArgs(1).WillReturnRows(rows)

	repo := NewRepositorySqlMock(db)
	service := NewServiceSql(repo)

	obtainedPayment := service.GetById(1)

	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, "AAA002", obtainedPayment.Codigo)
	assert.Equal(t, "ARS", obtainedPayment.Moneda)
}

func TestServiceSqlGetById_Mock_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "codigo", "moneda", "monto", "emisor", "receptor", "fecha"})

	mock.ExpectQuery("SELECT id, codigo, moneda, monto, emisor, receptor, fecha FROM Payments WHERE id = ?").WithArgs(1).WillReturnRows(rows)
	mock.ExpectExec("").WillReturnResult(nil)

	repo := NewRepositorySqlMock(db)
	service := NewServiceSql(repo)

	obtainedPayment := service.GetById(1)

	assert.Equal(t, models.Payment{}, obtainedPayment)
}

func TestServiceSqlUpdate_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	newPayment := models.Payment{
		Id:       1,
		Codigo:   "AAA002",
		Moneda:   "R$",
		Monto:    float64(256.34),
		Emisor:   "Rodrigo Vega Gimenez",
		Receptor: "Matias Agustin Perez",
		Fecha:    "2021-12-28",
	}

	rows := sqlmock.NewRows([]string{"id", "codigo", "moneda", "monto", "emisor", "receptor", "fecha"})
	rows.AddRow(1, "AAA002", "ARS", 856.34, "Rodrigo Vega", "Matias Perez", "2021-12-28")

	mock.ExpectPrepare("UPDATE Payments")
	mock.ExpectQuery("UPDATE Payments SET id = ?, codigo = ?, moneda = ?, monto = ?, emisor = ?, receptor = ?, fecha = ? WHERE id = ?").WithArgs(1).WillReturnRows(rows)

	repo := NewRepositorySqlMock(db)
	service := NewServiceSql(repo)

	changedPayment, err := service.Update(newPayment)

	assert.Nil(t, err)
	assert.Equal(t, newPayment.Codigo, changedPayment.Codigo)
	assert.Equal(t, newPayment.Moneda, changedPayment.Moneda)
	assert.Equal(t, newPayment.Emisor, changedPayment.Emisor)
}

// ==================================================================================
// ACA COMIENZA EL TEST USANDO LA DB COMO UNA BASE DE DATOS TRANSACCIONAL, USANDO TXDB. INCLUYE ACTIVIDAD.
func TestServiceSqlStore_Txdb(t *testing.T) {
	newPayment := models.Payment{
		Codigo:   "AAA002",
		Moneda:   "ARS",
		Monto:    float64(856.34),
		Emisor:   "Rodrigo Vega",
		Receptor: "Matias Perez",
		Fecha:    "2021-12-28",
	}

	// Aca abrimos la comunicacion con TXDB.
	db, err := db.InitDb()
	assert.Nil(t, err)

	repo := NewRepositorySqlMock(db)
	service := NewServiceSql(repo)

	insertedPayment, err := service.Store(newPayment.Codigo, newPayment.Moneda, newPayment.Emisor, newPayment.Receptor, newPayment.Fecha, newPayment.Monto)

	assert.Nil(t, err)
	assert.Equal(t, newPayment.Codigo, insertedPayment.Codigo)
}
