package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Inventory struct {
	ID        string
	Name      string
	Providers []Provider
}

type InventoryInterface interface {
	Create() error
	List() ([]Resource, error)
	Update() error
	Delete() error
}

var (
	ErrInvalidInventoryName      = errors.New("invalid name")
	ErrInvalidInventoryProviders = errors.New("invalid providers")
)

func (i *Inventory) IsValid() error {
	if i.Name == "" {
		return ErrInvalidInventoryName
	}
	if len(i.Providers) == 0 {
		return ErrInvalidInventoryProviders
	}
	return nil
}

func NewInventory(name string, providers []Provider) (*Inventory, error) {
	inventory := &Inventory{
		ID:        uuid.New().String(),
		Name:      name,
		Providers: providers,
	}

	err := inventory.IsValid()
	if err != nil {
		return nil, err
	}
	return inventory, nil
}
