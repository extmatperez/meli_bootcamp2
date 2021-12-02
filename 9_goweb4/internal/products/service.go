package internal

import (
	"fmt"
	"reflect"
	"strings"
)

type Service interface {
	GetAll() ([]Product, error)
	Store(name string, color string, price float64, stock int, code string, published bool, created_at string) (Product, error)
	FindById(id int64) (Product, error)
	FilterProducts(allProducts []Product, queryParams map[string]string) []Product
	Update(id int64, name string, color string, price float64, stock int, code string, published bool, created_at string) (Product, error)
	Delete(id int64) error
	UpdateNameAndPrice(id int64, name string, price float64) (Product, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Store(name string, color string, price float64, stock int, code string, published bool, created_at string) (Product, error) {
	newId, err := s.repository.LastId()

	if err != nil {
		return Product{}, err
	}

	product, err := s.repository.Store(newId+1, name, color, price, stock, code, published, created_at)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) FindById(id int64) (Product, error) {
	product, err := s.repository.FindById(id)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) FilterProducts(allProducts []Product, queryParams map[string]string) []Product {
	var fields []string

	//TODO: at the moment only filter strings
	fields = append(fields, "name", "color", "code")

	productsFiltered := allProducts

	for _, field := range fields {
		if len(queryParams[field]) != 0 && len(productsFiltered) != 0 {
			productsFiltered = filterProductsByField(productsFiltered, field, queryParams[field])
		}
	}

	return productsFiltered
}

func filterProductsByField(productsToFilter []Product, field, fieldValue string) []Product {
	var filteredProducts []Product
	var product Product

	productTypeOf := reflect.TypeOf(product)
	fieldIndex := 0

	for fieldIndex = 0; fieldIndex < productTypeOf.NumField(); fieldIndex++ {
		if strings.ToLower(productTypeOf.Field(fieldIndex).Name) == field {
			break
		}
	}

	// If index is out of range (field not found), return the same slice of products
	if fieldIndex == productTypeOf.NumField() {
		return productsToFilter
	}

	for _, productLoop := range productsToFilter {
		productFieldVal := fmt.Sprintf("%v", reflect.ValueOf(productLoop).Field(fieldIndex).Interface())
		if strings.Contains(strings.ToLower(productFieldVal), strings.ToLower(fieldValue)) {
			filteredProducts = append(filteredProducts, productLoop)
		}
	}

	return filteredProducts
}

func (s *service) Update(id int64, name string, color string, price float64, stock int, code string, published bool, created_at string) (Product, error) {
	return s.repository.Update(id, name, color, price, stock, code, published, created_at)
}

func (s *service) UpdateNameAndPrice(id int64, name string, price float64) (Product, error) {
	return s.repository.UpdateNameAndPrice(id, name, price)
}

func (s *service) Delete(id int64) error {
	return s.repository.Delete(id)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
