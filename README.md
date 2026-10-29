# 🎬 Movie Library REST API

A REST API for managing a personal movie library, built with Go.

## Tech Stack

- Go
- net/http
- JSON
- In-memory storage (PostgreSQL planned)

## Features

- [x] HTTP server with health check
- [ ] Movie CRUD operations
- [ ] Movie genres
- [ ] Watchlist management
- [ ] Search and filtering
- [ ] PostgreSQL integration

## Getting Started

**1. Clone the repository**

```bash
git clone <https://github.com/adikssl/movie-library-api>
cd movie-library-api
```

**2. Run the application**

```bash
go run ./cmd/api
```

The server runs at `http://localhost:8080`.

## API Endpoints

| Method | Endpoint | Description      |
| ------ | -------- | ---------------- |
| GET    | /health  | Check API status |

## Project Status

🚧 Under development.

This project is being developed as part of my Go backend learning journey.
