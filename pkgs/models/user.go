package models

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type User struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Spec       UserSpec `yaml:"spec"`
}

type UserSpec struct {
	Description string   `yaml:"description"`
	Name        string   `yaml:"name"`
	Roles       []string `yaml:"roles"`
}

func NewUser(apiVersion string, kind string, description string, name string, roles []string) User {
	res := User{
		APIVersion: apiVersion,
		Kind:       kind,
		Spec: UserSpec{
			Description: description,
			Name:        name,
			Roles:       roles,
		},
	}

	yamlData, err := yaml.Marshal(&res)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(yamlData))
	return res
}
