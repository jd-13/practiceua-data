package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNounCaseRules(t *testing.T) {
	// Read the file
	content, err := os.ReadFile("../data/nouncaserules/v1.json")
	require.NotEmpty(t, content)
	require.NoError(t, err)

	// Read the json
	var rules map[string]interface{}
	err = json.Unmarshal(content, &rules)
	require.NoError(t, err)

	cases := []string{"genitive", "accusative", "dative", "instrumental", "locative", "vocative"}

	// Validate singular rules
	assert.Contains(t, rules, "singular rules")
	singularRules, ok := rules["singular rules"].(map[string]interface{})
	require.True(t, ok)

	for _, caseName := range cases {
		assert.Contains(t, singularRules, caseName)
		caseRules, ok := singularRules[caseName].(map[string]interface{})
		require.True(t, ok)

		// Each case should have at least one rule
		assert.NotEmpty(t, caseRules)

		// Validate each rule is a non-empty string
		for key, rule := range caseRules {
			assert.NotEmpty(t, key)
			ruleString, ok := rule.(string)
			assert.True(t, ok)
			assert.NotEmpty(t, ruleString)
		}
	}

	// Validate plural rules
	assert.Contains(t, rules, "plural rules")
	pluralRules, ok := rules["plural rules"].(map[string]interface{})
	require.True(t, ok)

	pluralCases := []string{"nominative", "genitive", "accusative", "dative", "instrumental", "locative"}

	for _, caseName := range pluralCases {
		assert.Contains(t, pluralRules, caseName)
		caseRules, ok := pluralRules[caseName].(map[string]interface{})
		require.True(t, ok)

		// Each case should have at least one rule
		assert.NotEmpty(t, caseRules)

		// Validate each rule is a non-empty string
		for key, rule := range caseRules {
			assert.NotEmpty(t, key)
			ruleString, ok := rule.(string)
			assert.True(t, ok)
			assert.NotEmpty(t, ruleString)
		}
	}
}