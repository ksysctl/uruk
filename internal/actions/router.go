package actions

import (
	"github.com/ksysctl/uruk/internal/handlers"
)

// Define routes
func (r Engine) route() {
	r.router.GET("/products", handlers.GetProducts)
	r.router.GET("/products/:id", handlers.GetProduct)

	r.router.GET("/", handlers.GetHome)
}
