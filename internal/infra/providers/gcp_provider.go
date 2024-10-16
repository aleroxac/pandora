package providers

import (
	"context"
	"fmt"
	"io"
	"strconv"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/aleroxac/pandora/internal/entity"
	"google.golang.org/api/iterator"
)

type GCPProvider struct{}

func (g GCPProvider) ListInstances(w io.Writer, accountID string) ([]entity.Resource, error) {
	ctx := context.Background()
	client, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create compute client: %v", err)
	}
	defer client.Close()

	projectID := accountID
	zone := "europe-west1-b"

	req := &computepb.ListInstancesRequest{
		Project: projectID,
		Zone:    zone,
	}

	it := client.List(ctx, req)
	var resources []entity.Resource
	for {
		instance, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to list instances: %v", err)
		}

		resource := entity.Resource{
			ID:   strconv.FormatUint(instance.GetId(), 10),
			Name: instance.GetName(),
			Type: "GCP Instance",
			Spec: instance,
		}
		resources = append(resources, resource)
		fmt.Fprintf(w, "Instance found: %s\n", instance.GetName())
	}

	return resources, nil
}
