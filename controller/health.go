package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health heart beats
func Health(c *gin.Context) {
	c.String(http.StatusOK, "ok")
	return
}
