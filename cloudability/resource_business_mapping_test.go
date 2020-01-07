package cloudability

import (
	"testing"
	"reflect"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)


func TestFlattenStatements(t *testing.T) {
	cases := []struct {
		statement []cloudability.BusinessMappingStatement
		expected []map[string]interface{}
	}{
		{
			statement: []cloudability.BusinessMappingStatement{
				cloudability.BusinessMappingStatement{
					MatchExpression: "test-match_expression_1",
					ValueExpression: "test-value_expression_1",
				},
				cloudability.BusinessMappingStatement{
					MatchExpression: "test-match_expression_2",
					ValueExpression: "test-value_expression_2",
				},
			},
			expected: []map[string]interface{}{
				{
					"match_expression": "test-match_expression_1",
					"value_expression": "test-value_expression_1",
				},
				{
					"match_expression": "test-match_expression_2",
					"value_expression": "test-value_expression_2",
				},
			},
		},
	}
	for _, c := range cases {
		out := flattenStatements(c.statement)
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}
}