package target

import (
	"encoding/json"
	"github.com/sh3rp/cre/internal/model"
	"os"
	"strings"
)

type fileTarget struct {
	filename string
}

func NewFileTarget(filename string) Target {
	return fileTarget{filename}
}

func (g fileTarget) ReadConfig() (*model.ConfigObject, error) {
	file, err := os.Open(g.filename)
	stat, err := file.Stat()

	tokens := strings.Split(stat.Name(), ".")
	name := tokens[0]

	if err != nil {
		return nil, err
	}
	var jsonObj map[string]interface{}

	err = json.NewDecoder(file).Decode(&jsonObj)

	if err != nil {
		return nil, err
	}

	return getObject(name, jsonObj), nil
}

func getObject(name string, jsonObj map[string]interface{}) *model.ConfigObject {
	configObj := &model.ConfigObject{Name: name}

	for k, v := range jsonObj {
		var t model.DataType
		var value interface{}

		switch v.(type) {
		case map[string]interface{}:
			t = model.DataType_ConfigObject
			value = getObject(k, v.(map[string]interface{}))
		case int:
			t = model.DataType_Int
			value = v
		case float32:
		case float64:
			t = model.DataType_Float
			value = v
		case string:
			t = model.DataType_String
			value = v
		case bool:
			t = model.DataType_Boolean
			value = v
		default:
			t = model.DataType_String
			value = v
		}

		configObj.AddField(&model.ConfigField{
			FieldName: k,
			Value:     value,
			DataType:  t,
		})
	}

	return configObj
}
