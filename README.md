# Skyscrapr Terraform Provider for Cloudability

[![Build Status](https://travis-ci.com/skyscrapr/terraform-provider-cloudability.svg?branch=master)](https://travis-ci.com/skyscrapr/terraform-provider-cloudability)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12+
- [Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)
- [cloudability-sdk-go](https://github.com/skyscrapr/cloudability-sdk-go) 

## Installing the Provider

The provider is now in the offical [![community provider list for terraform](https://www.terraform.io/docs/providers/type/community-index.html)

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
