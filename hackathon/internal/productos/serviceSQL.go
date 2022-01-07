package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/hackathon/internal/models"
)

var print = fmt.Println

type ServiceSQL interface {
	LoadProductsOnDB() error
	Store(id int, description string, price float64) (models.Product, error)
	GetAll() ([]models.Product, error)
	GetById(id int) (models.Product, error)
	// GetLastId() (int, error)
	Update(models.Product) (models.Product, error)
	// UpdateNombrePrecio(id int, nombre string, precio float64) (models.Product, error)
	Delete(id int) error
}
type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}
func (s *serviceSQL) LoadProductsOnDB() error {
	s.repository.DeleteAll()
	byte_products, err := os.ReadFile("./../datos/products.txt")
	if err != nil {
		fmt.Println("error al abrir el archivo: ", err)

	}

	string_products := string(byte_products)
	slice_string_productos := strings.Split(string_products, "\n")
	for _, str := range slice_string_productos {
		print("PRODUCTO: ", str)
		splitted_product := strings.Split(str, "#$%#")
		intID, _ := strconv.Atoi(splitted_product[0])
		description := splitted_product[1]
		float_price, _ := strconv.ParseFloat(splitted_product[2], 64)
		product_to_store := models.Product{ID: intID, Description: description, Price: float_price}
		_, err := s.repository.Store(product_to_store)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *serviceSQL) Store(id int, description string, price float64) (models.Product, error) {
	nuevoProduct := models.Product{ID: id, Description: description, Price: price}
	productoCreado, err := s.repository.Store(nuevoProduct)
	if err != nil {
		return models.Product{}, err
	}
	return productoCreado, nil

}
func (s *serviceSQL) GetByIdProduct(id int) (models.Product, error) {
	producto, err := s.repository.GetById(id)
	if err != nil {
		return models.Product{}, err
	}
	return producto, nil
}

func (s *serviceSQL) GetAll() ([]models.Product, error) {
	productos, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (s *serviceSQL) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *serviceSQL) Update(producto models.Product) (models.Product, error) {
	productoActualizado, err := s.repository.Update(producto)
	if err != nil {
		return models.Product{}, err
	}
	return productoActualizado, nil
}
