package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingHandler godoc
// @Summary Ping API
// @Description Check if the API is live
// @Tags ping
// @Success 200 {string} string "pong"
// @Router /ping [get]
func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "okay",
	})
}
