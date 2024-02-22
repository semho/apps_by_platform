package models

type Repository interface {
	GetProduct(id int) Product
	GetProducts() []Product
	GetProductPage(page, pageSize int) (products []Product, totalAvailable int)
	GetProductPageCategory(categoryId, page, pageSize int) (products []Product, totalAvailable int)
	GetCategories() []Category
	Seed()
}
