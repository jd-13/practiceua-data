package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCountries(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/countries/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var countries []interface{}
	err = json.Unmarshal(content, &countries)
	require.NoError(t, err)

	// Validate each item
	for _, item := range countries {
		country, ok := item.(map[string]interface{})
		require.True(t, ok)

		assert.Contains(t, country, "flag")
		flag, ok := country["flag"].(string)
		assert.True(t, ok)
		assert.NotEmpty(t, flag)

		assert.Contains(t, country, "country")
		name, ok := country["country"].(string)
		assert.True(t, ok)
		assert.NotEmpty(t, name)

		assert.Contains(t, country, "genitive")
		genitive, ok := country["genitive"].(string)
		assert.True(t, ok)
		assert.NotEmpty(t, genitive)

		assert.Contains(t, country, "preposition")
		preposition, ok := country["preposition"].(string)
		assert.True(t, ok)
		assert.Contains(t, []string{"з", "із"}, preposition)

		assert.Contains(t, country, "languages")
		languages, ok := country["languages"].([]interface{})
		assert.True(t, ok)

		for _, language := range languages {
			_, ok = language.(string)
			assert.True(t, ok)
		}

		assert.Contains(t, country, "nationality")
		nationality, ok := country["nationality"].(map[string]interface{})
		require.True(t, ok)

		for _, key := range []string{"masculine", "feminine", "plural"} {
			assert.Contains(t, nationality, key)
			vals, ok := nationality[key].([]interface{})
			assert.True(t, ok)
			assert.NotEmpty(t, vals)
			for _, v := range vals {
				_, ok = v.(string)
				assert.True(t, ok)
			}
		}
	}
}