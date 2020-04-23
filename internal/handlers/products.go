package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ksysctl/uruk/internal/handlers/response"
	"github.com/ksysctl/uruk/internal/models"
	"github.com/ksysctl/uruk/internal/repository"
)

// GetProducts shows all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	var view = "products/products"
	var res = response.NewHandle(gin.H{
		"title":   "Catalog View",
		"product": products,
	})

	// Access to product repository from context
	repo := c.MustGet("repos").(repository.Repository)
	products = repo.Product.Get()

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

	// Access to product repository from context
	repo := c.MustGet("repos").(repository.Repository)
	product = repo.Product.GetByID(uint(productID))

	// Check if entity exists
	if product.ID == 0 {
		res.SetPayload("error", fmt.Sprintf("Product ID does not exist %s", id))
		res.GetNotFoundResponse(c, view)
		return
	}

	res.SetPayload("product", product)
	res.GetOkResponse(c, view)
}
