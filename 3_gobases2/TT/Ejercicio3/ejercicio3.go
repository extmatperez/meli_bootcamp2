package main

import (
	"errors"
	"fmt"
)

const (
	SMALL  = "pequeño"
	MEDIUM = "mediano"
	BIG    = "grande"
)

func main() {
	store := newStore()

	smallProduct, errSmall := newProduct(SMALL, "Caramelo", 10.0)
	mediumProduct, errMedium := newProduct(MEDIUM, "Pelota de futbol", 1300.0)
	bigProduct, errBig := newProduct(BIG, "Heladera", 50000.0)

	if errSmall != nil {
		fmt.Println(errSmall)
	} else {
		store.AddProduct(smallProduct)
	}

	if errMedium != nil {
		fmt.Println(errSmall)
	} else {
		store.AddProduct(mediumProduct)
	}

	if errBig != nil {
		fmt.Println(errSmall)
	} else {
		store.AddProduct(bigProduct)
	}

	fmt.Println(store.Total())
}

type Store struct {
	Products []ProductInterface
}

func (s Store) Total() float64 {
	var total float64
	for i := 0; i < len(s.Products); i++ {
		total += s.Products[i].CalculateCost()
	}

	return total
}

func (s *Store) AddProduct(productInterface ProductInterface) {
	s.Products = append(s.Products, productInterface)
}

type Product struct {
	ProductType string
	Name        string
	Price       float64
}

type SmallProduct struct {
	p Product
}

type MediumProduct struct {
	p Product
}

type BigProduct struct {
	p Product
}

func (p Product) CalculateCost() float64 {
	return p.Price
}

func (sP SmallProduct) CalculateCost() float64 {
	return sP.p.Price
}

func (mP MediumProduct) CalculateCost() float64 {
	return mP.p.Price * 1.03
}

func (bP BigProduct) CalculateCost() float64 {
	return bP.p.Price*1.06 + 2500
}

type ProductInterface interface {
	CalculateCost() float64
}

type EcommerceInterface interface {
	Total() float64
	Add() float64
}

func newProduct(productType, name string, price float64) (ProductInterface, error) {
	switch productType {
	case SMALL:
		return SmallProduct{
			p: Product{
				ProductType: productType,
				Name:        name,
				Price:       price,
			},
		}, nil
	case MEDIUM:
		return MediumProduct{
			p: Product{
				ProductType: productType,
				Name:        name,
				Price:       price,
			},
		}, nil
	case BIG:
		return BigProduct{
			p: Product{
				ProductType: productType,
				Name:        name,
				Price:       price,
			},
		}, nil
	default:
		errorMsg := "el producto '" + productType + "' no existe"
		return Product{}, errors.New(errorMsg)
	}
}

func newStore() Store {
	return Store{}
}
