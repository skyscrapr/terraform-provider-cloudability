# User Resource

This resource is used to manage users in Cloudability

## Example Usage

```hcl
resource "cloudability_user" "test_user" {
    email = "test.user@test.com.test"
    full_name = "Test User"
    role = "Administrator"
}
```

## Argument Reference

* `email` - (Optional) The email address of the user
* `full_name` - (Required) The full name of the user
* `role` - (Optional) Role assigned to the user: [User|Administrator] Defaults to User
* `restricted` - (Required) True if the user is allowed to have no filter set applied, false if they must always have a filter set applied
* `default_dimension_filter_set_id` - (Optional) Filter set id used by default for the user
* `shared_dimension_filter_set_ids` - Not supported yet

## Attribute Reference

* `email` - The email address of the user
* `full_name` - The full name of the user
* `role` - Role assigned to the user: [User|Administrator]
* `restricted` - True if the user is allowed to have no filter set applied, false if they must always have a filter set applied
* `default_dimension_filter_set_id` - Filter set id used by default for the user
* `shared_dimension_filter_set_ids` - Not supported yet
