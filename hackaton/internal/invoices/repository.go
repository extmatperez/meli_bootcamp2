package internal

import (
	"errors"
	"log"
	"strings"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/pkg/db"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/pkg/store"
)

type InvoiceRepository interface {
	ImportAllInvoices() error
	StoreInvoice(models.Invoice) (models.Invoice, error)
	UpdateInvoice(models.Invoice) (models.Invoice, error)
	UpdateTotalsOfInvoices() error
}

type repository_invoice struct {
	arr store.SaveFile
}

func NewInvoiceRepository(arr store.SaveFile) InvoiceRepository {
	return &repository_invoice{arr}
}

func (r *repository_invoice) ImportAllInvoices() error {
	invoices_string, err := r.arr.ReadLines("/Users/rovega/Documents/GitHub/meli_bootcamp2/hackaton/cmd/server/data/invoices.txt")
	if err != nil {
		return err
	}

	for _, invoice := range invoices_string {
		only_invoice := strings.Split(invoice, "#$%#")
		id := only_invoice[0]
		datetime := only_invoice[1]
		id_costumer := only_invoice[2]
		total := 0.0

		db := db.StorageDB
		query := "INSERT INTO Invoice(id, `datetime`, idCustomer, total) VALUES (?,?,?,?)"
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(id, datetime, id_costumer, total)

		if err != nil {
			return errors.New("No se pudo guardar elemento en BD.")
		}
	}
	return nil
}

func (r *repository_invoice) StoreInvoice(invoice models.Invoice) (models.Invoice, error) {
	db := db.StorageDB
	query := "INSERT INTO Invoice(`datetime`, idCustomer, total) VALUES (?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(invoice.Datetime, invoice.IdCustomer, invoice.Total)
	if err != nil {
		return models.Invoice{}, err
	}

	idCreado, _ := result.LastInsertId()
	invoice.Id = int(idCreado)
	return invoice, nil
}

func (r *repository_invoice) UpdateInvoice(invoice models.Invoice) (models.Invoice, error) {
	db := db.StorageDB
	query := "UPDATE Invoice SET `datetime` = ?, idCustomer = ?, total = ? WHERE id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(invoice.Datetime, invoice.IdCustomer, invoice.Total, invoice.Id)
	if err != nil {
		return models.Invoice{}, err
	}
	updatedRows, _ := result.RowsAffected()
	if updatedRows == 0 {
		return models.Invoice{}, errors.New("No se encontro al invoice.")
	}
	return invoice, nil
}

func (r *repository_invoice) UpdateTotalsOfInvoices() error {
	var inv []models.Invoice
	db := db.StorageDB
	rows, err := db.Query("SELECT id, datetime, idCustomer, total FROM Invoice")
	if err != nil {
		log.Fatal(err)
		return err
	}

	for rows.Next() {
		var i models.Invoice
		err := rows.Scan(&i.Id, &i.Datetime, &i.IdCustomer, &i.Total)
		if err != nil {
			log.Fatal(err)
			return err
		}
		inv = append(inv, i)
	}

	for _, inv := range inv {
		total := 0.00
		// Busco la cantidad en la tabla de store para esa factura.
		var sales []models.Sale
		rows, err := db.Query("SELECT id, idProduct, idInvoice, quantity FROM Sale WHERE idInvoice = ?", inv.Id)
		if err != nil {
			log.Fatal(err)
			return err
		}

		// Se recorre el resultado de la query.
		for rows.Next() {
			var sale models.Sale
			err := rows.Scan(&sale.Id, &sale.IdProduct, &sale.IdInvoice, &sale.Quantity)
			if err != nil {
				log.Fatal(err)
				return err
			}
			sales = append(sales, sale)
		}

		for _, s := range sales {
			var pr models.Product
			rows, err := db.Query("SELECT id, description, price FROM Product WHERE id = ?", s.IdProduct)
			if err != nil {
				log.Fatal(err)
			}

			// Se recorre el resultado de la query.
			for rows.Next() {
				err := rows.Scan(&pr.Id, &pr.Description, &pr.Price)
				if err != nil {
					log.Fatal(err)
				}
			}
			total = total + (float64(s.Quantity) * pr.Price)
		}

		// Ahora hago el update con la suma de los productos de las cantidades y los precios de los items de la factura.
		stmt, err := db.Prepare("UPDATE Invoice SET total = ? WHERE id = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		result, err := stmt.Exec(total, inv.Id)
		if err != nil {
			return err
		}
		updatedRows, _ := result.RowsAffected()
		if updatedRows == 0 {
			return errors.New("No se pudo modificar el total del registro.")
		}
	}
	return nil
}
