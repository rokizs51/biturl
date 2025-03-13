package handler

import (
	"net/http"
	"url-shortnere/config"
	"url-shortnere/internal/service"

	"github.com/gin-gonic/gin"
)

func ResetToken(c *gin.Context) {
	service := service.NewRateLimiterService(&config.NewConfig().RateLimiterConfig)
	identifier := c.ClientIP()
	result := service.ResetToken(c, identifier)

	if !result {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed resetting token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success resetting token"})
}
