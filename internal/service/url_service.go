package service

import (
	"context"
	"errors"
	"log"
	"time"
	"url-shortnere/internal/models"
	"url-shortnere/internal/repository"
	"url-shortnere/internal/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type URLService struct {
	urlRepository *repository.URLRepository
}

// new url service
func NewURLService(repo repository.URLRepository) *URLService {
	return &URLService{
		urlRepository: &repo,
	}
}

func (s *URLService) ShortenUrl(ctx context.Context, urlRequest models.UrlMappingRequest) (*models.UrlMapping, error) {
	var shortUrl string

	if urlRequest.CustomSlug != "" {
		shortUrl = urlRequest.CustomSlug
	} else {
		shortUrl = utils.ShortenURLHash(urlRequest.LongURL)
	}

	exist, err := s.urlRepository.FindByShortURL(ctx, shortUrl)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}
	if exist != nil && urlRequest.LongURL != exist.LongURL {
		return nil, errors.New("custom slug already exists")
	}
	if exist != nil {
		return exist, nil
	}

	url := models.UrlMapping{
		ShortURL:       shortUrl,
		LongURL:        urlRequest.LongURL,
		CreationDate:   time.Now(),
		ExpirationDate: time.Now().AddDate(1, 0, 0),
		ClickCount:     0,
	}

	err = s.urlRepository.Insert(ctx, url)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (s *URLService) GetOriginalURL(ctx context.Context, shortUrl string) (*models.UrlMapping, error) {
	mapping, err := s.urlRepository.FindByShortURL(ctx, shortUrl)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("URL not found")
		}
		return nil, err
	}

	// update click count
	err = s.urlRepository.UpdateClickCount(ctx, shortUrl)
	if err != nil {
		log.Printf("Failed to update click count : %v", err)
	}
	return mapping, nil
}
