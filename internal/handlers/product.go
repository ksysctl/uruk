package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/ksysctl/uruk/internal/handlers/response"
	"github.com/ksysctl/uruk/internal/models"
)

// GetProducts shows all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	var view = "products/products"
	var res = response.NewHandle(gin.H{
		"title":   "Catalog View",
		"product": products,
	})

	// Get entities from database
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&products)

	res.SetPayload("products", products)
	res.GetOkResponse(c, view)
}

// GetProduct shows single product by ID
func GetProduct(c *gin.Context) {
	var product models.Product
	var view = "products/product"
	var res = response.NewHandle(gin.H{
		"title":   "Product View",
		"product": product,
	})

	// Obtain parameters
	id := c.Param("id")

	// Parse parameter
	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		res.SetPayload("error", fmt.Sprintf("Error parsing %s type(%T) as Uint", id, id))
		res.GetBadResponse(c, view)
		return
	}

	// Get entity from database
	db := c.MustGet("db").(*gorm.DB)
	db.Where(
		&models.Product{ID: uint(productID)},
	).First(&product)

	// Check if entity exists
	if product.ID == 0 {
		res.SetPayload("error", fmt.Sprintf("Product ID does not exist %s", id))
		res.GetNotFoundResponse(c, view)
		return
	}

	res.SetPayload("product", product)
	res.GetOkResponse(c, view)
}
