package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigField(t *testing.T) {
	field1 := &ConfigField{
		FieldName: "the_string",
		DataType:  DataType_String,
		Value:     "a test value",
	}
	field2 := &ConfigField{
		FieldName: "the_int",
		DataType:  DataType_Int,
		Value:     1,
	}
	field3 := &ConfigField{
		FieldName: "the_bool",
		DataType:  DataType_Boolean,
		Value:     true,
	}

	field4 := &ConfigField{
		FieldName: "the_foreign_int",
		DataType:  DataType_Int,
		Value:     2,
	}

	table1 := ConfigObject{
		Name: "stuff",
	}

	table1.AddField(field1)
	table1.AddField(field2)

	table2 := &ConfigObject{
		Name: "more_stuff",
	}

	table2.AddField(field3)
	table2.AddField(field4).WithDependency(field1)
	assert.True(t, field4.HasDependencyOn(field1))
	assert.Equal(t, field4.GetDependency(), field1)
}
