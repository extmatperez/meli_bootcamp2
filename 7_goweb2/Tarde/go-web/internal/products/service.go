package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type Service interface {
	GetAll() ([]Product, error)
	Store(nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (Product, error)
	FindById(id int) (Product, error)
	Update(id int, nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (Product, error)
	Delete(id int) error
	UpdateNameAndPrice(id int, nombre string, precio int) (Product, error)
	Filter(params map[string][]string) ([]*Product, error)
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{repo: rep}
}

func (serv *service) GetAll() ([]Product, error) {
	prods, err := serv.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("No se puedo extraer los datos %w", err)
	} else {
		if prods != nil {
			return prods, nil
		} else {
			return []Product{}, nil
		}
	}
}

func (serv *service) Store(nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (Product, error) {

	id, err := serv.repo.LastID()
	if err != nil {
		return Product{}, fmt.Errorf("No se pudo crear el producto %w", err)
	}

	prod, err := serv.repo.Store(id+1, nombre, color, precio, stock, codigo, publicado, fechaCreacion)

	if err != nil {
		return Product{}, fmt.Errorf("No se pudo crear el producto %w", err)
	}

	return prod, nil

}

func (serv *service) FindById(id int) (Product, error) {
	prods, err := serv.repo.GetAll()
	if err != nil {
		return Product{}, fmt.Errorf("No se puedo extraer los datos %w", err)
	}

	for _, value := range prods {
		if value.ID == id {
			return value, nil
		}
	}

	return Product{}, fmt.Errorf("Producto id: %d no encontrado", id)

}

func (serv *service) Update(id int, nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	prod, err := serv.repo.Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)

	if err != nil {
		return prod, fmt.Errorf("No se pudo actualizar el producto, %w", err)
	}

	return prod, nil
}

func (serv *service) Delete(id int) error {
	err := serv.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("No se pudo eliminar el producto, %w", err)
	}
	return nil
}

func (serv *service) UpdateNameAndPrice(id int, nombre string, precio int) (Product, error) {

	prod, err := serv.repo.UpdateNameAndPrice(id, nombre, precio)

	if err != nil {
		return Product{}, fmt.Errorf("No se pudo actualizar el producto, %w", err)
	}

	return prod, nil

}

func (serv *service) Filter(params map[string][]string) ([]*Product, error) {

	var filtrados []*Product

	products, err := serv.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("No se puedo extraer los datos %w", err)
	}

	for i := 0; i < len(products); i++ {
		var coincide []bool

		if params["id"] == nil {
			coincide = append(coincide, true)
		} else {
			id, err := strconv.Atoi(params["id"][0])
			if err == nil && id == products[i].ID {
				coincide = append(coincide, true)
			}
		}
		if params["nombre"] == nil {
			coincide = append(coincide, true)
		} else {
			if strings.Contains(strings.ToLower(products[i].Nombre), strings.ToLower(params["nombre"][0])) {
				coincide = append(coincide, true)
			}
		}

		if params["color"] == nil {
			coincide = append(coincide, true)
		} else {
			fmt.Println(products[i].Color, params["color"][0])
			if strings.Contains(strings.ToLower(products[i].Color), strings.ToLower(params["color"][0])) {
				coincide = append(coincide, true)
			}
		}

		if params["precio"] == nil {
			coincide = append(coincide, true)
		} else {

			if len(params["precio"]) == 2 {
				precioFrom := 0
				precioFrom, err := strconv.Atoi(params["precio"][0])
				if err != nil {
					coincide = append(coincide, true)
				} else {
					precioTo, err := strconv.Atoi(params["precio"][1])
					if err != nil {
						coincide = append(coincide, true)
					} else {
						if precioFrom <= products[i].Precio && products[i].Precio <= precioTo {
							coincide = append(coincide, true)
						}
					}
				}
			} else {
				precioTo, err := strconv.Atoi(params["precio"][0])
				if err == nil && products[i].Precio <= precioTo {
					coincide = append(coincide, true)
				}
			}

		}
		if params["stock"] == nil {
			coincide = append(coincide, true)
		} else {

			stockTo := 0
			if len(params["stock"]) == 2 {
				stockFrom := 0
				stockFrom, err := strconv.Atoi(params["stock"][0])
				if err != nil {
					coincide = append(coincide, true)
				} else {
					stockTo, err := strconv.Atoi(params["stock"][1])
					if err != nil {
						coincide = append(coincide, true)
					} else {
						if stockFrom <= products[i].Stock && products[i].Stock <= stockTo {
							coincide = append(coincide, true)
						}
					}
				}
			}
			stockTo, err := strconv.Atoi(params["stock"][0])
			if err == nil && products[i].Stock <= stockTo {
				coincide = append(coincide, true)
			}
		}

		if params["codigo"] == nil {
			coincide = append(coincide, true)
		} else if params["codigo"][0] == products[i].Codigo {
			coincide = append(coincide, true)
		}
		if params["publicado"] != nil {
			if params["publicado"][0] == "true" && products[i].Publicado {
				coincide = append(coincide, true)
			} else if params["publicado"][0] == "false" && !products[i].Publicado {
				coincide = append(coincide, true)
			}
		} else {
			coincide = append(coincide, true)
		}
		if params["fechaCreacion"] == nil {
			coincide = append(coincide, true)
		} else if params["fechaCreacion"][0] == products[i].FechaCreacion {
			coincide = append(coincide, true)
		}
		//	fmt.Printf("len= %d idCoincide=", len(coincide), params["id") == products[i].ID)
		if len(coincide) > 7 {
			filtrados = append(filtrados, &products[i])
		}
	}

	return filtrados, nil

}
