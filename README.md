# mini-marketplace-products

Demo HTTP service: in-memory NetApp-style marketplace products, JSON APIs, and a small HTML UI.

## Requirements

- Go **1.21+**

## Commands

| Command | Action |
|---------|--------|
| `make run` | Start server at **http://127.0.0.1:8080/** |
| `PORT=3000 make run` | Other port |
| `make` | Run tests and build `bin/server` |
| `make test` | Tests only |

Also: `go run ./cmd/server`, `npm start`, `npm test`.

## HTTP

| Method | Path | Response |
|--------|------|----------|
| GET | `/` | HTML UI |
| GET | `/api/health` | `{"status":"ok"}` |
| GET | `/api/products` | JSON array of products |
| GET | `/api/products/{id}` | One product or 404 |

## Layout

```
cmd/server/main.go
internal/httpserver/   # HTTP router and handlers
internal/data/         # In-memory catalog
internal/models/       # Product model
internal/web/          # Embedded index.html
```
