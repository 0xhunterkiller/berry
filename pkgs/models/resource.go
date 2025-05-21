package models

type Resource struct {
	APIVersion string       `yaml:"apiVersion"`
	Kind       string       `yaml:"kind"`
	Spec       ResourceSpec `yaml:"spec"`
}

type ResourceSpec struct {
	Description string   `yaml:"description"`
	Name        string   `yaml:"name"`
	Verbs       []string `yaml:"verbs"`
}

func NewResource(apiVersion string, kind string, description string, name string, verbs []string) Resource {
	res := Resource{
		APIVersion: apiVersion,
		Kind:       kind,
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
