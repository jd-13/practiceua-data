package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClothing(t *testing.T) {
	content, err := os.ReadFile("../data/clothing/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	var clothing map[string]interface{}
	err = json.Unmarshal(content, &clothing)
	require.NoError(t, err)

	require.NotEmpty(t, clothing)

	for key, item := range clothing {
		entry, ok := item.(map[string]interface{})
		require.True(t, ok, "item %q should be an object", key)

		assert.Contains(t, entry, "singluar")
		singular, ok := entry["singluar"].(string)
		assert.True(t, ok, "item %q singluar should be a string", key)
		assert.NotEmpty(t, singular, "item %q singluar should not be empty", key)

		assert.Contains(t, entry, "plural")
		plural, ok := entry["plural"].(string)
		assert.True(t, ok, "item %q plural should be a string", key)
		assert.NotEmpty(t, plural, "item %q plural should not be empty", key)
	}
}
