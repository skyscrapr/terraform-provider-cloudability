# Terraform Provider for Cloudability

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12+
- [Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)
- [cloudability-sdk-go](https://github.com/skyscrapr/cloudability-sdk-go) 

## Installing the Provider

Download the latest binary for your target platform from the [releases](https://github.com/skyscrapr/terraform-provider-cloudability/releases) tab.

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

## TODO

- Acceptance Tests
- Makefile
- Examples (perhaps link to the modules repo)
- Travis CI (how to manage secrets)