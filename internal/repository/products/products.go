package products

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"fmt"
	"github.com/ksysctl/uruk/internal/models"
)

// ProductRepository is the product repository
type ProductRepository interface {
	GetByID(id uint) models.Product
	Get() []models.Product
}

// New returns the product repository
func New(env string, db *gorm.DB) ProductRepository {
	if env == gin.ReleaseMode {
		fmt.Println("Using MySQL repository")
		return mysqlProductRepository{
			db: db,
		}
	}

	fmt.Println("Using Mock repository")
	return mockProductRepository{}
}
