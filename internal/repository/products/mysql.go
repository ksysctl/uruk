package products

import (
	"github.com/jinzhu/gorm"

	"github.com/ksysctl/uruk/internal/models"
)

type mysqlProductRepository struct {
	db *gorm.DB
}

// GetByID returns product by ID
func (r mysqlProductRepository) GetByID(id uint) models.Product {
	var product models.Product

	r.db.Where(
		&models.Product{ID: uint(id)},
	).First(&product)

	return product
}

// GetByID returns all products
func (r mysqlProductRepository) Get() []models.Product {
	var products []models.Product

	r.db.Find(&products)

	return products
}
