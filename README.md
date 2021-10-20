# Twilio Terraform Provider
[![Tests](https://github.com/twilio/terraform-provider-twilio/actions/workflows/test.yml/badge.svg)](https://github.com/twilio/terraform-provider-twilio/actions/workflows/test.yml)
[![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/twilio/terraform-provider-twilio?)](https://github.com/Twilio/terraform-provider-twilio/releases/latest)
[![Terraform Registry](https://img.shields.io/badge/registry-twilio-green?logo=terraform&style=flat)](https://registry.terraform.io/providers/twilio/twilio/latest)

## Project Status

:warning: This project is currently in **PILOT** and under active development. If you would like to participate in the pilot, please sign up for [Twilio Insiders](https://twil.io/insiders).

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) v0.15.x
- [Go](https://golang.org/doc/install) 1.15+ (to build the provider plugin)

## Resource Documentation

Documentation on the available resources that can be managed by this provider and their parameters can be found [here](twilio/resources/README.md).

Note that file upload resources are currently not available.

## Building The Provider

Clone repository:

```sh
git clone git@github.com:twilio/terraform-provider-twilio
```

Enter the provider directory and build the provider:

```sh
make build
```

## Installing and Using the Provider

1. Run `make install` to install and build the twilio-terraform-provider.
2. Configure the Twilio provider with your twilio credentials in your Terraform configuration file (e.g. main.tf). These can also be set via `TWILIO_ACCOUNT_SID` and `TWILIO_AUTH_TOKEN` environment variables.
3. Add your resource configurations to your Terraform configuration file (e.g. main.tf).

```terraform
terraform {
  required_providers {
    twilio = {
      source  = "twilio/twilio"
      version = "0.9.0"
    }
  }
}

# Credentials can be found at www.twilio.com/console.
provider "twilio" {
  //  username defaults to TWILIO_API_KEY with TWILIO_ACCOUNT_SID as the fallback env var
  //  password  defaults to TWILIO_API_SECRET with TWILIO_AUTH_TOKEN as the fallback env var
}

resource "twilio_api_accounts_keys_v2010" "key_name" {
  friendly_name = "terraform key"
}

output "messages" {
  value = twilio_api_accounts_keys_v2010.key_name
}
```

4. Run `terraform init` and `terraform apply` to initialize and apply changes to your Twilio infrastructure.

### Using environment variables

You can use credentials stored in environment variables for your setup:

#### OPTION 1 (recommended)
* `TWILIO_ACCOUNT_SID` = your Account SID from [your console](https://www.twilio.com/console)
* `TWILIO_API_KEY` = an API Key created in [your console](https://twil.io/get-api-key)
* `TWILIO_API_SECRET` = the secret for the API Key (you would have received this when you created an API key)
* _(optional)_ `TWILIO_REGION` = the Region for the account

#### OPTION 2
* `TWILIO_ACCOUNT_SID` = your Account SID from [your console](https://www.twilio.com/console)
* `TWILIO_AUTH_TOKEN` = your Auth Token from [your console](https://www.twilio.com/console)
* _(optional)_ `TWILIO_REGION` = the Region for the account

## Examples

For usage examples, checkout the [documentation in usage.md](usage.md) and the [examples folder](examples).

## Developing the Provider

The boilerplate includes the following:

- `Makefile` contains helper functions used to build, package and install the Twilio Terraform Provider. It's currently written for MacOS Terraform provider development, but you can change the variables at the top of the file to match your OS_ARCH.

  The `install` function is configured to install the provider into the ~/.terraform.d/plugins/ folder. To use the locally installed provider, prefix the source path with `local` in your Terraform configuration file (i.e., `source = "local/twilio/twilio"`). Then run `terraform init` to use the locally built provider.

- `examples` contains sample Terraform configuration that can be used to test the Twilio provider
- `twilio` contains the main provider code. This will be where the provider's resources and data source implementations will be defined.

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.15+ is _required_).

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
make build
...
$GOPATH/bin/terraform-provider-twilio
...
```

In order to run the full suite of Acceptance tests, run `make testacc`. Provide your Account SID and Auth Token as environment variables to properly configure the test suite.

_Note:_ Acceptance tests create real resources, and often cost money to run.

```sh
 make testacc TWILIO_ACCOUNT_SID=YOUR_ACCOUNT_SID TWILIO_AUTH_TOKEN=YOUR_AUTH_TOKEN
```

You can also specify a particular suite to run like so:

```shell
 make testacc TEST=./twilio/ TWILIO_ACCOUNT_SID=YOUR_ACCOUNT_SID TWILIO_AUTH_TOKEN=YOUR_AUTH_TOKEN
```

An example test file can be found [here](https://github.com/twilio/terraform-provider-twilio/blob/main/twilio/resources_flex_test.go).

### Debugging

First:

```sh
export TF_LOG=TRACE
```

then refer to the [Terraform Debugging Documentation](https://www.terraform.io/docs/internals/debugging.html).

### Debugging with Delve / Goland

You can build and debug the provider locally. When using Goland you can set break point and step through code:

```sh
$ dlv debug main.go -- -debug
Type 'help' for list of commands.
(dlv) c
Provider started, to attach Terraform set the TF_REATTACH_PROVIDERS env var:

	TF_REATTACH_PROVIDERS='{"registry.terraform.io/twilio/twilio":{...}}}'
```

Copy the `TF_REATTACH_PROVIDERS` and run Terraform with this value set:

```sh
$ TF_REATTACH_PROVIDERS='...' terraform init
$ TF_REATTACH_PROVIDERS='...' terraform plan
...
```

Terraform will use the binary running under `dlv` instead of the `twilio/twilio` registry version. For further details
refer to the [Terraform Debugging Providers](https://www.terraform.io/docs/extend/debugging.html) documentation.
