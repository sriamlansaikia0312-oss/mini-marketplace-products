package models

// Product is a NetApp-style marketplace listing (in-memory demo).
type Product struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Cloud  string `json:"cloud"` // aws | azure | gcp
	SKU    string `json:"sku"`
	Family string `json:"family,omitempty"`
}
