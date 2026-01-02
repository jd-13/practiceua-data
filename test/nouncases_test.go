package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNounCases(t *testing.T) {
	// Read the noun cases file
	content, err := os.ReadFile("../data/nouncases/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var nouns map[string]interface{}
	err = json.Unmarshal(content, &nouns)
	require.NoError(t, err)

	// Read the rules file to validate caseRule values
	rulesContent, err := os.ReadFile("../data/nouncaserules/v1.json")
	require.NoError(t, err)
	var rules map[string]interface{}
	err = json.Unmarshal(rulesContent, &rules)
	require.NoError(t, err)

	// Extract valid rule names from the rules structure
	validRules := make(map[string]bool)
	singularRules, ok := rules["singular"].(map[string]interface{})
	require.True(t, ok)
	for _, caseRules := range singularRules {
		caseRulesMap, ok := caseRules.(map[string]interface{})
		require.True(t, ok)
		for ruleName := range caseRulesMap {
			validRules[ruleName] = true
		}
	}
	pluralRules, ok := rules["plural"].(map[string]interface{})
	require.True(t, ok)
	for _, caseRules := range pluralRules {
		caseRulesMap, ok := caseRules.(map[string]interface{})
		require.True(t, ok)
		for ruleName := range caseRulesMap {
			validRules[ruleName] = true
		}
	}

	cases := []string{"nominative", "genitive", "accusative", "dative", "instrumental", "locative", "vocative"}

	for nounWord, noun := range nouns {
		nounMap, ok := noun.(map[string]interface{})
		require.True(t, ok, "noun %s should be a map", nounWord)

		for _, singularPlural := range []string{"singular", "plural"} {
			assert.Contains(t, nounMap, singularPlural, "noun %s should have %s", nounWord, singularPlural)

			quantityMap, ok := nounMap[singularPlural].(map[string]interface{})
			require.True(t, ok, "noun %s %s should be a map", nounWord, singularPlural)

			for _, caseName := range cases {
				// Plural doesn't have vocative case
				if singularPlural == "plural" && caseName == "vocative" {
					assert.NotContains(t, quantityMap, caseName, "noun %s plural should not have vocative", nounWord)
					continue
				}

				assert.Contains(t, quantityMap, caseName, "noun %s %s should have %s", nounWord, singularPlural, caseName)

				caseMap, ok := quantityMap[caseName].(map[string]interface{})
				require.True(t, ok, "noun %s %s %s should be a map", nounWord, singularPlural, caseName)

				// Must always have text
				assert.Contains(t, caseMap, "text", "noun %s %s %s should have text", nounWord, singularPlural, caseName)
				text, ok := caseMap["text"].(string)
				require.True(t, ok, "noun %s %s %s text should be a string", nounWord, singularPlural, caseName)
				assert.NotEmpty(t, text, "noun %s %s %s text should not be empty", nounWord, singularPlural, caseName)

				// Singular nominative doesn't have a case rule
				if !(singularPlural == "singular" && caseName == "nominative") {
					assert.Contains(t, caseMap, "caseRule", "noun %s %s %s should have caseRule", nounWord, singularPlural, caseName)
					caseRule, ok := caseMap["caseRule"].(string)
					require.True(t, ok, "noun %s %s %s caseRule should be a string", nounWord, singularPlural, caseName)
					assert.NotEmpty(t, caseRule, "noun %s %s %s caseRule should not be empty", nounWord, singularPlural, caseName)

					// Validate that the caseRule exists in the rules file
					assert.True(t, validRules[caseRule], "noun %s %s %s has invalid caseRule '%s'", nounWord, singularPlural, caseName, caseRule)
				}
			}
		}
	}
}