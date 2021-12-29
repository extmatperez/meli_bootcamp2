package internal

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/TTarde/proyecto/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/18_storage2/TTarde/proyecto/pkg/db"
	"github.com/extmatperez/meli_bootcamp2/18_storage2/TTarde/proyecto/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestServiceUpdateMock(t *testing.T) {
	dataByte := []byte(per)

	trNuevo := Transaccion{
		ID:                1,
		CodigoTransaccion: 556111,
		Moneda:            "Pesos",
		Monto:             80.00,
		Emisor:            "Locomotion",
		Receptor:          "Disney",
		FechaTransaccion:  "13/08/2021",
	}

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)
	service := NewService(repo)

	per_actualizada, _ := service.Update(trNuevo.ID, trNuevo.CodigoTransaccion, trNuevo.Moneda, trNuevo.Monto, trNuevo.Emisor, trNuevo.Receptor, trNuevo.FechaTransaccion)

	assert.Equal(t, trNuevo, per_actualizada)



}


func TestServiceDeleteMock(t *testing.T) {
	dataByte := []byte(perso)
  
	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)
	service := NewService(repo)

	err := service.Delete(1)
	todos, _ := service.GetAll()
  
	assert.Equal(t, err, nil)
	assert.Equal(t, len(todos), 1)
  }

  func TestStoreSQL(t *testing.T) {
	  transaccionNueva := models.Transaccion{
		CodigoTransaccion: 556111,
		Moneda:            "Pesos",
		Monto:             80.00,
		Emisor:            "Locomotion",
		Receptor:          "Disney",
		FechaTransaccion:  "13/08/2021",
	  }

	  repo := NewRepositorySQL()

	  service := NewServiceSQL(repo)

	  transaccionCreada, _ := service.Store(transaccionNueva.CodigoTransaccion, transaccionNueva.Moneda, transaccionNueva.Monto, transaccionNueva.Emisor, transaccionNueva.Receptor, transaccionNueva.FechaTransaccion)

	  assert.Equal(t, transaccionNueva.Emisor, transaccionCreada.Emisor)
  }


  func TestGetByNameSQL(t *testing.T) {
	transaccionNueva := models.Transaccion{
	  CodigoTransaccion: 556111,
	  Moneda:            "Pesos",
	  Monto:             80.00,
	  Emisor:            "Locomotion",
	  Receptor:          "Disney",
	  FechaTransaccion:  "13/08/2021",
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	transaccionCargada := service.GetByName("Locomotion")

	assert.Equal(t, transaccionNueva.Emisor, transaccionCargada.Emisor)
}


func TestStoreServiceSQLTxdb(t *testing.T) {
	transaccionNueva := models.Transaccion{
		CodigoTransaccion: 556111,
		Moneda:            "Pesos",
		Monto:             80.00,
		Emisor:            "Mars",
		Receptor:          "Line",
		FechaTransaccion:  "13/08/2021",
	  }

	  db, err := db.InitDb()
	  assert.Nil(t, err)

	  repo := NewRepositorySQLMock(db)
	  defer db.Close()

	  service := NewServiceSQL(repo)

	  transaccionCreada, _ := service.Store(transaccionNueva.CodigoTransaccion, transaccionNueva.Moneda, transaccionNueva.Monto, transaccionNueva.Emisor, transaccionNueva.Receptor, transaccionNueva.FechaTransaccion)

	  assert.Equal(t, transaccionNueva.Emisor, transaccionCreada.Emisor)
}


func TestGetByNameSQLTxdb(t *testing.T) {
	transaccionNueva := models.Transaccion{
	  CodigoTransaccion: 556111,
	  Moneda:            "Pesos",
	  Monto:             80.00,
	  Emisor:            "Locomotion",
	  Receptor:          "Disney",
	  FechaTransaccion:  "13/08/2021",
	}

	db, err := db.InitDb()
	assert.Nil(t, err)

	repo := NewRepositorySQLMock(db)

	service := NewServiceSQL(repo)

	transaccionCargada := service.GetByName("Locomotion")

	assert.Equal(t, transaccionNueva.Emisor, transaccionCargada.Emisor)
}


func TestUpdateSQLTxdb(t *testing.T) {
	trNuevo := models.Transaccion{
		ID:                1,
		CodigoTransaccion: 556111,
		Moneda:            "Sopes",
		Monto:             80.00,
		Emisor:            "Locomotion",
		Receptor:          "Disney",
		FechaTransaccion:  "13/08/2021",
		Articulo: models.Producto{},
	}

	db, err := db.InitDb()
	assert.Nil(t, err)

	repo := NewRepositorySQLMock(db)
	defer db.Close()

	service := NewServiceSQL(repo)


	per_actualizada, _ := service.Update(trNuevo)

	assert.Equal(t, trNuevo, per_actualizada)

}


func TestGetByNameSQLMock(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "codigo_transaccion", "moneda", "monto", "emisor", "receptor", "fecha_transaccion"})
	rows.AddRow(1, 5444, "pesos", 5222.00, "Carlos", "Pancho", "12/05/2021")
	mock.ExpectQuery("SELECT id, codigo_transaccion, moneda, monto, emisor, receptor, fecha_transaccion FROM transacciones WHERE emisor = ?").WithArgs("Carlos").WillReturnRows(rows)

	repo := NewRepositorySQLMock(db)

	service := NewServiceSQL(repo)

	transaccionCargada := service.GetByName("Carlos")

	assert.Equal(t, "Carlos", transaccionCargada.Emisor)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestStoreSQLMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	trNuevo := models.Transaccion{
		CodigoTransaccion: 556111,
		Moneda:            "Sopes",
		Monto:             80.00,
		Emisor:            "LPeren",
		Receptor:          "Daios",
		FechaTransaccion:  "13/08/2021",
		Articulo: models.Producto{},
	}


	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(2, 1))

	repo := NewRepositorySQLMock(db)
	service := NewServiceSQL(repo)

	transaccionCargada, err := service.Store(trNuevo.CodigoTransaccion, trNuevo.Moneda, trNuevo.Monto, trNuevo.Emisor, trNuevo.Receptor, trNuevo.FechaTransaccion)

	assert.Equal(t, 2, transaccionCargada.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
	

}