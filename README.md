# Movies CRUD API

A RESTful API for managing movies, built with Go, Gin, GORM, and PostgreSQL.

## Features

- CRUD operations for movies
- JWT-based authentication
- PostgreSQL database
- Docker support
- Dependency injection with UberFx

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose
- PostgreSQL (if running locally)

## Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd task_itv
```

2. Build and run with Docker Compose:
```bash
docker-compose up --build
```

The API will be available at `http://localhost:8080`

## API Endpoints

### Authentication
- `POST /login` - Get JWT token (currently returns dummy token)

### Movies
- `POST /api/movies` - Create a new movie
- `GET /api/movies` - Get all movies
- `GET /api/movies/:id` - Get a specific movie
- `PUT /api/movies/:id` - Update a movie
- `DELETE /api/movies/:id` - Delete a movie

## Environment Variables

Create a `.env` file with the following variables:
```
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=3210
DB_NAME=movies_db
JWT_SECRET=your-secret-key
PORT=8080
```

## Development

1. Install dependencies:
```bash
go mod download
```

2. Run the application:
```bash
go run main.go
```

## Testing

To test the API, you can use tools like curl or Postman. Here are some example requests:

Create a movie:
```bash
curl -X POST http://localhost:8080/api/movies \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-token" \
  -d '{"title":"Inception","director":"Christopher Nolan","year":2010,"plot":"A thief who steals corporate secrets..."}'
```

Get all movies:
```bash
curl -X GET http://localhost:8080/api/movies \
  -H "Authorization: Bearer your-token"
```

## License

MIT 