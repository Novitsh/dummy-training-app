// handlers/about.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}
