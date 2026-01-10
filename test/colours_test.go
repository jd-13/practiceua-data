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
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var colours map[string]interface{}
	err = json.Unmarshal(content, &colours)
	require.NoError(t, err)

	// Ensure we have colours to validate
	require.NotEmpty(t, colours, "colours should not be empty")

	// Validate each colour
	for colourName := range colours {
		colourData, ok := colours[colourName].(map[string]interface{})
		require.True(t, ok, "colour %s should be a map", colourName)

		// Validate gender forms
		for _, gender := range []string{"masculine", "feminine", "neuter", "plural"} {
			assert.Contains(t, colourData, gender, "colour %s should have %s form", colourName, gender)
			declension, ok := colourData[gender].(string)
			require.True(t, ok, "colour %s %s should be a string", colourName, gender)
			assert.NotEmpty(t, declension, "colour %s %s should not be empty", colourName, gender)
		}
	}
}
