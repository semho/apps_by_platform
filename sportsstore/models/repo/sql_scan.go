package repo

import (
	"database/sql"
	"sportsstore/models"
)

func scanProducts(rows *sql.Rows) ([]models.Product, error) {
	products := make([]models.Product, 0, 10)
	for rows.Next() {
		p := models.Product{Category: &models.Category{}}
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category.ID, &p.Category.CategoryName)
		if err != nil {
			return products, err
		}
		products = append(products, p)
	}
	return products, nil
}

func scanProduct(row *sql.Row) (models.Product, error) {
	p := models.Product{Category: &models.Category{}}
	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category.ID, &p.Category.CategoryName)
	return p, err
}
