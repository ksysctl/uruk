package actions

import (
	"github.com/gin-gonic/gin"
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
	r.route()

	return r.router
}
