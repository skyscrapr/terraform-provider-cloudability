package cloudability

import (
	"testing"
	"reflect"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)


func TestFlattenFilters(t *testing.T) {
	cases := []struct {
		filters []cloudability.ViewFilter
		expected []map[string]interface{}
	}{
		{
			filters: []cloudability.ViewFilter{
				cloudability.ViewFilter{
					Field: "test-field",
					Comparator: "test-comparator",
					Value: "test-value",
				},
			},
			expected: []map[string]interface{}{
				{
					"field": "test-field",
					"comparator": "test-comparator",
					"value": "test-value",
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