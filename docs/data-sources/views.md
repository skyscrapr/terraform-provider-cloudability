# Views Datasource

This datasource tracks all the views in Cloudability.

## Example Usage

```hcl
data "cloudability_views" "test" {}
```

## Argument Reference

NONE

## Attribute Reference

- `view_id` - Unique identifier for the View object
- `title` - The name of the view as it will appear to the end users
- `shared_with_users` - The discrete list of users (by their unique identifier) that the view should be shared with
- `shared_with_organization` - Whether the view should be accessible to the entire organization
- `owner_id` - Unique identifier for the user who created the view
- `filters` - list of filter objects. If multiple filters are applied on the same dimension they are OR'd, however if they are on different dimensions they are AND'd. See below regarding filter specifics.
- - `field`
- - `comparator`
- - `value`
