package models

import "time"

type UrlMapping struct {
	ShortURL       string    `bson:"short_url"`
	LongURL        string    `bson:"long_url"`
	CreationDate   time.Time `bson:"creation_date"`
	ExpirationDate time.Time `bson:"expiration_date"`
	ClickCount     int       `bson:"click_count"`
}

type UrlMappingRequest struct {
	LongURL    string `json:"long_url" binding:"required"`
	CustomSlug string `json:"custom_slug"`
}
