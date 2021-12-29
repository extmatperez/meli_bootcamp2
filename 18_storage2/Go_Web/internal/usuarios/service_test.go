package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/18_storage2/Go_Web/internal/models"
	db "github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/18_storage2/Go_Web/pkg/db"
	"github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/18_storage2/Go_Web/pkg/store"
	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
	usedGetAll bool
}

var userss string = `
[{"id":1,"nombre":"Ida","apellido":"Tieman","email":"itieman0@npr.org","edad":82,"altura":187,"activo":true,"fecha_creacion":"06/15/2021"},
{"id":2,"nombre":"Law","apellido":"Lafee","email":"llafee1@barnesandnoble.com","edad":70,"altura":142,"activo":true,"fecha_creacion":"07/12/2021"}]

`

func (s *StubRepository) GetAll() ([]Usuario, error) {
	var salida []Usuario
	err := json.Unmarshal([]byte(userss), &salida)
	s.usedGetAll = true
	return salida, err
}

func (s *StubRepository) Store(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error) {
	return Usuario{}, nil
}

func (s *StubRepository) Update(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error) {
	return Usuario{}, nil
}

func (s *StubRepository) Delete(id int) error {
	return nil
}

func (s *StubRepository) EditarNombreEdad(id int, nombre string, edad int) (Usuario, error) {
	return Usuario{}, nil
}

func (s *StubRepository) LastID() (int, error) {
	return 0, nil
}

func TestGetAllService(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	misUsuarios, _ := service.GetAll()
	assert.Equal(t, 2, len(misUsuarios))
	assert.True(t, stubRepo.usedGetAll)
}

func TestGetAllServiceMock(t *testing.T) {
	dataByte := []byte(userss)
	var usuariosEsperados []Usuario
	json.Unmarshal(dataByte, &usuariosEsperados)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	misUsuarios, _ := service.GetAll()
	assert.Equal(t, usuariosEsperados, misUsuarios)
}
func TestGetAllServiceMockError(t *testing.T) {
	// dataByte := []byte(userss)
	// var usuariosEsperados []Usuario
	// json.Unmarshal(dataByte, &usuariosEsperados)
	errorEsperado := errors.New("No hay datos en el mock")

	dbMock := store.Mock{Err: errorEsperado}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	misUsuarios, errorRecibido := service.GetAll()
	assert.Equal(t, errorEsperado, errorRecibido)
	assert.Nil(t, misUsuarios)
}

func TestStoreServiceMock(t *testing.T) {
	// Arrange
	usuarioNuevo := Usuario{3, "Juan", "Perez", "correo", 20, 180, true, "fecha"}

	dbMock := store.Mock{Data: []byte(`[]`)}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	usuarioCreado, _ := service.Store(usuarioNuevo.Nombre, usuarioNuevo.Apellido, usuarioNuevo.Email, usuarioNuevo.Edad, usuarioNuevo.Altura, usuarioNuevo.Activo, usuarioNuevo.FechaCreacion)
	assert.Equal(t, usuarioNuevo.Nombre, usuarioCreado.Nombre)
	assert.Equal(t, usuarioNuevo.Altura, usuarioCreado.Altura)
	//assert.Nil(t, misUsuarios)
}

func TestStoreServiceMockError(t *testing.T) {
	// Arrange
	usuarioNuevo := Usuario{3, "Juan", "Perez", "correo", 20, 180, true, "fecha"}

	errorEsperado := errors.New("No hay datos en el mock")

	dbMock := store.Mock{Data: []byte(`[]`), Err: errorEsperado}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	usuarioCreado, err := service.Store(usuarioNuevo.Nombre, usuarioNuevo.Apellido, usuarioNuevo.Email, usuarioNuevo.Edad, usuarioNuevo.Altura, usuarioNuevo.Activo, usuarioNuevo.FechaCreacion)
	assert.Equal(t, errorEsperado, err)
	assert.Equal(t, "", usuarioCreado.Nombre)
}

func TestUpdateServiceMock(t *testing.T) {
	// Arrange
	usuarioUp := Usuario{2, "Juan", "Perez", "correo", 20, 180, true, "fecha"}

	dataByte := []byte(userss)
	var usuariosEsperados []Usuario
	json.Unmarshal(dataByte, &usuariosEsperados)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	usuarioActualizado, _ := service.Update(usuarioUp.ID, usuarioUp.Nombre, usuarioUp.Apellido, usuarioUp.Email, usuarioUp.Edad, usuarioUp.Altura, usuarioUp.Activo, usuarioUp.FechaCreacion)
	assert.Equal(t, usuarioUp, usuarioActualizado)
	assert.True(t, dbMock.ValidRead)
}

