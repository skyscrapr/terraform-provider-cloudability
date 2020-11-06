# Cloudability Provider

The Cloudability provider is used to manage a Cloudability configuration defined in terraform. The supported resources are listed in this documentation. 

The provider needs to be configured with the proper credentials before it can be used.

## Example Usage

```hcl
provider cloudability {
    apikey = var.cloudability_apikey
}
```

## Argument Reference

The following arguments are supported:

* `apikey` - (Required) API Key used to authenticate with the Cloudability API. Obtained from Cloudability preferences.
