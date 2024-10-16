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

func (i *Inventory) IsValid() error {
	if i.Name == "" {
		return errors.New("invalid name")
	}
	if len(i.Providers) == 0 {
		return errors.New("invalid providers")
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