func TestUpdateServiceMockError(t *testing.T) {
	// Arrange
	usuarioUp := Usuario{2, "Juan", "Perez", "correo", 20, 180, true, "fecha"}

	dataByte := []byte(userss)
	var usuariosEsperados []Usuario
	json.Unmarshal(dataByte, &usuariosEsperados)

	errorEsperado := errors.New("No hay datos en el mock")

	dbMock := store.Mock{Data: []byte(`[]`), Err: errorEsperado}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.Update(usuarioUp.ID, usuarioUp.Nombre, usuarioUp.Apellido, usuarioUp.Email, usuarioUp.Edad, usuarioUp.Altura, usuarioUp.Activo, usuarioUp.FechaCreacion)
	assert.Equal(t, errorEsperado, err)
}
func TestDeleteServiceMock(t *testing.T) {
	// Arrange

	dataByte := []byte(userss)
	var usuariosEsperados []Usuario
	json.Unmarshal(dataByte, &usuariosEsperados)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	err := service.Delete(1)
	assert.Nil(t, err)
}
func TestDeleteServiceMockError(t *testing.T) {
	// Arrange

	dataByte := []byte(userss)
	var usuariosEsperados []Usuario
	json.Unmarshal(dataByte, &usuariosEsperados)
	errorEsperado := errors.New("No hay datos en el mock")

	dbMock := store.Mock{Data: dataByte, Err: errorEsperado}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	err := service.Delete(5)
	assert.Equal(t, errorEsperado, err)
}

func TestDeleteServiceMockError2(t *testing.T) {
	// Arrange

	dataByte := []byte(userss)
	var usuariosEsperados []Usuario
	json.Unmarshal(dataByte, &usuariosEsperados)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	err := service.Delete(5)
	assert.NotNil(t, err)
}
func TestDeleteService(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	err := service.Delete(1)

	assert.Nil(t, err)
}

func TestStoreServiceSQL(t *testing.T) {
	// Arrange
	usuarioNuevo := models.Usuario{Nombre: "Juan", Apellido: "Perez", Email: "correo", Edad: 20, Altura: 180, Activo: true, FechaCreacion: "fecha"}

	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	usuarioCreado, _ := service.Store(usuarioNuevo.Nombre, usuarioNuevo.Apellido, usuarioNuevo.Email, usuarioNuevo.Edad, usuarioNuevo.Altura, usuarioNuevo.Activo, usuarioNuevo.FechaCreacion)
	assert.Equal(t, usuarioNuevo.Nombre, usuarioCreado.Nombre)
	assert.Equal(t, usuarioNuevo.Altura, usuarioCreado.Altura)
	//assert.Nil(t, misUsuarios)
}

