package model

type ConfigGroup struct {
	Name    string
	Context string
	Tables  []ConfigObject
}
