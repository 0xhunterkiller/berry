package models

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

func NewResource(apiVersion string, name string, description string, verbs []string) Resource {
	res := Resource{
		APIVersion: apiVersion,
		Kind:       "Resource",
		Spec: ResourceSpec{
			Description: description,
			Name:        name,
			Verbs:       verbs,
		},
	}

	// yamlData, err := yaml.Marshal(&res)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(yamlData))
	return res
}
