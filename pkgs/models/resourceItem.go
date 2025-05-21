package models

type ResourceItem struct {
	Name  string   `yaml:"name" json:"name" validate:"required"`
	Verbs []string `yaml:"verbs" json:"verbs" validate:"required"`
}

func NewResourceItem(resourceName string, verbs []string) ResourceItem {
	res := ResourceItem{
		Name:  resourceName,
		Verbs: verbs,
	}

	// yamlData, err := yaml.Marshal(&res)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(yamlData))

	return res
}
