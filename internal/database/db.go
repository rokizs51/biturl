package database

import (
	"context"
	"fmt"
	"url-shortnere/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context) (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	fmt.Println("Success Connecting to DB")
	return client.Database("bitURL"), nil
}

func Insert(ctx context.Context, db *mongo.Database, mapping models.UrlMapping) error {
	collection := db.Collection("url_mappings")
	_, err := collection.InsertOne(ctx, mapping)
	return err
}
