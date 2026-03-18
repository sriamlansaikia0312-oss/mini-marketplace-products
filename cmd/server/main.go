package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/netapp/mini-marketplace-products/internal/data"
	"github.com/netapp/mini-marketplace-products/internal/models"
	"github.com/netapp/mini-marketplace-products/internal/web"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/products", handleListProducts)
	mux.HandleFunc("GET /api/products/{id}", handleGetProduct)
	mux.HandleFunc("GET /", handleIndex)

	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}
	log.Printf("open in browser: http://localhost%s/", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(web.IndexHTML)
}

func handleListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data.Products)
}

func handleGetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var found *models.Product
	for i := range data.Products {
		if data.Products[i].ID == id {
			found = &data.Products[i]
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if found == nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
		return
	}
	_ = json.NewEncoder(w).Encode(found)
}
