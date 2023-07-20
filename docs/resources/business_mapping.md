# Business Mapping Resource

Business Mapping is an activity that allows you to map your cloud cost and usage to custom dimensions and custom metrics that are important to report on within your organization.
This resource maps to BUSINESS_DIMENSION Business Mapping.

## Example Usage

```hcl
resource "cloudability_business_mapping" "cloud_service_provider" {
    name = "Cloud Service Provider"
    default_value = "Unknown Cloud Service Provider"
    kind = "BUSINESS_DIMENSION"
    statement {
        match_expression = "DIMENSION['vendor'] == 'Amazon'"
        value_expression = "'Amazon'"
    }
    statement {
        match_expression = "DIMENSION['vendor'] == 'Azure'"
        value_expression = "'Azure'"
    }
}
```

## Argument Reference

* `name` - (Required) Name for the Business Dimension/Metric.
* `kind` - (Optional) Should be hardcoded to BUSINESS_DIMENSION (BUSINESS_METRIC not supported)
* `default_value` - (Optional) If no rule matches then take the value resolved by this expression as the fall back value.
* `statement` - (Optional) List of statements

## Attribute Reference

* `index` - An integer value between 1 and 10 representing the ID for the Business Dimension.
* `kind` - Should be hardcoded in the provider resource to BUSINESS_DIMENSION
* `name` - Name for the Business Dimension
* `default_value` - If no rule matches then take the value resolved by this expression as the fall back value.
* `statement` -  List of statement objects.
* `updated_at` - Datetime the item was updated.
