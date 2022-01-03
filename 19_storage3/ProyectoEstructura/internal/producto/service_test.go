package internal

import (
	"context"
	"testing"
	"time"

	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/19_storage3/ProyectoEstructura/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/19_storage3/ProyectoEstructura/pkg/db"
	"github.com/stretchr/testify/assert"
)

// var prodServiceData string = `[
// 	{
// 	 "id": 1,
// 	 "nombre": "Espejo",
// 	 "color": "Azul",
// 	 "precio": 40,
// 	 "stock": 10,
// 	 "codigo": "156ujma8ssssDA",
// 	 "publicado": true,
// 	 "fecha_creacion": "25/10/2020"
// 	},
// 	{
// 	 "id": 3,
// 	 "nombre": "Auricular",
// 	 "color": "Verde",
// 	 "precio": 25.3,
// 	 "stock": 10,
// 	 "codigo": "156ujma8ssssDA",
// 	 "publicado": true,
// 	 "fecha_creacion": "25/10/2020"
// 	}
//    ]`

// type storeServiceMock struct {
// 	Data    interface{}
// 	SpyRead bool
// }

// func (s *storeServiceMock) Read(data interface{}) error {
// 	s.SpyRead = true
// 	json.Unmarshal([]byte(prodData), &data)

// 	return nil
// }
// func (s *storeServiceMock) Write(data interface{}) error {
// 	return nil
// }

// func TestUpdateService(t *testing.T) {
// 	//arr
// 	dataByte := []byte(prodData)

// 	dbMock := store.Mock{Data: dataByte}
// 	storeStub := store.FileStore{Mock: &dbMock}
// 	// store := storeServiceMock{}
// 	repo := NewRepository(&storeStub)
// 	ser := NewService(repo)

// 	miProd, err := ser.Update(1, "EspejoUpdated", "Azul", 40, 10, "156ujma8ssssDA", true, "25/10/2020")

// 	var expected []Producto
// 	json.Unmarshal([]byte(prodServiceData), &expected)

// 	assert.Equal(t, expected[0].ID, miProd.ID)
// 	assert.Nil(t, err)
// 	// assert.True(t, store.SpyRead)
// }

func TestStoreServiceSQL(t *testing.T) {
	//Arrange
	repo := newRepositorySQL()

	service := NewServiceSQL(repo)
	prodParaCrear := models.Producto{Nombre: "Leche", Color: "Blanco", Precio: 100, Stock: 4, Codigo: "LEC992", Publicado: true, FechaCreacion: "28/12/2021"}
	prodObtenidos, err := service.Store(prodParaCrear.Nombre, prodParaCrear.Color, prodParaCrear.Precio, prodParaCrear.Stock, prodParaCrear.Codigo, prodParaCrear.Publicado, prodParaCrear.FechaCreacion)

	assert.Nil(t, err)
	assert.Equal(t, prodObtenidos.Nombre, prodParaCrear.Nombre)
}
func TestGetByNameServiceSQL(t *testing.T) {
	//Arrange
	nombreProd := "Yogurt"

	repo := newRepositorySQL()

	service := NewServiceSQL(repo)

	prodObtenidos, err := service.GetByName(nombreProd)

	assert.Nil(t, err)
	assert.True(t, len(prodObtenidos) >= 0)
}
func TestGetByNameServiceSQLFailNotFoundAny(t *testing.T) {
	//Arrange
	nombreProd := "Yogurt"

	repo := newRepositorySQL()

	service := NewServiceSQL(repo)

	prodObtenidos, err := service.GetByName(nombreProd)

	assert.Nil(t, err)
	assert.True(t, len(prodObtenidos) == 0)
}

func TestGetAllSQLContext(t *testing.T) {
	//Arrange
	repo := newRepositorySQL()

	service := NewServiceSQL(repo)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	prodObtenidos, err := service.GetAll(ctx)

	assert.Nil(t, err)
	assert.True(t, len(prodObtenidos) >= 0)
	// assert.Equal(t, .Nombre, .Nombre)
}

func TestUpdateSQLContext(t *testing.T) {
	//Arrange
	repo := newRepositorySQL()

	service := NewServiceSQL(repo)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	prod := models.Producto{
		ID:            1,
		Nombre:        "Yogursin",
		Color:         "Rojo",
		Precio:        120,
		Stock:         5,
		Codigo:        "YOU777",
		Publicado:     true,
		FechaCreacion: "29/12/2021"}

	prodActualizado, err := service.Update(ctx, prod)

	assert.Nil(t, err)
	assert.True(t, prodActualizado == prod)
	// assert.Equal(t, .Nombre, .Nombre)
}

func TestStoreServiceSQLTxdb(t *testing.T) {
	//Arrange

	prodParaCrear := models.Producto{Nombre: "Leche", Color: "Blanco", Precio: 100, Stock: 4, Codigo: "LEC992", Publicado: true, FechaCreacion: "28/12/2021"}

	db, err := db.InitDb()
	assert.Nil(t, err)
	repo := NewRepositorySQLMock(db)
	defer db.Close()

	service := NewServiceSQL(repo)

	prodCreado, _ := service.Store(prodParaCrear.Nombre, prodParaCrear.Color, prodParaCrear.Precio, prodParaCrear.Stock, prodParaCrear.Codigo, prodParaCrear.Publicado, prodParaCrear.FechaCreacion)

	assert.Equal(t, prodParaCrear.Nombre, prodCreado.Nombre)
	assert.Equal(t, prodParaCrear.Color, prodCreado.Color)
}
