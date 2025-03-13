package main

import (
	"context"
	"log"
	"url-shortnere/config"
	"url-shortnere/internal/handler"
	"url-shortnere/internal/middleware"
	"url-shortnere/internal/repository"
	"url-shortnere/internal/service"

	"url-shortnere/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	db, err := database.ConnectDB(context.Background(), cfg)
	database.InitializeRedis(cfg)

	if err != nil {
		log.Fatalf("Failed connect to database: %v", err)
	}

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// router.Use(middleware.RateLimiter(cfg))

	urlRepo := repository.NewUrlRepository(db)
	urlService := service.NewURLService(*urlRepo)
	urlHandler := handler.NewURLHandler(*urlService)

	router.GET("/ping", middleware.RateLimiter(cfg), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/reset-limiter", handler.ResetToken)
	router.POST("/api/shorten", middleware.RateLimiter(cfg), urlHandler.ShortenUrl)
	router.GET("/:shortURL", urlHandler.RedirectUrl)
	log.Println("RUNING....")
	router.Run(":8080")
}
