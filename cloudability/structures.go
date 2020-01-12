package cloudability

import (
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)

func flattenVerification(in cloudability.Verification) []map[string]interface{} {
	var out = make([]map[string]interface{}, 1, 1)
	m := make(map[string]interface{})
	m["state"] = in.State
	m["last_verification_attempted_at"] = in.LastVerificationAttemptedAt
	m["message"] = in.Message
	out[0]= m
	return out
}

func flattenAuthorization(in cloudability.Authorization) []map[string]interface{} {
	var out = make([]map[string]interface{}, 1, 1)
	m := make(map[string]interface{})
	m["type"] = in.Type
	m["role_name"] = in.RoleName
	m["external_id"] = in.ExternalId
	out[0]= m
	return out
}

func flattenStatements(in []*cloudability.BusinessMappingStatement) []map[string]interface{} {
	var out = make([]map[string]interface{}, len(in), len(in))
	for i, v := range in {
		m := make(map[string]interface{})
		m["match_expression"] = v.MatchExpression
		m["value_expression"] = v.ValueExpression
		out[i] = m
	}
	return out
}

func inflateStatements(in []map[string]interface{}) []cloudability.BusinessMappingStatement {
	var out = make([]cloudability.BusinessMappingStatement, len(in))
	for i, v := range in {
		m := cloudability.BusinessMappingStatement{
			MatchExpression: v["match_expression"].(string),
			ValueExpression: v["value_expression"].(string),
		}
		out[i] = m
	}
	return out
}

func flattenFilters(in []cloudability.ViewFilter) []map[string]interface{} {
	var out = make([]map[string]interface{}, len(in), len(in))
	for i, v := range in {
		m := make(map[string]interface{})
		m["field"] = v.Field
		m["comparator"] = v.Comparator
		m["value"] = v.Value
		out[i] = m
	}
	return out
}