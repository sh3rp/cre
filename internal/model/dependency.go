package model

type ConfigDependency struct {
	LocalField       *ConfigField
	LocalFieldHash   string
	ForeignField     *ConfigField
	ForeignFieldHash string
}
