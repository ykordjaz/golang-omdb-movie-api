# 🎬 Go Movie Info API

A simple backend Go web server that lets you fetch movie data from the OMDB API using a title (and optional year) via a REST endpoint.

---

## ✨ Features

- ✅ Returns movie title, year, and poster URL from the OMDB API  
- ✅ Supports movie title query via `?title=...`
- ✅ Optional `year` query param to narrow the search
- ✅ /movie: Uses t=title to fetch one exact match
- ✅ /search: Uses s=keyword to find multiple possible matches (e.g., all movies with “batman”)
- ✅ Reads your API key securely from a `.env` file  
- ✅ Written with clean Go code using `net/http` and `encoding/json`

---

## 📦 Tech Stack

- Language: Go (Golang)
- HTTP server: `net/http`
- API client: Native `http.Get`
- Environment variables: [`github.com/joho/godotenv`](https://github.com/joho/godotenv)
- External API: [OMDB API](https://www.omdbapi.com/)

---

## 🚀 Getting Started

### 1. Clone the project

```bash
git clone https://github.com/your-username/go-movie-api.git
cd go-movie-api
