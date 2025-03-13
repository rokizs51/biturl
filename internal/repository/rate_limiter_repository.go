package repository

import (
	"context"
	"encoding/json"
	"time"
	"url-shortnere/internal/database"

	"github.com/redis/go-redis/v9"
)

type RateLimiterRepository struct {
	redis *redis.Client
}

type BucketState struct {
	Tokens     float64   `json:"tokens"`
	LastRefill time.Time `json:"last_refill"`
}

func NewRateLimiterRepository() *RateLimiterRepository {
	return &RateLimiterRepository{redis: database.GetRedis()}
}

func (r *RateLimiterRepository) GetBucket(ctx context.Context, identifier string) (*BucketState, error) {
	data, err := r.redis.Get(ctx, "bucket:"+identifier).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var bucketState BucketState
	err = json.Unmarshal([]byte(data), &bucketState)
	if err != nil {
		return nil, err
	}

	return &bucketState, nil
}

func (r *RateLimiterRepository) UpdateBucket(ctx context.Context, identifier string, bucketState *BucketState) error {
	data, err := json.Marshal(bucketState)
	if err != nil {
		return err
	}

	return r.redis.Set(ctx, "bucket:"+identifier, data, 0).Err()
}

func (r *RateLimiterRepository) ResetTokens(ctx context.Context, identifier string, bucketState *BucketState) error {
	data, err := json.Marshal(bucketState)
	if err != nil {
		return err
	}

	return r.redis.Set(ctx, "bucket:"+identifier, data, 0).Err()
}
