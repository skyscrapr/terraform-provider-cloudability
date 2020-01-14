module github.com/skyscrapr/terraform-provider-cloudability

go 1.13

require (
	github.com/aws/aws-sdk-go v1.25.3
	github.com/hashicorp/terraform v0.12.18
	github.com/hashicorp/terraform-plugin-sdk v1.4.1
	github.com/skyscrapr/cloudability-sdk-go v0.0.0-20200103111937-50807147ead0
)

replace github.com/skyscrapr/cloudability-sdk-go => ../cloudability-sdk-go
