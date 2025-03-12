package main

import (
	"context"
	"log"
	"url-shortnere/internal/handler"
	"url-shortnere/internal/repository"
	"url-shortnere/internal/service"

	"url-shortnere/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	db, err := database.ConnectDB(context.Background())
	if err != nil {
		log.Fatalf("Failed connect to database: %v", err)
	}

	router := gin.Default()
	urlRepo := repository.NewUrlRepository(db)
	urlService := service.NewURLService(*urlRepo)
	urlHandler := handler.NewURLHandler(*urlService)

	router.POST("/api/shorten", urlHandler.ShortenUrl)
	router.GET("/:shortURL", urlHandler.RedirectUrl)
	log.Println("RUNING....")
	router.Run(":8080")
}
