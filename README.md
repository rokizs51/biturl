# Bit URL - URL Shortener Service

Bit URL is a simple and efficient URL shortening service built with Go. It allows users to convert long URLs into shorter, more manageable links while providing features like link expiration, click tracking, and rate limiting.

## Features

- URL Shortening: Convert long URLs into short, hash-based URLs
- Link Redirection: Automatic redirection from short URLs to original URLs
- Link Expiration: URLs automatically expire after one year
- Click Tracking: Track the number of times a shortened URL is accessed
- Rate Limiting: Protect the API from abuse using token bucket algorithm
- RESTful API: Simple and easy-to-use API endpoints
- Modern Web Frontend: React-based UI with TypeScript and Tailwind CSS

## Tech Stack

### Backend
- **Go**: Backend programming language
- **Gin**: Web framework
- **MongoDB**: Database for storing URL mappings
- **Redis**: For rate limiting and caching
- **Crypto**: For secure hash generation

### Frontend
- **React**: UI library
- **TypeScript**: Type-safe JavaScript
- **Tailwind CSS**: Utility-first CSS framework
- **Vite**: Build tool and development server

## Prerequisites

- Go 1.x or higher
- MongoDB
- Redis
- Node.js and npm
- Git

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener
   ```

2. Install backend dependencies:
   ```bash
   go mod download
   ```

3. Install frontend dependencies:
   ```bash
   cd web
   npm install
   ```

4. Configure MongoDB connection in `internal/database/db.go`

5. Configure Redis connection in `internal/database/redis.go`

6. Configure rate limiting parameters in your configuration file:
   ```go
   RateLimiter: {
     Tokens: 100,        // Maximum tokens per bucket
     RefillRate: 1,      // Tokens added per minute
   }
   ```

7. Run the backend:
   ```bash
   go run cmd/main.go
   ```

8. Run the frontend development server:
   ```bash
   cd web
   npm run dev
   ```

The backend server will start on `http://localhost:8080`
The frontend development server will start on `http://localhost:5173`

## API Documentation

### Shorten URL

- **Endpoint**: `POST /api/shorten`
- **Content-Type**: `application/json`
- **Request Body**:
  ```json
  {
    "long_url": "https://example.com/very/long/url"
  }
  ```
- **Response**:
  ```json
  {
    "short_url": "abc123",
    "long_url": "https://example.com/very/long/url",
    "expires": "2024-02-14T15:30:00Z"
  }
  ```

### Redirect to Original URL

- **Endpoint**: `GET /:shortURL`
- **Response**: Redirects to the original URL
- **Error Response** (if URL not found):
  ```json
  {
    "error": "URL not found"
  }
  ```

## Project Structure

```
├── cmd/
│   └── main.go           # Application entry point
├── config/
│   └── config.go         # Application configuration
├── internal/
│   ├── database/         # Database and Redis connections
│   ├── handler/          # HTTP request handlers
│   ├── middleware/       # HTTP middlewares (rate limiting, etc.)
│   ├── models/           # Data models
│   ├── repository/       # Database operations
│   ├── service/          # Business logic
│   └── utils/            # Utility functions
├── web/                  # Frontend application
│   ├── src/              # React source code
│   ├── public/           # Static assets
│   └── package.json      # Frontend dependencies
└── README.md
```

## Rate Limiting

The API implements rate limiting using a token bucket algorithm:

- Each client gets a bucket with a maximum number of tokens
- Tokens are refilled at a constant rate
- Each API request consumes one token
- When a bucket is empty, requests are rejected until tokens are refilled

## Error Handling

The API returns appropriate HTTP status codes:

- `200 OK`: Successful operation
- `400 Bad Request`: Invalid input
- `404 Not Found`: URL not found
- `429 Too Many Requests`: Rate limit exceeded
- `500 Internal Server Error`: Server-side errors

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.