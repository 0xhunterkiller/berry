package resource

import (
	"fmt"

	"github.com/0xhunterkiller/berry/internal/models"
)

type resourceService struct {
	store ResourceStoreIface
}

func NewResourceService(store ResourceStoreIface) ResourceServiceIface {
	return &resourceService{store: store}
}

func (svc *resourceService) createResource(name string, description string) error {

	var resource models.ResourceModel

	resource.Name = name
	resource.Description = description

	err := resource.Validate()
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	err = svc.store.createResource(&resource)
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}

	return nil
}

func (svc *resourceService) deleteResource(id string) error {
	err := svc.store.deleteResource(id)
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}
	return nil
}

type ResourceServiceIface interface {
	createResource(name string, description string) error
	deleteResource(id string) error
}

var _ ResourceServiceIface = &resourceService{}
