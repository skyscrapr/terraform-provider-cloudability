# Business Metric Resource

Business Mapping is an activity that allows you to map your cloud cost and usage to custom dimensions and custom metrics that are important to report on within your organization.
This resource maps to BUSINESS_METRIC Business Mapping.

## Example Usage

```hcl
resource "cloudability_business_metric" "test" {
	name = "Cost (Surcharge) %s"
	number_format = "number"
	default_value_expression = "METRIC['unblended_cost']"
	pre_match_expression = "DIMENSION['vendor'] == 'Amazon'"
	statement {
		match_expression = "DIMENSION['vendor'] == 'Amazon' || DIMENSION['vendor'] == 'Azure'"
		value_expression = "METRIC['unblended_cost'] * 1.15"
	}
}
```

## Argument Reference

* `name` - (Required) Name for the Business Dimension/Metric.
* `number_format` - (Optional) Can be set to "currency" or "number". Defaults to "number"
* `default_value_expression` - (Optional) If no rule matches then take the value resolved by this expression as the fall back value.
* `pre_match_expression` - (Optional) The preMatchExpression represents a centralized/global expression which is evaluated before all other defined expressions contained within statements section; It can be left empty (without a defined expression) when creating a Business Metric; if left empty, it will be omitted from the response.
* `statement` - (Optional) List of statements

## Attribute Reference

* `index` - An integer value representing the ID for the Business Metric.
* `default_value` - If there are no matches for the statements defined, this is the fall back value.
* `updated_at` - Datetime the item was updated.
