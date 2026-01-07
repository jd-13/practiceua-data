package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMonths(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/months/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var months map[string]interface{}
	err = json.Unmarshal(content, &months)
	require.NoError(t, err)

	// Ensure we have months to validate
	require.NotEmpty(t, months, "months should not be empty")

	// Validate each month
	for monthName := range months {
		monthData, ok := months[monthName].(map[string]interface{})
		require.True(t, ok, "month %s should be a map", monthName)

		// Validate text field
		assert.Contains(t, monthData, "text", "month %s should have text", monthName)
		textValue, ok := monthData["text"].(string)
		require.True(t, ok, "month %s text should be a string", monthName)
		assert.NotEmpty(t, textValue, "month %s text should not be empty", monthName)

		// Validate meaning field
		assert.Contains(t, monthData, "meaning", "month %s should have meaning", monthName)
		meaningValue, ok := monthData["meaning"].(string)
		require.True(t, ok, "month %s meaning should be a string", monthName)
		assert.NotEmpty(t, meaningValue, "month %s meaning should not be empty", monthName)
	}
}
