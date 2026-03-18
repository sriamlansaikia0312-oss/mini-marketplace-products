package data

import "github.com/netapp/mini-marketplace-products/internal/models"

// Products is fixed in-memory seed data (mini MCP).
var Products = []models.Product{
	{ID: "np-aws-console", Name: "NetApp Console", Cloud: "aws", SKU: "NTAP-CONSOLE-AWS", Family: "Console"},
	{ID: "np-azure-console", Name: "NetApp Console", Cloud: "azure", SKU: "NTAP-CONSOLE-AZ", Family: "Console"},
	{ID: "np-gcp-console", Name: "NetApp Console", Cloud: "gcp", SKU: "NTAP-CONSOLE-GCP", Family: "Console"},
	{ID: "np-aws-dii", Name: "Data Infrastructure Insights", Cloud: "aws", SKU: "NTAP-DII-AWS", Family: "DII"},
	{ID: "np-azure-dii", Name: "Data Infrastructure Insights", Cloud: "azure", SKU: "NTAP-DII-AZ", Family: "DII"},
	{ID: "np-aws-ic", Name: "IntraCluster", Cloud: "aws", SKU: "NTAP-IC-AWS", Family: "IntraCluster"},
}
