# Skyscrapr Terraform Provider for Cloudability

[![Go Reference](https://pkg.go.dev/badge/github.com/skyscrapr/terraform-provider-cloudability.svg)](https://pkg.go.dev/github.com/skyscrapr/terraform-provider-cloudability)
[![Go Report Card](https://goreportcard.com/badge/github.com/skyscrapr/terraform-provider-cloudability)](https://goreportcard.com/report/github.com/skyscrapr/terraform-provider-cloudability)
![Github Actions Workflow](https://github.com/skyscrapr/terraform-provider-cloudability/actions/workflows/release.yml/badge.svg)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/skyscrapr/terraform-provider-cloudability)
![License](https://img.shields.io/dub/l/vibe-d.svg)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12+
- [Go](https://golang.org/doc/install) 1.22.7 (to build the provider plugin)
- [cloudability-sdk-go](https://github.com/skyscrapr/cloudability-sdk-go) 

## Installing the Provider

The provider is registered in the official [terraform registry](https://registry.terraform.io/providers/skyscrapr/cloudability/latest) 

This enables the provider to be auto-installed when you run ```terraform init```

You can also download the latest binary for your target platform from the [releases](https://github.com/skyscrapr/terraform-provider-cloudability/releases) tab.

## Building the Provider

- Clone the repo:
    ```sh
    $ mkdir -p terraform-provider-cloudability
    $ cd terraform-provider-cloudability
    $ git clone https://github.com/skyscrapr/terraform-provider-cloudability
    ```

- Build the provider: (NOTE: the install directory will allow using this provider by the current user)
    ```sh
    $ go build -o ~/.terraform.d/plugins/terraform-provider-cloudability
    ```

## Examples

Please see the [terraform-cloudabiliy-modules](https://github.com/skyscrapr/terraform-cloudability-modules) repo for example usage.

## TODO

- Fix Acceptance Tests 
- - There is an issue with the support for multiple providers in the acceptance tests. Needs revisiting.
