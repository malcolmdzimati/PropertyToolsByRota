package router

import (
	"PropertyToolsByRotaBackend/internal/middleware"
	middleware_utils "PropertyToolsByRotaBackend/internal/utils"
	"net/http"

	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) *gin.Engine {

	// === Group: /v1/api ===
	api := r.Group("/v1/api")

	// Setup the public routes
	SetupPublicRoutes(api)

	// === Group: /v1/api/secure ===
	secure := api.Group("/secure")
	secure.Use(middleware_utils.GinMiddleware(auth_middleware.EnsureValidToken()))

	// Setup the get routes
	SetupGetRoutes(secure)

	// Setup the post routes
	SetupPostRoutes(secure)

	return r
}

func SetupPublicRoutes(group *gin.RouterGroup) {
	// GET /ping (public)
	group.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "okay"})
	})

	// POST /ping (public)
	group.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}

func SetupGetRoutes(group *gin.RouterGroup) {
	// GET /ping (public)
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}

func SetupPostRoutes(group *gin.RouterGroup) {
	// POST /ping (public)
	group.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}

func scopedPostRoutes(group *gin.RouterGroup) {
	group.POST("", func(c *gin.Context) {
		token := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		claims := token.CustomClaims.(*auth_middleware.CustomClaims)

		if !claims.HasScope("write:messages") {
			c.JSON(http.StatusForbidden, gin.H{"message": "Insufficient scope"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Secure POST received with proper scope"})
	})
}
