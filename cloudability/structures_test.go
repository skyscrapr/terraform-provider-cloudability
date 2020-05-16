package cloudability

import (
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
	"reflect"
	"testing"
)

func TestFlattenVerfication(t *testing.T) {
	cases := []struct {
		verification *cloudability.Verification
		expected     []map[string]interface{}
	}{
		{
			verification: &cloudability.Verification{
				State:                       "test-state",
				LastVerificationAttemptedAt: "test-lastVerificationAttemptedAt",
				Message:                     "test-message",
			},
			expected: []map[string]interface{}{
				{
					"state":                          "test-state",
					"last_verification_attempted_at": "test-lastVerificationAttemptedAt",
					"message":                        "test-message",
				},
			},
		},
	}
	for _, c := range cases {
		out := flattenVerification(c.verification)
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}
}

func TestFlattenAuthorization(t *testing.T) {
	cases := []struct {
		authorization *cloudability.Authorization
		expected      []map[string]interface{}
	}{
		{
			authorization: &cloudability.Authorization{
				Type:       "test-type",
				RoleName:   "test-role_name",
				ExternalID: "test-external_id",
			},
			expected: []map[string]interface{}{
				{
					"type":        "test-type",
					"role_name":   "test-role_name",
					"external_id": "test-external_id",
				},
			},
		},
	}
	for _, c := range cases {
		out := flattenAuthorization(c.authorization)
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}
}

func TestFlattenStatements(t *testing.T) {
	cases := []struct {
		statements []*cloudability.BusinessMappingStatement
		expected   []map[string]interface{}
	}{
		{
			statements: []*cloudability.BusinessMappingStatement{
				{
					MatchExpression: "test-match_expression_1",
					ValueExpression: "test-value_expression_1",
				},
				{
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
		out := flattenStatements(c.statements)
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}
}

func TestFlattenFilters(t *testing.T) {
	cases := []struct {
		filters  []*cloudability.ViewFilter
		expected []map[string]interface{}
	}{
		{
			filters: []*cloudability.ViewFilter{
				{
					Field:      "test-field",
					Comparator: "test-comparator",
					Value:      "test-value",
				},
			},
			expected: []map[string]interface{}{
				{
					"field":      "test-field",
					"comparator": "test-comparator",
					"value":      "test-value",
				},
			},
		},
	}
	for _, c := range cases {
		out := flattenFilters(c.filters)
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}
}
