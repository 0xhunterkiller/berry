package permission

import (
	"fmt"

	"github.com/0xhunterkiller/berry/internal/models"
)

type permissionService struct {
	store PermissionStoreIface
}

func NewPermissionService(store PermissionStoreIface) PermissionServiceIface {
	return &permissionService{store: store}
}

func (svc *permissionService) createPermission(name string, description string) (string, error) {

	var permission models.PermissionModel

	permission.Name = name
	permission.Description = description

	err := permission.Validate()
	if err != nil {
		return "", fmt.Errorf("validation error: %w", err)
	}

	err = svc.store.createPermission(&permission)
	if err != nil {
		return "", fmt.Errorf("db error: %w", err)
	}

	if permission.ID == "" {
		return "", fmt.Errorf("db error: permission id not available")
	}

	return permission.ID, nil
}

func (svc *permissionService) deletePermission(id string) error {
	err := svc.store.deletePermission(id)
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}
	return nil
}

type PermissionServiceIface interface {
	createPermission(name string, description string) (string, error)
	deletePermission(id string) error
}

var _ PermissionServiceIface = &permissionService{}
