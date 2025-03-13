package middleware

import (
	"log"
	"net/http"
	"strconv"
	"url-shortnere/config"
	"url-shortnere/internal/service"

	"github.com/gin-gonic/gin"
)

func RateLimiter(cfg *config.Config) gin.HandlerFunc {
	rateLimiter := service.NewRateLimiterService(&cfg.RateLimiterConfig)
	return func(ctx *gin.Context) {
		identifier := ctx.ClientIP()
		allowed, count, err := rateLimiter.IsAllowed(ctx, identifier)
		if err != nil {
			log.Printf("Rate limiter error: %v, falling back to allow traffic", err)
			ctx.Next()
			return
		}

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Rate limiting error",
				"details": err.Error(),
			})
			return
		}

		ctx.Header("X-RateLimit-Limit", strconv.Itoa(cfg.RateLimiterConfig.Tokens))
		ctx.Header("X-RateLimit-Remaining", strconv.Itoa(count))
		ctx.Header("X-Rate-Limit-Type", "token-bucket")

		if !allowed {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			return
		}
		ctx.Next()
	}
}
