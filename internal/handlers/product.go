package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/ksysctl/uruk/internal/handlers/response"
	"github.com/ksysctl/uruk/internal/models"
)

// GetProducts returns all products
func GetProducts(c *gin.Context) {
	var products []models.Product

	contentType := response.SetType(
		c.DefaultQuery("type", response.TypeHTML),
	)

	db := c.MustGet("db").(*gorm.DB)
	db.Find(&products)

	if contentType == response.TypeHTML {
		c.HTML(http.StatusOK, "products/products", gin.H{
			"title":    "Catalog Products",
			"products": products,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": products,
		})
	}
}

// GetProduct shows single product
func GetProduct(c *gin.Context) {
	var product []models.Product

	contentType := response.SetType(
		c.DefaultQuery("type", response.TypeHTML),
	)

	id := c.Param("id")
	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error parsing string Uint %s", id))
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Where(
		&models.Product{ID: uint(productID)},
	).First(&product)

	if contentType == response.TypeHTML {
		c.HTML(http.StatusOK, "products/products", gin.H{
			"title":    "Product View",
			"products": product,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": product,
		})
	}
}
