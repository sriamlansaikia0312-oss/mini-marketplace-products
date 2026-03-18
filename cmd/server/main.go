package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/netapp/mini-marketplace-products/internal/httpserver"
)

func main() {
	addr := address()
	h := httpserver.New()
	log.Println("mini-marketplace-products")
	log.Printf("  http://127.0.0.1%s/\n", addr)
	log.Fatal(http.ListenAndServe(addr, h))
}

func address() string {
	p := strings.TrimSpace(os.Getenv("PORT"))
	switch {
	case p == "":
		return ":8080"
	case strings.HasPrefix(p, ":"):
		return p
	default:
		return ":" + p
	}
}
