package models

type Repository interface {
	GetProduct(id int) Product
	GetProducts() []Product
	GetProductPage(page, pageSize int) (product []Product, totalAvailable int)
	GetCategories() []Category
	Seed()
}
