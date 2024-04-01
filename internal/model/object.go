package model

type ConfigObject struct {
	Name string

	fields       []*ConfigField
	dependencies []*ConfigDependency
}

func (t *ConfigObject) AddField(field *ConfigField) *ConfigField {
	t.fields = append(t.fields, field)
	field.parent = t
	return field
}

func (t *ConfigObject) AddDependency(dependency *ConfigDependency) {
	t.dependencies = append(t.dependencies, dependency)
}

func (t *ConfigObject) GetFields() []*ConfigField {
	return t.fields
}

func (t *ConfigObject) HasDependencies() bool {
	return t.dependencies != nil && len(t.dependencies) > 0
}
