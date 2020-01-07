package cloudability

import (
	"testing"
	"reflect"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)


func TestFlattenVerfication(t *testing.T) {
	cases := []struct {
		verification *cloudability.Verification
		expected []map[string]interface{}
	}{
		{
			verification: &cloudability.Verification{
				State: "test-state",
				LastVerificationAttemptedAt: "test-lastVerificationAttemptedAt",
				Message: "test-message",
			},
			expected: []map[string]interface{}{
				{
					"state": "test-state",
					"last_verification_attempted_at": "test-lastVerificationAttemptedAt",
					"message": "test-message",
				},
			},
		},
	}
	for _, c := range cases {
		out := flattenVerification(*c.verification)
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}
}

func TestFlattenAuthorization(t *testing.T) {
	cases := []struct {
		authorization *cloudability.Authorization
		expected []map[string]interface{}
	}{
		{
			authorization: &cloudability.Authorization{
				Type: "test-type",
				RoleName: "test-role_name",
				ExternalId: "test-external_id",
			},
			expected: []map[string]interface{}{
				{
					"type": "test-type",
					"role_name": "test-role_name",
					"external_id": "test-external_id",
				},
			},
		},
	}
	for _, c := range cases {
		out := flattenAuthorization(*c.authorization)
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}
}
