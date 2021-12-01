package internal

import (
	"fmt"
	"reflect"
	"strings"
)

type Service interface {
	GetAll() ([]Product, error)
	Store(name string, color string, stock int, code string, published bool, created_at string) (Product, error)
	FindById(id int64) (Product, error)
	FilterProducts(allProducts []Product, queryParams map[string]string) []Product
	LoadProducts() error
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

func (s *service) Store(name string, color string, stock int, code string, published bool, created_at string) (Product, error) {
	newId, err := s.repository.LastId()

	if err != nil {
		return Product{}, err
	}

	product, err := s.repository.Store(newId+1, name, color, stock, code, published, created_at)

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

func (s *service) LoadProducts() error {
	err := s.repository.LoadProducts()

	if err != nil {
		return err
	}

	return nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
