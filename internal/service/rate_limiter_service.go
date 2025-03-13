package service

import (
	"context"
	"time"
	"url-shortnere/config"
	"url-shortnere/internal/repository"
)

type RateLimiterService struct {
	repo   *repository.RateLimiterRepository
	config *config.RateLimiterConfig
}

func NewRateLimiterService(cfg *config.RateLimiterConfig) *RateLimiterService {
	return &RateLimiterService{
		repo:   repository.NewRateLimiterRepository(),
		config: cfg,
	}
}

func (s *RateLimiterService) IsAllowed(ctx context.Context, identifier string) (bool, int, error) {
	bucket, err := s.repo.GetBucket(ctx, identifier)
	if err != nil {
		return false, 0, err
	}

	now := time.Now()
	if bucket == nil {
		bucket = &repository.BucketState{
			Tokens:     float64(s.config.Tokens),
			LastRefill: now,
		}
	}

	elapsed := now.Sub(bucket.LastRefill).Minutes()
	tokenToAdd := elapsed * s.config.RefillRate
	bucket.Tokens = min(float64(s.config.Tokens), bucket.Tokens+tokenToAdd)
	bucket.LastRefill = now

	if bucket.Tokens < 1 {
		return false, int(bucket.Tokens), nil
	}
	bucket.Tokens -= 1
	err = s.repo.UpdateBucket(ctx, identifier, bucket)
	if err != nil {
		return false, 0, err
	}

	return true, int(bucket.Tokens), nil
}

func (s *RateLimiterService) ResetToken(ctx context.Context, identifier string) bool {
	bucket, err := s.repo.GetBucket(ctx, identifier)
	if err != nil {
		return false
	}
	bucket.Tokens = float64(s.config.Tokens)

	err = s.repo.UpdateBucket(ctx, identifier, bucket)
	if err != nil {
		return false
	}

	return true
}
