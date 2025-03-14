package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"url-shortnere/internal/database"
	"url-shortnere/internal/models"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
	redis      *redis.Client
}

// new url repo
func NewUrlRepository(db *mongo.Database) *URLRepository {
	return &URLRepository{
		db:         db,
		collection: db.Collection("url_mappings"),
		redis:      database.GetRedis(),
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
	val, err := r.redis.Get(ctx, shortUrl).Result()
	if err == redis.Nil {
		err := r.collection.FindOne(ctx, bson.M{"short_url": shortUrl}).Decode(&mapping)
		fmt.Println(err)
		if err != nil {
			return nil, err
		}
		// struct to json
		jsonData, _ := json.Marshal(mapping)
		// cache in redis
		r.redis.Set(ctx, shortUrl, jsonData, 1*time.Hour)
	} else if err != nil {
		log.Fatal(err)
	} else {
		// json to struct
		err := json.Unmarshal([]byte(val), &mapping)
		if err != nil {
			log.Fatal("Error Decoding JSON", err)
		}
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
