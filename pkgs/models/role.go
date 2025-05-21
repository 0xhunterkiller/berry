package models

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

func NewRole(apiVersion string, name string, description string, resources []ResourceItem) Role {
	res := Role{
		APIVersion: apiVersion,
		Kind:       "Role",
		Spec: RoleSpec{
			Description: description,
			Name:        name,
			Resources:   resources,
		},
	}

	// yamlData, err := yaml.Marshal(&res)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(yamlData))
	return res
}
