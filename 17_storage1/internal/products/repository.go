package internal

import (
	"database/sql"
	"log"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/internal/models"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/pkg/db"
)

type Repository interface {
	Store(product models.Product) (models.Product, error)
	// GetAll()
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (rep *repository) Store(product models.Product) (models.Product, error) {
	//se ejecuta el init y se crea la conexion ocn la base de datos
	db := db.StorageDB
	//se prepara la consulta
	stmt, err := db.Prepare("INSERT INTO products(name, price, size) VALUES(?, ?, ?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	//me guardo en una variable el result para obtener luego el id creado
	var result sql.Result
	//se ejecuta el statment y devuelve un result y un error
	result, err = stmt.Exec(product.Name, product.Price, product.Size)
	if err != nil {
		return models.Product{}, err
	}

	//accedo al metodo lastInsertID del result, el error no lo controlo porque se supone que ya controle todos los errores
	//EXTRA: result tambine tiene el metodo result.RowsAffected() para saber cuantas filas fueron afectadas
	id_Creado, _ := result.LastInsertId()
	// RECORDAR: el LastId me devuelve un int64 asi que si mi id es int lo casteo a int
	product.ID = int(id_Creado)

	return product, nil

}

// func GetAll(rep *repository) ([]models.Product, error) {
// 	var myProducts []models.Product
// 	db := db.StorageDB

// 	return [].[]models.Product{}, errors.New("err")
// }
