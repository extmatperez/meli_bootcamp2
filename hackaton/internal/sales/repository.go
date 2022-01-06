package internal

import (
	"errors"
	"log"
	"strings"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/pkg/db"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/pkg/store"
)

var sales []models.Sale

type SaleRepository interface {
	ImportAllSales() error
	StoreSale(models.Sale) (models.Sale, error)
	UpdateSale(models.Sale) (models.Sale, error)
}

type repository_sale struct {
	arr store.SaveFile
}

func NewSaleRepository(arr store.SaveFile) SaleRepository {
	return &repository_sale{arr}
}

func (r *repository_sale) ImportAllSales() error {
	sales_string, err := r.arr.ReadLines("/Users/rovega/Documents/GitHub/meli_bootcamp2/hackaton/cmd/server/data/sales.txt")
	if err != nil {
		return err
	}

	for _, sale := range sales_string {
		only_sale := strings.Split(sale, "#$%#")
		id := only_sale[0]
		id_product := only_sale[1]
		id_invoice := only_sale[2]
		quantity := only_sale[3]

		db := db.StorageDB
		query := "INSERT INTO Sale(id, idProduct, idInvoice, quantity) VALUES (?,?,?,?)"
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(id, id_product, id_invoice, quantity)

		if err != nil {
			return errors.New("No se pudo guardar elemento en BD.")
		}
	}
	return nil
}

func (r *repository_sale) StoreSale(sale models.Sale) (models.Sale, error) {
	db := db.StorageDB
	query := "INSERT INTO Sale(idProduct, idInvoice, quantity) VALUES (?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(sale.IdProduct, sale.IdInvoice, sale.Quantity)
	if err != nil {
		return models.Sale{}, err
	}

	idCreado, _ := result.LastInsertId()
	sale.Id = int(idCreado)
	return sale, nil
}

func (r *repository_sale) UpdateSale(sale models.Sale) (models.Sale, error) {
	db := db.StorageDB
	query := "UPDATE Sale SET idProduct = ?, idInvoice = ?, quantity = ? WHERE id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(sale.IdProduct, sale.IdInvoice, sale.Quantity, sale.Id)
	if err != nil {
		return models.Sale{}, err
	}
	updatedRows, _ := result.RowsAffected()
	if updatedRows == 0 {
		return models.Sale{}, errors.New("No se encontro al sale.")
	}
	return sale, nil
}
