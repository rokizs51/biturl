package repository

import (
	"context"
	"fmt"
	"url-shortnere/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

// new url repo
func NewUrlRepository(db *mongo.Database) *URLRepository {
	return &URLRepository{
		db:         db,
		collection: db.Collection("url_mappings"),
	}
}

// insert record to db
func (r *URLRepository) Insert(ctx context.Context, mapping models.UrlMapping) error {
	_, err := r.collection.InsertOne(ctx, mapping)
	return err
}

// find by short url
func (r *URLRepository) FindByShortURL(ctx context.Context, shortUrl string) (*models.UrlMapping, error) {
	var mapping models.UrlMapping
	err := r.collection.FindOne(ctx, bson.M{"short_url": shortUrl}).Decode(&mapping)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return &mapping, nil
}

// update clickcount
func (r *URLRepository) UpdateClickCount(ctx context.Context, shortUrl string) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"short_url": shortUrl},
		bson.M{"$inc": bson.M{"click_count": 1}},
	)
	return err
}
