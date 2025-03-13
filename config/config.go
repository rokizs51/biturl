package config

type RateLimiterConfig struct {
	Enabled    bool
	Tokens     int
	RefillRate float64
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type Config struct {
	DatabaseDSN       string
	RedisConfig       RedisConfig
	RateLimiterConfig RateLimiterConfig
}

func NewConfig() *Config {
	return &Config{
		DatabaseDSN: "mongodb://localhost:27017",
		RedisConfig: RedisConfig{
			Host:     "localhost",
			Port:     6379,
			Password: "",
			DB:       0,
		},
		RateLimiterConfig: RateLimiterConfig{
			Enabled:    true,
			Tokens:     20,
			RefillRate: 1,
		},
	}
}
