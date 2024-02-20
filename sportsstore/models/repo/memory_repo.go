package repo

import (
	"platform/services"
	"sportsstore/models"
)

type MemoryRepo struct {
	products   []models.Product
	categories []models.Category
}

func RegisterMemoryRepoService() {
	services.AddSingleton(func() models.Repository {
		repo := &MemoryRepo{}
		repo.Seed()
		return repo
	})
}

func (repo *MemoryRepo) GetProduct(id int) models.Product {
	for _, p := range repo.products {
		if p.ID == id {
			return p
		}
	}
	return models.Product{}
}

func (repo *MemoryRepo) GetProducts() []models.Product {
	return repo.products
}

func (repo *MemoryRepo) GetCategories() []models.Category {
	return repo.categories
}
