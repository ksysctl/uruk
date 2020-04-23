package products

import (
	"github.com/ksysctl/uruk/internal/models"
)

type mockProductRepository struct {
	// @TODO: ...mock.Mock
}

// GetByID returns product by ID
func (r mockProductRepository) GetByID(id uint) models.Product {
	return models.Product{
		// @TODO: Implement it
	}
}

// GetByID returns all products
func (r mockProductRepository) Get() []models.Product {
	return []models.Product{
		// @TODO: Implement it
	}
}
