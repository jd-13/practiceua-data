package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestColours(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/colours/v1.json")
	require.NoError(t, err)
	require.NotEmpty(t, content)

	// Read the json
	var colours map[string]interface{}
	err = json.Unmarshal(content, &colours)
	require.NoError(t, err)

	// Ensure we have colours to validate
	require.NotEmpty(t, colours, "colours should not be empty")

	// Validate each colour
	for colourName, colourValue := range colours {
		colourData, ok := colourValue.(map[string]interface{})
		require.True(t, ok, "colour %s should be a map", colourName)

		// Validate gender forms
		for _, gender := range []string{"masculine", "feminine", "neuter", "plural"} {
			assert.Contains(t, colourData, gender, "colour %s should have %s form", colourName, gender)

			values, ok := colourData[gender].([]interface{})
			require.True(t, ok, "colour %s %s should be an array", colourName, gender)
			require.NotEmpty(t, values, "colour %s %s should not be empty", colourName, gender)

			for _, v := range values {
				s, ok := v.(string)
				require.True(t, ok, "colour %s %s entries should be strings", colourName, gender)
				assert.NotEmpty(t, s, "colour %s %s entries should not be empty", colourName, gender)
			}
		}
	}
}
