package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	startTime time.Time
}

func NewHealthHandler() *HealthHandler {

	return &HealthHandler{
		startTime: time.Now(),
	}
}

func (h *HealthHandler) Live(c *gin.Context) {

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "alive",
		},
	)
}

func (h *HealthHandler) Ready(c *gin.Context) {

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "ready",
		},
	)
}

func (h *HealthHandler) Health(c *gin.Context) {

	c.JSON(
		http.StatusOK,
		gin.H{
			"application": "AutoParts Pro API",
			"status":      "healthy",
			"uptime":      time.Since(h.startTime).String(),
			"time":        time.Now().UTC(),
		},
	)
}
