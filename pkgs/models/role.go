package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Role struct {
	APIVersion string   `yaml:"apiVersion" json:"apiVersion" validate:"required"`
	Kind       string   `yaml:"kind" json:"kind" validate:"required"`
	Spec       RoleSpec `yaml:"spec" json:"spec" validate:"required"`
}

type RoleSpec struct {
	Description string         `yaml:"description" json:"description"`
	Name        string         `yaml:"name" json:"name" validate:"required"`
	Resources   []ResourceItem `yaml:"resources" json:"resources" validate:"required"`
}

func NewRole(apiVersion string, name string, description string, resources []ResourceItem) (*Role, error) {
	res := Role{
		APIVersion: apiVersion,
		Kind:       "Role",
		Spec: RoleSpec{
			Description: description,
			Name:        name,
			Resources:   resources,
		},
	}
	validate := validator.New()

	err := validate.Struct(res)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed validation tag '%s'\n", err.Field(), err.Tag())
		}
		return nil, err
	}

	// yamlData, err := yaml.Marshal(&res)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(yamlData))
	return &res, nil
}
