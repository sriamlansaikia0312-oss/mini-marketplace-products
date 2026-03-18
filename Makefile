# Run all commands from the repository root.

.PHONY: help all build test run clean fmt vet

BIN     := bin/server
CMD_PKG := ./cmd/server
PORT    ?= 8080

help: ## List targets
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?##' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-10s %s\n", $$1, $$2}'
	@echo ""
	@echo "  make run          → http://127.0.0.1:8080/"
	@echo "  PORT=3000 make run"

all: test build ## Test, then build (default: make)

build: ## Write bin/server
	@mkdir -p bin
	go build -o $(BIN) $(CMD_PKG)
	@echo "Built $(BIN)"

test: ## go test ./...
	go test ./... -count=1

run: ## Run server (foreground)
	PORT=$(PORT) go run $(CMD_PKG)

clean: ## Remove bin/server
	rm -f $(BIN)
	rmdir bin 2>/dev/null || true

fmt: ## go fmt ./...
	go fmt ./...

vet: ## go vet ./...
	go vet ./...
