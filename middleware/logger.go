// middleware/logger.go
package middleware

import (
	"github.com/gin-gonic/gin"
)

// Logger returns a named request-logging middleware using Gin's built-in logger.
func Logger() gin.HandlerFunc {
	return gin.Logger()
}
