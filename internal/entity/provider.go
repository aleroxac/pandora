package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Provider struct {
	ID        string
	Name      string
	AccountID string
	Resources []Resource
}

type ProviderInterface interface {
	Create() error
	List() ([]Resource, error)
	Update() error
	Delete() error
}

type Instance struct {
	ID                string
	Name              string
	Status            string
	Image             string
	Region            string
	Zone              string
	Tier              string
	Spot              bool
	Disks             interface{}
	NetworkInterfaces interface{}
	CreationTimestamp string
	Tags              map[string]string
}

func (p *Provider) isValid() error {
	if p.Name == "" {
		return errors.New("invalid name")
	}
	if p.AccountID == "" {
		return errors.New("invalid account_id")
	}
	return nil
}

func NewProvider(name string, account_id string, resources []Resource) (*Provider, error) {
	provider := &Provider{
		ID:        uuid.New().String(),
		Name:      name,
		AccountID: account_id,
		Resources: resources,
	}

	err := provider.isValid()
	if err != nil {
		return nil, err
	}
	return provider, nil
}
