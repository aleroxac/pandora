package entity_test

import (
	"testing"

	"github.com/aleroxac/pandora/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestInventory_NewInventory(t *testing.T) {
	i, err := entity.NewInventory(
		"Inventory 1",
		[]entity.Provider{
			{},
		},
	)
	assert.Nil(t, err)

	assert.NotNil(t, i.Name)
	assert.NotEmpty(t, i.Name)

	assert.NotNil(t, i.Providers)
	assert.NotEmpty(t, i.Providers)

	assert.Equal(t, "Inventory 1", i.Name)
}

func TestInventory_InvalidName(t *testing.T) {
	_, err := entity.NewInventory(
		"",
		[]entity.Provider{
			{},
		},
	)
	assert.ErrorIs(t, err, entity.ErrInvalidInventoryName)
}

func TestInventory_InvalidProviders(t *testing.T) {
	_, err := entity.NewInventory(
		"Inventory 1",
		[]entity.Provider{},
	)
	assert.ErrorIs(t, err, entity.ErrInvalidInventoryProviders)
}
