package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFruitsAndVegetables(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/fruitsandvegetables/v1.json")
	require.NoError(t, err)
	require.NotEmpty(t, content)

	// Read the json
	var items map[string]map[string]string
	err = json.Unmarshal(content, &items)
	require.NoError(t, err)

	// Ensure we have items to validate
	require.NotEmpty(t, items, "fruitsandvegetables should not be empty")

	// Validate each item
	for itemName, itemData := range items {
		// Validate singular and plural forms
		for _, form := range []string{"singluar", "plural"} {
			value, ok := itemData[form]
			require.True(t, ok, "item %s should have %s form", itemName, form)
			assert.NotEmpty(t, value, "item %s %s form should not be empty", itemName, form)
		}
	}
}
