package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHome shows homepage
func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index", gin.H{})
}
