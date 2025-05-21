package models

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type ResourceItem struct {
	Name  string   `yaml:"name"`
	Verbs []string `yaml:"verbs"`
}

func NewResourceItem(resourceName string, verbs []string) ResourceItem {
	res := ResourceItem{
		Name:  resourceName,
		Verbs: verbs,
	}

	yamlData, err := yaml.Marshal(&res)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(yamlData))

	return res
}
