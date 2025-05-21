package models

type User struct {
	APIVersion string   `yaml:"apiVersion" json:"apiVersion" validate:"required"`
	Kind       string   `yaml:"kind" json:"kind" validate:"required"`
	Spec       UserSpec `yaml:"spec" json:"spec" validate:"required"`
}

type UserSpec struct {
	Description string   `yaml:"description" json:"description"`
	Name        string   `yaml:"name" json:"name" validate:"required"`
	Roles       []string `yaml:"roles" json:"roles" validate:"required"`
}

func NewUser(apiVersion string, name string, description string, roles []string) User {
	res := User{
		APIVersion: apiVersion,
		Kind:       "User",
		Spec: UserSpec{
			Description: description,
			Name:        name,
			Roles:       roles,
		},
	}

	// yamlData, err := yaml.Marshal(&res)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(yamlData))
	return res
}
