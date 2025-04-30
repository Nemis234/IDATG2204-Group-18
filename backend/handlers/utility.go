package handlers

import (
	"database/sql"
)

func GetProduct(productRows *sql.Rows) (Product, error) {
	var p Product
	var c Category
	var b Brand
	if err := productRows.Scan(
		&p.ProductID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.StockQuantity,
		&c.ID,
		&b.ID); err != nil {
		return p, err
	}

	if err := db.QueryRow(queryCategories, c.ID).Scan(
		&c.Name,
		&c.Description); err != nil {
		return p, err
	}
	if err := db.QueryRow(queryBrands, b.ID).Scan(
		&b.Name,
		&b.Description); err != nil {
		return p, err
	}

	p.Category = c
	p.Brand = b
	return p, nil
}
