package actions

import (
	"github.com/gin-gonic/gin"

	"github.com/ksysctl/uruk/internal/handlers"
)

// Engine is the framework's instance
type Engine struct {
	router *gin.Engine
}

// Init initializes engine
func (r Engine) Init() *gin.Engine {
	r.router = gin.Default()

	r.render()
	r.logger()
	r.storage()

	r.router.GET("/products", handlers.GetProducts)
	r.router.GET("/products/:id", handlers.GetProduct)

	r.router.GET("/", handlers.GetHome)

	return r.router
}
