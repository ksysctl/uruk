package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/ksysctl/uruk/internal/repository/products"
)

// Repository contains supported repositories
type Repository struct {
	Product products.ProductRepository
}

// New returns supported Repositories
func New(env string, db *gorm.DB) Repository {
	return Repository{
		products.New(env, db),
	}
}
