package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ResourceItem struct {
	Name  string   `yaml:"name" json:"name" validate:"required"`
	Verbs []string `yaml:"verbs" json:"verbs" validate:"required"`
}

func NewResourceItem(resourceName string, verbs []string) (*ResourceItem, error) {
	res := ResourceItem{
		Name:  resourceName,
		Verbs: verbs,
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
