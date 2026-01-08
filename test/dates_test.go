package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDates(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/dates/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	require.NoError(t, err)

	// Validate ordinal numbers
	numbers, ok := data["numbers"].(map[string]interface{})
	require.True(t, ok)

	for key, value := range numbers {
		assert.NotEmpty(t, key)

		ordinal, ok := value.(map[string]interface{})
		require.True(t, ok)

		// Validate nominative
		assert.Contains(t, ordinal, "nominative")
		nominative, ok := ordinal["nominative"].(string)
		assert.True(t, ok)
		assert.NotEmpty(t, nominative)

		// Validate genitive
		assert.Contains(t, ordinal, "genitive")
		genitive, ok := ordinal["genitive"].(string)
		assert.True(t, ok)
		assert.NotEmpty(t, genitive)

		// Validate locative
		assert.Contains(t, ordinal, "locative")
		locative, ok := ordinal["locative"].([]interface{})
		assert.True(t, ok)
		assert.NotEmpty(t, locative)
		for _, loc := range locative {
			locStr, ok := loc.(string)
			assert.True(t, ok)
			assert.NotEmpty(t, locStr)
		}
	}

	// Validate months
	months, ok := data["months"].([]interface{})
	require.True(t, ok)
	require.NotEmpty(t, months)

	for _, monthData := range months {
		month, ok := monthData.(map[string]interface{})
		require.True(t, ok)

		// Validate nominative
		assert.Contains(t, month, "nominative")
		nominative, ok := month["nominative"].(string)
		assert.True(t, ok)
		assert.NotEmpty(t, nominative)

		// Validate genitive
		assert.Contains(t, month, "genitive")
		genitive, ok := month["genitive"].(string)
		assert.True(t, ok)
		assert.NotEmpty(t, genitive)

		// Validate locative
		assert.Contains(t, month, "locative")
		locative, ok := month["locative"].([]interface{})
		assert.True(t, ok)
		assert.NotEmpty(t, locative)
		for _, loc := range locative {
			locStr, ok := loc.(string)
			assert.True(t, ok)
			assert.NotEmpty(t, locStr)
		}
	}
}
