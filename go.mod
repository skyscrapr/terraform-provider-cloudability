module github.com/skyscrapr/terraform-provider-cloudability

go 1.16

// replace github.com/skyscrapr/cloudability-sdk-go => ../cloudability-sdk-go

// replace github.com/terraform-providers/terraform-provider-aws => ../terraform-provider-aws

require (
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.26.1
	github.com/hashicorp/terraform-plugin-testing v1.2.0
	github.com/hashicorp/yamux v0.0.0-20210826001029-26ff87cf9493 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/skyscrapr/cloudability-sdk-go v0.0.6
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20210921142501-181ce0d877f6 // indirect
)
