# 🎬 Movie Library API

A RESTful API for managing a movie library, built with **Go (Golang)**.

The project follows a layered architecture and provides CRUD operations for managing movies, including validation and structured error handling.

## 🚀 Features

- Create, retrieve, update, and delete movies
- Retrieve a movie by ID or get all movies
- Input validation
- Structured JSON responses and error handling
- Layered architecture
- Dependency Injection
- Thread-safe in-memory storage

## 🛠 Technologies

- **Language:** Go
- **HTTP Router:** Gorilla Mux
- **Data Format:** JSON
- **Architecture:** Layered Architecture
- **Storage:** In-memory Map
- **Version Control:** Git

## 📁 Project Structure

```text
movie-library-api/
│
├── internal/
│   ├── handler/
│   │   └── dto/
│   ├── model/
│   ├── repository/
│   └── service/
│
├── main.go
├── go.mod
├── go.sum
├── .gitignore
└── README.md
```

## Architecture

The application uses a layered architecture to separate responsibilities:

```text
Client
  │
  ▼
Handler
  │
  ▼
Service
  │
  ▼
Repository
  │
  ▼
In-Memory Storage
```

- **Handler:** Handles HTTP requests and responses.
- **Service:** Contains business logic and validation.
- **Repository:** Manages movie data.
- **Model:** Defines the Movie entity.
- **DTO:** Defines request and error response structures.

## API Endpoints

| Method | Endpoint       | Description       |
| ------ | -------------- | ----------------- |
| POST   | `/movies`      | Create a movie    |
| GET    | `/movies`      | Get all movies    |
| GET    | `/movies/{id}` | Get a movie by ID |
| PUT    | `/movies/{id}` | Update a movie    |
| DELETE | `/movies/{id}` | Delete a movie    |

## Movie Model

```json
{
	"id": 1,
	"title": "Interstellar",
	"year": 2014,
	"director": "Christopher Nolan",
	"rating": 8.7,
	"description": "A team of explorers travels through a wormhole."
}
```

## ⚙️ Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/adikssl/movie-library-api.git
cd movie-library-api
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the application

```bash
go run .
```

The API will be available at:

```text
http://localhost:8080
```

## Testing

You can test the API using **Postman** or any HTTP client.

The API currently uses in-memory storage, so movie data is not persisted after the application stops.

## Future Improvements

- PostgreSQL integration
- Unit and integration tests
- Database migrations
- Pagination and filtering
- Improved API documentation
