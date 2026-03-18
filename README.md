# mini-marketplace-products

Tiny demo: in-memory NetApp-style marketplace products and **two read APIs** (Go). Use Node only to run via npm.

## Run

```bash
npm start
# or: go run ./cmd/server
```

Server: `http://localhost:8080` (override with `PORT`).

**In the browser:** open **http://localhost:8080/** — you’ll see a table of products; click a row for detail (loaded from the API).

## APIs

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/products` | All products (JSON array) |
| `GET` | `/api/products/{id}` | One product, e.g. `np-aws-console` |

## Layout

```
cmd/server/main.go    # HTTP server
internal/data/        # in-memory product list
internal/models/      # JSON shape
internal/web/         # UI (embedded HTML)
```

Requires **Go 1.22+** (path patterns). For CI see `.github/workflows/ci.yml`.
