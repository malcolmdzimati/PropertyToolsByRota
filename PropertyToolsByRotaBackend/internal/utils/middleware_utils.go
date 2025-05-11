package middleware_utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GinMiddleware adapts a standard http.Handler middleware to Gin's middleware signature
func GinMiddleware(mw func(http.Handler) http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Update gin's context with potentially modified request
			c.Request = r
			c.Next()
		})

		// Run the middleware chain
		mw(finalHandler).ServeHTTP(c.Writer, c.Request)

		// Prevent Gin from continuing the chain again
		c.Abort()
	}
}
