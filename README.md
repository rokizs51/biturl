# Bit URL - URL Shortener Service

Bit URL is a simple and efficient URL shortening service built with Go. It allows users to convert long URLs into shorter, more manageable links while providing features like link expiration and click tracking.

## Features

- URL Shortening: Convert long URLs into short, hash-based URLs
- Link Redirection: Automatic redirection from short URLs to original URLs
- Link Expiration: URLs automatically expire after one year
- Click Tracking: Track the number of times a shortened URL is accessed
- RESTful API: Simple and easy-to-use API endpoints

## Tech Stack

- **Go**: Backend programming language
- **Gin**: Web framework
- **MongoDB**: Database for storing URL mappings
- **Crypto**: For secure hash generation

## Prerequisites

- Go 1.x or higher
- MongoDB
- Git

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Configure MongoDB connection in `internal/database/db.go`

4. Run the application:
   ```bash
   go run cmd/main.go
   ```

The server will start on `http://localhost:8080`

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
├── internal/
│   ├── database/         # Database connection and configuration
│   ├── handler/          # HTTP request handlers
│   ├── models/           # Data models
│   ├── repository/       # Database operations
│   ├── service/          # Business logic
│   └── utils/            # Utility functions
└── README.md
```

## Error Handling

The API returns appropriate HTTP status codes:

- `200 OK`: Successful operation
- `400 Bad Request`: Invalid input
- `404 Not Found`: URL not found
- `500 Internal Server Error`: Server-side errors

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.