package action

import (
	"fmt"

	"github.com/0xhunterkiller/berry/internal/models"
)

type actionService struct {
	store ActionStoreIface
}

func NewActionService(store ActionStoreIface) ActionServiceIface {
	return &actionService{store: store}
}

func (svc *actionService) createAction(name string, description string) error {

	var action models.ActionModel

	action.Name = name
	action.Description = description

	err := action.Validate()
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	err = svc.store.createAction(&action)
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}

	return nil
}

func (svc *actionService) deleteAction(id string) error {
	err := svc.store.deleteAction(id)
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}
	return nil
}

type ActionServiceIface interface {
	createAction(name string, description string) error
	deleteAction(id string) error
}

var _ ActionServiceIface = &actionService{}
