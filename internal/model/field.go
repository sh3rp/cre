package model

type DataType int

const (
	DataType_Int DataType = iota
	DataType_Float
	DataType_String
	DataType_Boolean
	DataType_ConfigObject
)

type ConfigField struct {
	FieldName string
	DataType  DataType
	Value     interface{}

	parent *ConfigObject
}

func NewConfigField(name string, dType DataType) *ConfigField {
	return &ConfigField{FieldName: name, DataType: dType}
}

func (f *ConfigField) WithDependency(foreign *ConfigField) *ConfigField {
	f.parent.AddDependency(&ConfigDependency{
		LocalField:   f,
		ForeignField: foreign,
	})
	return f
}

func (f *ConfigField) GetDependency() *ConfigField {
	if !f.parent.HasDependencies() {
		return nil
	}

	for _, dep := range f.parent.dependencies {
		if dep.LocalField == f {
			return dep.ForeignField
		}
	}

	return nil
}

func (f *ConfigField) HasDependencyOn(field *ConfigField) bool {
	return f.GetDependency() == field
}
