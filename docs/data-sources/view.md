# Views Datasource

This datasource tracks a View in Cloudability.

## Example Usage

```hcl
data "cloudability_view" "test" {
    id = "12345"
}
```

## Argument Reference

- `title` - (Required) The name of the view as it will appear to the end users
- `shared_with_users` - (Optional) The discrete list of users (by their unique identifier) that the view should be shared with
- `shared_with_organization` - (Required) Whether the view should be accessible to the entire organization
- `filters` - (Required) list of filter objects. If multiple filters are applied on the same dimension they are OR'd, however if they are on different dimensions they are AND'd. See below regarding filter specifics.
- - `field`
- - `comparator`
- - `value`

## Attribute Reference

- `owner_id` - (Optional) Unique identifier for the user who created the view