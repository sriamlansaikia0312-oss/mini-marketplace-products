package data

import (
	"testing"

	"github.com/netapp/mini-marketplace-products/internal/models"
)

func TestListCount(t *testing.T) {
	if n := len(List()); n != 6 {
		t.Fatalf("want 6 products, got %d", n)
	}
}

func TestGet(t *testing.T) {
	p, ok := Get("np-aws-console")
	if !ok {
		t.Fatal("expected product")
	}
	if p.Cloud != models.CloudAWS || p.SKU != "NTAP-CONSOLE-AWS" {
		t.Fatalf("wrong product: %+v", p)
	}
	if _, ok := Get("missing"); ok {
		t.Fatal("expected not found")
	}
}
