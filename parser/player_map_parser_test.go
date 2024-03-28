package parser

import (
	"testing"
	"reflect"
)

type formationTestCase[]struct {
		formation string
		expected  map[string]interface{}
}

func TestParseFormationStringIntoMap(t *testing.T) {
	testCases := formationTestCase{
		{
			formation: "4-4-2",
			expected: map[string]interface{}{
				"D": 4,
				"M": 4,
				"F":  2,
				"original": "4-4-2",
			},
		},
		{
			formation: "4-3-3",
			expected: map[string]interface{}{
				"D": 4,
				"M": 3,
				"F":  3,
				"original": "4-3-3",
			},
		},
		{
			formation: "4-2-3-1",
			expected: map[string]interface{}{
				"D":  4,
				"DM": 2,
				"AM": 3,
				"F":  1,
				"original": "4-2-3-1",
			},
		},
		{
			formation: "3-4-3",
			expected: map[string]interface{} {
				"D": 3,
				"M": 4,
				"F":  3,
				"original": "3-4-3",
			},
		},
		{
			formation: "4-4-1-1",
			expected: map[string]interface{} {
				"D": 4,
				"M": 4,
				"AM":  1,
				"F":  1,
				"original": "4-4-1-1",
			},
		},
	}

	for _, tc := range testCases {
		actual := ParseFormationStringIntoMap(tc.formation)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, actual)
		}
	}
	
}