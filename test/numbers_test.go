package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNumbers(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/numbers/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var numbers map[string]interface{}
	err = json.Unmarshal(content, &numbers)
	require.NoError(t, err)

	// Validate cardinal
	assert.Contains(t, numbers, "cardinal")
	for k, v := range numbers["cardinal"].(map[string]interface{}) {
		if _, ok := v.(string); ok {
			// Do nothing
		} else if values, ok := v.([]interface{}); ok {
			assert.Len(t, values, 3)
			for _, value := range values {
				_, ok = value.(string)
				assert.True(t, ok)
			}
		} else {
			t.Errorf("Invalid type for key %s", k)
		}
	}

	// Validate ordinal
	assert.Contains(t, numbers, "ordinal")
	for k, v := range numbers["ordinal"].(map[string]interface{}) {
		if values, ok := v.([]interface{}); ok {
			assert.Len(t, values, 4)
			for _, value := range values {
				_, ok = value.(string)
				assert.True(t, ok)
			}
		} else {
			t.Errorf("Invalid type for key %s", k)
		}
	}
}