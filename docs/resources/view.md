# View Resource

This resource is used to manage views in Cloudability. 

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

resource "cloudability_view" "Amazon" {
    title = "Amazon"
    filter {
        # This prefix might need to be coded in the provider or sdk
        field = "category${cloudability_business_mapping.cloud_service_provider.id}"
		comparator = "=="
		value = "Amazon"
    }
}

resource "cloudability_view" "Azure" {
    title = "Azure"
    filter {
        field = "category${cloudability_business_mapping.cloud_service_provider.id}"
		comparator = "=="
		value = "Azure"
    }
}
```

## Argument Reference

* `title` - (Required) The name of the view as it will appear to the end users
* `shared_with_organization` - (Optional) Whether the view should be accessible to the entire organization. Defaults to True
* `filter` - (Optional) list of filter objects. If multiple filters are applied on the same dimension they are OR'd, however if they are on different dimensions they are AND'd.

## Attribute Reference

* `title` - The name of the view as it will appear to the end users
* `shared_with_organization` - Whether the view should be accessible to the entire organization. Defaults to True
* `filter` - list of filter objects. If multiple filters are applied on the same dimension they are OR'd, however if they are on different dimensions they are AND'd.
