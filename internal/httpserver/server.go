package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/netapp/mini-marketplace-products/internal/data"
	"github.com/netapp/mini-marketplace-products/internal/web"
)

// Options configures the server handler.
type Options struct {
	Quiet bool // no per-request logs (tests)
}

// New builds the default handler with access logs.
func New() http.Handler {
	return NewWithOptions(Options{})
}

// NewWithOptions builds the HTTP handler.
func NewWithOptions(o Options) http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/api/health", serveHealth)
	m.HandleFunc("/api/products/", serveProductByID)
	m.HandleFunc("/api/products", serveProductList)
	m.HandleFunc("/", serveIndex)
	if o.Quiet {
		return m
	}
	return withAccessLog(m)
}

func withAccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t0 := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(t0).Round(time.Millisecond))
	})
}

func serveHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeMethodNotAllowed(w)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeMethodNotAllowed(w)
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(web.IndexHTML)
}

func serveProductList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeMethodNotAllowed(w)
		return
	}
	if r.URL.Path != "/api/products" {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, http.StatusOK, data.List())
}

func serveProductByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeMethodNotAllowed(w)
		return
	}
	id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/api/products/"), "/")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	p, ok := data.Get(id)
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	writeJSON(w, http.StatusOK, p)
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		log.Printf("encode json: %v", err)
	}
}

func writeMethodNotAllowed(w http.ResponseWriter) {
	w.Header().Set("Allow", http.MethodGet)
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
