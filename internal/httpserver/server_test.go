package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func quiet() http.Handler {
	return NewWithOptions(Options{Quiet: true})
}

func TestHealth(t *testing.T) {
	ts := httptest.NewServer(quiet())
	t.Cleanup(ts.Close)

	res, err := http.Get(ts.URL + "/api/health")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status %d", res.StatusCode)
	}
}

func TestProductsList(t *testing.T) {
	ts := httptest.NewServer(quiet())
	t.Cleanup(ts.Close)

	res, err := http.Get(ts.URL + "/api/products")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status %d", res.StatusCode)
	}
}

func TestProductByID(t *testing.T) {
	ts := httptest.NewServer(quiet())
	t.Cleanup(ts.Close)

	res, err := http.Get(ts.URL + "/api/products/np-aws-console")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status %d", res.StatusCode)
	}

	res2, err := http.Get(ts.URL + "/api/products/unknown")
	if err != nil {
		t.Fatal(err)
	}
	defer res2.Body.Close()
	if res2.StatusCode != http.StatusNotFound {
		t.Fatalf("want 404, got %d", res2.StatusCode)
	}
}

func TestIndex(t *testing.T) {
	ts := httptest.NewServer(quiet())
	t.Cleanup(ts.Close)

	res, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status %d", res.StatusCode)
	}
}

func TestPostProductsNotAllowed(t *testing.T) {
	ts := httptest.NewServer(quiet())
	t.Cleanup(ts.Close)

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/products", nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("status %d", res.StatusCode)
	}
}
