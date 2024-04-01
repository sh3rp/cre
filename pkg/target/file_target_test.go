package target

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileTarget(t *testing.T) {
	target := NewFileTarget("../../test/go_version.json")
	config, err := target.ReadConfig()
	assert.NoError(t, err)
	for _, field := range config.GetFields() {
		t.Logf("%s = %+v", field.FieldName, field.Value)
	}
}
