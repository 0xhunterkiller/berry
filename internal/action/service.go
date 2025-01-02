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

func (svc *actionService) createAction(name string, description string) (string, error) {

	var action models.ActionModel

	action.Name = name
	action.Description = description

	err := action.Validate()
	if err != nil {
		return "", err
	}

	err = svc.store.createAction(&action)
	if err != nil {
		return "", err
	}

	if action.ID == "" {
		return "", fmt.Errorf("db error: action id not available")
	}

	return action.ID, nil
}

func (svc *actionService) deleteAction(id string) error {
	err := svc.store.deleteAction(id)
	if err != nil {
		return err
	}
	return nil
}

type ActionServiceIface interface {
	createAction(name string, description string) (string, error)
	deleteAction(id string) error
}

var _ ActionServiceIface = &actionService{}
