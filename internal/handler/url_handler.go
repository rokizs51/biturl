package handler

import (
	"fmt"
	"net/http"
	"url-shortnere/internal/models"
	"url-shortnere/internal/service"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	service service.URLService
}

func NewURLHandler(urlService service.URLService) *URLHandler {
	return &URLHandler{
		service: urlService,
	}
}

func (h *URLHandler) ShortenUrl(c *gin.Context) {
	var request models.UrlMappingRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request format",
		})
		return
	}

	result, err := h.service.ShortenUrl(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_url": result.ShortURL,
		"long_url":  result.LongURL,
		"expires":   result.ExpirationDate,
	})

}

func (h *URLHandler) RedirectUrl(c *gin.Context) {
	shortUrl := c.Param("shortURL")
	if shortUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Short URL is required",
		})
		return
	}

	// find url map
	mapping, err := h.service.GetOriginalURL(c.Request.Context(), shortUrl)
	fmt.Println(mapping, err)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "URL not found",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, mapping.LongURL)
}
