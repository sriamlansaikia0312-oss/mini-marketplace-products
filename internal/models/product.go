package models

const (
	CloudAWS   = "aws"
	CloudAzure = "azure"
	CloudGCP   = "gcp"
)

// Product represents one marketplace listing (demo data).
type Product struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Cloud  string `json:"cloud"`
	SKU    string `json:"sku"`
	Family string `json:"family,omitempty"`
}
