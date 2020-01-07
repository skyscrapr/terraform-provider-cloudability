# Terraform Provider for Cloudability

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12+
- [Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)

Setup Notes for Mac

Git Setup
Install xcode
Install xcode command line tools:
    xcode-select --install

Go Setup
Install Go

Provider template setup:
Setup provider boilerplate from terraform docs. This will create main.go and provider.go

Setup go modules
go mod init github.com/skyscrapr/terraform-provider-cloudability

## Installing for testing

go build -o ~/.terraform.d/plugins/terraform-provider-cloudability
