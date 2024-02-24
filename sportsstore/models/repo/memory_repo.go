package repo

import (
	"math"
	"sportsstore/models"
)

type MemoryRepo struct {
	products   []models.Product
	categories []models.Category
}

//func RegisterMemoryRepoService() {
//	services.AddSingleton(func() models.Repository {
//		repo := &MemoryRepo{}
//		repo.Seed()
//		return repo
//	})
//}

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

func (repo *MemoryRepo) GetProductPage(page, pageSize int) ([]models.Product, int) {
	return getPage(repo.products, page, pageSize), len(repo.products)
}

func getPage(src []models.Product, page, pageSize int) []models.Product {
	start := (page - 1) * pageSize
	if page > 0 && len(src) > start {
		end := (int)(math.Min((float64)(len(src)), (float64)(start+pageSize)))
		return src[start:end]
	}
	return []models.Product{}
}

func (repo *MemoryRepo) GetProductPageCategory(category, page, pageSize int) ([]models.Product, int) {
	if category == 0 {
		return repo.GetProductPage(page, pageSize)
	} else {
		filteredProducts := make([]models.Product, 0, len(repo.products))
		for _, p := range repo.products {
			if p.Category.ID == category {
				filteredProducts = append(filteredProducts, p)
			}
		}
		return getPage(filteredProducts, page, pageSize), len(filteredProducts)
	}
}