func TestGetOneServiceSQL(t *testing.T) {
	//Arrange
	searchedUser := models.Usuario{
		Nombre:        "Juan",
		Apellido:      "Perez",
		Email:         "correo",
		Edad:          20,
		Altura:        180,
		Activo:        true,
		FechaCreacion: "fecha",
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	personaCargada := service.GetOne(1)

	assert.Equal(t, searchedUser.Nombre, personaCargada.Nombre)
	assert.Equal(t, searchedUser.Apellido, personaCargada.Apellido)
	// assert.Nil(t, misPersonas)
}

func TestGetByNameServiceSQL(t *testing.T) {
	//Arrange
	user := []models.Usuario{
		{ID: 1,
			Nombre:        "Juan",
			Apellido:      "Kevin",
			Email:         "correo",
			Edad:          20,
			Altura:        180,
			Activo:        true,
			FechaCreacion: "fecha"},
		{ID: 2,
			Nombre:        "Juan",
			Apellido:      "Perez",
			Email:         "correo",
			Edad:          20,
			Altura:        180,
			Activo:        true,
			FechaCreacion: "fecha"},
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	personaCargada, _ := service.GetByName(user[0].Nombre)
	fmt.Println(personaCargada)
	assert.Equal(t, user, personaCargada)

}

func TestUpdateServiceSQL(t *testing.T) {
	//Arrange
	userUpdate := models.Usuario{
		ID:            1,
		Nombre:        "Juan",
		Apellido:      "Kevin",
		Email:         "correo",
		Edad:          20,
		Altura:        180,
		Activo:        true,
		FechaCreacion: "fecha",
	}

	repo := NewRepositorySQL()

	// service := NewServiceSQL(repo)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	personaCargada, err := repo.Update(ctx, userUpdate)
	fmt.Println(err)
	assert.Equal(t, userUpdate.Nombre, personaCargada.Nombre)
	assert.Equal(t, userUpdate.Apellido, personaCargada.Apellido)
	// assert.Nil(t, misPersonas)
}

func TestGetAllServiceSQL(t *testing.T) {
	//Arrange
	user := []models.Usuario{
		{ID: 1,
			Nombre:        "Juan",
			Apellido:      "Kevin",
			Email:         "correo",
			Edad:          20,
			Altura:        180,
			Activo:        true,
			FechaCreacion: "fecha"},
		{ID: 2,
			Nombre:        "Juan",
			Apellido:      "Perez",
			Email:         "correo",
			Edad:          20,
			Altura:        180,
			Activo:        true,
			FechaCreacion: "fecha"},
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	usuarios, _ := service.GetAll()
	fmt.Println(usuarios)
	assert.Equal(t, user, usuarios)

}

func TestDeleteServiceSQL(t *testing.T) {
	//Arrange

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	err := service.Delete(5)

	assert.Nil(t, err)

}

func TestGetOneServiceSQLMock1(t *testing.T) {
	//Arrange
	searchedUser := models.Usuario{
		Nombre:        "Juan",
		Apellido:      "Kevin",
		Email:         "correo",
		Edad:          20,
		Altura:        180,
		Activo:        true,
		FechaCreacion: "fecha",
	}
	db := db.StorageDB
	repo := NewRepositorySQLMock(db)

	service := NewServiceSQL(repo)

	personaCargada := service.GetOne(1)

	assert.Equal(t, searchedUser.Nombre, personaCargada.Nombre)
	assert.Equal(t, searchedUser.Apellido, personaCargada.Apellido)
	// assert.Nil(t, misPersonas)
}

func TestGetOneServiceSQLMock(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "nombre", "apellido", "email", "edad", "altura", "activo", "fecha_creacion"})
	rows.AddRow(1, "Juan", "Kevin", "correo", 20, 180, true, "fecha")
	mock.ExpectQuery("SELECT id, nombre,apellido, email, edad, altura, activo,fecha_creacion FROM users WHERE id = ?").WithArgs(1).WillReturnRows(rows)

	repo := NewRepositorySQLMock(db)
	service := NewServiceSQL(repo)

	personaCargada := service.GetOne(1)

	assert.Equal(t, "Juan", personaCargada.Nombre)
	assert.Equal(t, "Kevin", personaCargada.Apellido)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestStoreServiceSQLMock(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("").WillReturnResult(sqlmock.NewResult(9, 1))

	usuarioNuevo := models.Usuario{Nombre: "Juan", Apellido: "Perez", Email: "correo", Edad: 20, Altura: 180, Activo: true, FechaCreacion: "fecha"}

	repo := NewRepositorySQLMock(db)
	service := NewServiceSQL(repo)

	usuarioCreado, _ := service.Store(usuarioNuevo.Nombre, usuarioNuevo.Apellido, usuarioNuevo.Email, usuarioNuevo.Edad, usuarioNuevo.Altura, usuarioNuevo.Activo, usuarioNuevo.FechaCreacion)

	assert.Equal(t, "Juan", usuarioCreado.Nombre)
	assert.Equal(t, "Perez", usuarioCreado.Apellido)
	assert.Equal(t, 9, usuarioCreado.ID)
	assert.Nil(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestStoreServiceSQLTXDB(t *testing.T) {
	// Arrange
	usuarioNuevo := models.Usuario{Nombre: "Juan", Apellido: "Perez", Email: "correo", Edad: 20, Altura: 180, Activo: true, FechaCreacion: "fecha"}

	db, err := db.InitDB()
	assert.Nil(t, err)

	repo := NewRepositorySQLMock(db)
	defer db.Close()

	service := NewServiceSQL(repo)

	usuarioCreado, _ := service.Store(usuarioNuevo.Nombre, usuarioNuevo.Apellido, usuarioNuevo.Email, usuarioNuevo.Edad, usuarioNuevo.Altura, usuarioNuevo.Activo, usuarioNuevo.FechaCreacion)
	assert.Equal(t, usuarioNuevo.Nombre, usuarioCreado.Nombre)
	assert.Equal(t, usuarioNuevo.Altura, usuarioCreado.Altura)
	//assert.Nil(t, misUsuarios)
}

func TestStore_GetOneServiceSQLTXDB(t *testing.T) {
	// Arrange
	usuarioNuevo := models.Usuario{Nombre: "Juan", Apellido: "Perez", Email: "correo", Edad: 20, Altura: 180, Activo: true, FechaCreacion: "fecha"}

	db, err := db.InitDB()
	assert.Nil(t, err)

	repo := NewRepositorySQLMock(db)
	defer db.Close()

	service := NewServiceSQL(repo)

	usuarioCreado, _ := service.Store(usuarioNuevo.Nombre, usuarioNuevo.Apellido, usuarioNuevo.Email, usuarioNuevo.Edad, usuarioNuevo.Altura, usuarioNuevo.Activo, usuarioNuevo.FechaCreacion)
	assert.Equal(t, usuarioNuevo.Nombre, usuarioCreado.Nombre)
	assert.Equal(t, usuarioNuevo.Altura, usuarioCreado.Altura)

	fmt.Println(usuarioCreado)

	usuarioConsultado := service.GetOne(usuarioCreado.ID)

	fmt.Println(usuarioConsultado)

	assert.Equal(t, usuarioNuevo.Nombre, usuarioConsultado.Nombre)
	assert.Equal(t, usuarioNuevo.Altura, usuarioConsultado.Altura)
	//assert.Nil(t, misUsuarios)
}

func TestUpdate_DeleteServiceSQLTXDB(t *testing.T) {
	// Arrange
	usuarioUpdate := models.Usuario{ID: 4, Nombre: "Pipe", Apellido: "Roldan", Email: "correo", Edad: 20, Altura: 180, Activo: true, FechaCreacion: "fecha"}

	db, err := db.InitDB()
	assert.Nil(t, err)

	repo := NewRepositorySQLMock(db)
	defer db.Close()

	service := NewServiceSQL(repo)

	usuarioActualizado, _ := service.Update(context.Background(), usuarioUpdate)
	assert.Equal(t, usuarioUpdate.Nombre, usuarioActualizado.Nombre)
	assert.Equal(t, usuarioUpdate.Altura, usuarioActualizado.Altura)

	fmt.Println(usuarioActualizado)

	usuarioConsultado := service.GetOne(4)

	fmt.Println(usuarioConsultado)

	assert.Equal(t, usuarioUpdate.Nombre, usuarioConsultado.Nombre)
	assert.Equal(t, usuarioUpdate.Altura, usuarioConsultado.Altura)
	//assert.Nil(t, misUsuarios)
}
func TestUpdateServiceSQLTXDB_Fail(t *testing.T) {
	// Arrange
	usuarioUpdate := models.Usuario{ID: 70, Nombre: "Pipe", Apellido: "Roldan", Email: "correo", Edad: 20, Altura: 180, Activo: true, FechaCreacion: "fecha"}

	db, err := db.InitDB()
	assert.Nil(t, err)

	repo := NewRepositorySQLMock(db)
	defer db.Close()

	service := NewServiceSQL(repo)

	_, err = service.Update(context.Background(), usuarioUpdate)
	assert.Error(t, err)

}

func TestStore_GetOneServiceSQLMock(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	usuarioNuevo := models.Usuario{Nombre: "Juan", Apellido: "Perez", Email: "correo", Edad: 20, Altura: 180, Activo: true, FechaCreacion: "fecha"}

	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(13, 1))

	repo := NewRepositorySQLMock(db)
	defer db.Close()

	service := NewServiceSQL(repo)

	usuarioCreado, _ := service.Store(usuarioNuevo.Nombre, usuarioNuevo.Apellido, usuarioNuevo.Email, usuarioNuevo.Edad, usuarioNuevo.Altura, usuarioNuevo.Activo, usuarioNuevo.FechaCreacion)
	assert.Equal(t, usuarioNuevo.Nombre, usuarioCreado.Nombre)
	assert.Equal(t, usuarioNuevo.Altura, usuarioCreado.Altura)

	fmt.Println(usuarioCreado)

	rows := sqlmock.NewRows([]string{"id", "nombre", "apellido", "email", "edad", "altura", "activo", "fecha_creacion"})
	rows.AddRow(13, "Juan", "Perez", "correo", 20, 180, true, "fecha")
	mock.ExpectQuery("SELECT id, nombre,apellido, email, edad, altura, activo,fecha_creacion FROM users WHERE id = ?").WithArgs(13).WillReturnRows(rows)

	usuarioConsultado := service.GetOne(13)

	fmt.Println(usuarioConsultado)

	assert.Equal(t, usuarioConsultado.Nombre, usuarioCreado.Nombre)
	assert.Equal(t, usuarioConsultado.Altura, usuarioCreado.Altura)
	//assert.Nil(t, misUsuarios)
}
