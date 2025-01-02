package role

import (
	"fmt"

	"github.com/0xhunterkiller/berry/internal/models"
)

type roleService struct {
	store RoleStoreIface
}

func NewRoleService(store RoleStoreIface) RoleServiceIface {
	return &roleService{store: store}
}

func (svc *roleService) createRole(name string, description string) (string, error) {

	var role models.RoleModel

	role.Name = name
	role.Description = description

	err := role.Validate()
	if err != nil {
		return "", err
	}

	err = svc.store.createRole(&role)
	if err != nil {
		return "", err
	}

	if role.ID == "" {
		return "", fmt.Errorf("db error: role id is not available")
	}

	return role.ID, nil
}

func (svc *roleService) deleteRole(id string) error {
	err := svc.store.deleteRole(id)
	if err != nil {
		return err
	}
	return nil
}

type RoleServiceIface interface {
	createRole(name string, description string) (string, error)
	deleteRole(id string) error
}

var _ RoleServiceIface = &roleService{}
