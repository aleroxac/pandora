package entity

import (
	"errors"
	"io"

	"github.com/google/uuid"
)

type Resource struct {
	ID   string
	Name string
	Type string
	Spec interface{}
}

type ResourceInterface interface {
	ListInstances(w io.Writer, accountID string) ([]Resource, error)
}

func (r *Resource) isValid() error {
	if r.Name == "" {
		return errors.New("invalid name")
	}
	if r.Type == "" {
		return errors.New("invalid type")
	}
	if r.Spec == nil {
		return errors.New("invalid spec")
	}
	return nil
}

func NewResource(name string, resource_type string, spec interface{}) (*Resource, error) {
	resource := &Resource{
		ID:   uuid.New().String(),
		Name: name,
		Type: resource_type,
		Spec: spec,
	}

	err := resource.isValid()
	if err != nil {
		return nil, err
	}
	return resource, nil
}
