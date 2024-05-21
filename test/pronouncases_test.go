package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPronounCases(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/pronouncases/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var pronouns map[string]interface{}
	err = json.Unmarshal(content, &pronouns)
	require.NoError(t, err)

	cases := []string{"nominative", "genitive", "accusative", "dative", "instrumental", "locative"}

	// Validate personal pronouns
	personalPronouns, ok := pronouns["personal"].(map[string]interface{})
	require.True(t, ok)

	for _, aspect := range []string{"singular", "masculine", "feminine", "neuter", "plural"} {
		assert.Contains(t, personalPronouns, aspect)
		aspectMap, ok := personalPronouns[aspect].([]interface{})
		require.True(t, ok)

		for _, pronoun := range aspectMap {
			pronounMap, ok := pronoun.(map[string]interface{})
			require.True(t, ok)

			for _, caseName := range cases {
				assert.Contains(t, pronounMap, caseName)
				pronounString, ok := pronounMap[caseName].(string)
				require.True(t, ok)
				assert.NotEmpty(t, pronounString)
			}
		}
	}

	// Validate possessive pronouns
	possessivePronouns, ok := pronouns["possessive"].(map[string]interface{})
	require.True(t, ok)

	for _, gender := range []string{"masculine", "feminine", "neuter", "plural"} {
		assert.Contains(t, possessivePronouns, gender)
		pronounMaps, ok := possessivePronouns[gender].([]interface{})
		require.True(t, ok)

		for _, pronoun := range pronounMaps {
			pronounMap, ok := pronoun.(map[string]interface{})
			require.True(t, ok)

			for _, caseName := range cases {
				assert.Contains(t, pronounMap, caseName)

				if caseName == "accusative" {
					accusativeMap, ok := pronounMap[caseName].(map[string]interface{})
					require.True(t, ok)

					assert.Contains(t, accusativeMap, "animate")
					pronounString, ok := accusativeMap["animate"].(string)
					assert.True(t, ok)
					assert.NotEmpty(t, pronounString)

					assert.Contains(t, accusativeMap, "inanimate")
					pronounString, ok = accusativeMap["inanimate"].(string)
					assert.True(t, ok)
					assert.NotEmpty(t, pronounString)
				} else {
					pronounString, ok := pronounMap[caseName].(string)
					require.True(t, ok)
					assert.NotEmpty(t, pronounString)
				}
			}
		}
	}
}
