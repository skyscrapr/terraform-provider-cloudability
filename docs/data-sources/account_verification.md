# Account Verification Datasource

This datasource tracks an account verification with Cloudability.

## Example Usage

```hcl
data "cloudability_account_verification" "aws_payer_account" {
    vendor_account_id = var.aws_payer_account_id
}
```

## Argument Reference

* `vendor_account_id` - (Required) 12 digit string corresponding to your AWS account ID
* `vendor_key` - (Optional) Hardcoded to "aws"
* `retry_count` - (Optional) Number of times to retry the verification. Default 20
* `retry_wait` - (Optional) Number of seconds to wait between verification retries. Default 5

## Attribute Reference

* `vendor_account_id` - 12 digit string corresponding to your AWS account ID
* `vendor_key` - Hardcoded to "aws"
* `retry_count` - Number of times to retry the verification
* `retry_wait` - Number of seconds to wait between verification retries
* `state` - Examples: unverified, verified, error
* `last_verification_attempted_at` - Date timestamp, example: 1970-01-01T00:00:00.000Z
* `message` - Error message for credentials in error state
* `external_id` - The external ID used to prevent confused deputies. Generated by Cloudability
