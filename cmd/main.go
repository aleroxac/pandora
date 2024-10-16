package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aleroxac/pandora/internal/entity"
	"github.com/aleroxac/pandora/internal/infra/providers"
)

func main() {
	gcpProvider := providers.GCPProvider{}
	resources, err := gcpProvider.ListInstances(log.Writer(), "aleroxac")
	if err != nil {
		log.Fatalf("Error listing instances: %v", err)
	}

	inventory := entity.Inventory{
		ID:   "inventory-001",
		Name: "test",
		Providers: []entity.Provider{
			{
				ID:        "provider-001",
				Name:      "gcp",
				AccountID: "aleroxac",
				Resources: resources,
			},
		},
	}

	jsonData, err := json.MarshalIndent(inventory, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling inventory to JSON: %v", err)
	}

	os.Stdout.Write(jsonData)
}
