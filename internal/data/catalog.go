package data

import (
	"sync"

	"github.com/netapp/mini-marketplace-products/internal/models"
)

var (
	mu   sync.RWMutex
	list []models.Product
	byID map[string]models.Product
)

func init() {
	seed := []models.Product{
		{ID: "np-aws-console", Name: "NetApp Console", Cloud: models.CloudAWS, SKU: "NTAP-CONSOLE-AWS", Family: "Console"},
		{ID: "np-azure-console", Name: "NetApp Console", Cloud: models.CloudAzure, SKU: "NTAP-CONSOLE-AZ", Family: "Console"},
		{ID: "np-gcp-console", Name: "NetApp Console", Cloud: models.CloudGCP, SKU: "NTAP-CONSOLE-GCP", Family: "Console"},
		{ID: "np-aws-dii", Name: "Data Infrastructure Insights", Cloud: models.CloudAWS, SKU: "NTAP-DII-AWS", Family: "DII"},
		{ID: "np-azure-dii", Name: "Data Infrastructure Insights", Cloud: models.CloudAzure, SKU: "NTAP-DII-AZ", Family: "DII"},
		{ID: "np-aws-ic", Name: "IntraCluster", Cloud: models.CloudAWS, SKU: "NTAP-IC-AWS", Family: "IntraCluster"},
	}
	byID = make(map[string]models.Product, len(seed))
	for _, p := range seed {
		byID[p.ID] = p
	}
	list = seed
}

// List returns a shallow copy of all products.
func List() []models.Product {
	mu.RLock()
	defer mu.RUnlock()
	out := make([]models.Product, len(list))
	copy(out, list)
	return out
}

// Get returns one product by id.
func Get(id string) (models.Product, bool) {
	mu.RLock()
	defer mu.RUnlock()
	p, ok := byID[id]
	return p, ok
}
