package models

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Role struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Spec       RoleSpec `yaml:"spec"`
}

type RoleSpec struct {
	Description string         `yaml:"description"`
	Name        string         `yaml:"name"`
	Resources   []ResourceItem `yaml:"resources"`
}

func NewRole(apiVersion string, kind string, description string, name string, resources []ResourceItem) Role {
	res := Role{
		APIVersion: apiVersion,
		Kind:       kind,
		Spec: RoleSpec{
			Description: description,
			Name:        name,
			Resources:   resources,
		},
	}

	yamlData, err := yaml.Marshal(&res)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(yamlData))
	return res
}
