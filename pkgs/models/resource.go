package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Resource struct {
	APIVersion string       `yaml:"apiVersion" json:"apiVersion" validate:"required"`
	Kind       string       `yaml:"kind" json:"kind" validate:"required"`
	Spec       ResourceSpec `yaml:"spec" json:"spec" validate:"required"`
}

type ResourceSpec struct {
	Description string   `yaml:"description" json:"description"`
	Name        string   `yaml:"name" json:"name" validate:"required"`
	Verbs       []string `yaml:"verbs" json:"verbs" validate:"required"`
}

func (dat Resource) Validate() error {
	validate := validator.New()

	err := validate.Struct(dat)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("field '%s' failed validation tag '%s'", err.Field(), err.Tag())
		}
	}
	return nil
}

func NewResource(apiVersion string, name string, description string, verbs []string) (*Resource, error) {
	res := Resource{
		APIVersion: apiVersion,
		Kind:       "Resource",
		Spec: ResourceSpec{
			Description: description,
			Name:        name,
			Verbs:       verbs,
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
