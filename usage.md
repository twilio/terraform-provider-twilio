# Twilio Provider

The Twilio provider is used to interact with the resources supported by Twilio.
The provider needs to be configured with the proper credentials before it can be used.

- [Configuration](#configuration)
  - [Arguments](#arguments)
  - [Specify an Edge Location](#specify-an-edge-location)
  - [Specify a Subaccount](#specify-a-subaccount)
- [Examples](#examples)
  - [Create and delete API Keys](#create-and-delete-api-keys)
  - [Buy and configure a Phone Number](#buy-and-configure-a-phone-number)
  - [Import a Twilio Phone Number](#import-a-twilio-phone-number)
  - [Define a Studio Flow](#define-a-studio-flow)
  - [Importing an existing Flex project](#importing-an-existing-flex-project)

## Configuration

### Arguments

The following arguments are supported:

- `account_sid` - (Required) Twilio Account SID. This can also be set via the `TWILIO_ACCOUNT_SID` environment variable.
- `auth_token` - (Required) Auth token for the account. This can also be set via the `TWILIO_AUTH_TOKEN` environment variable.
- `subaccount_sid` - (Optional) Twilio Subaccount SID. This can also be set via the `TWILIO_SUBACCOUNT_SID` environment variable.
- `edge` - (Optional) The [Edge](https://www.twilio.com/docs/global-infrastructure/edge-locations#public-edge-locations) location to be used by the Twilio client. This can also be set via the `TWILIO_EDGE` environment variable.
- `region` - (Optional) The [Region](https://www.twilio.com/docs/global-infrastructure/edge-locations/legacy-regions) to be used by the Twilio client. This can also be set via the `TWILIO_REGION` environment variable.

#### Example Usage

```terraform
terraform {
  required_providers {
    twilio = {
      source  = "twilio/twilio"
      version = ">=0.4.0"
    }
  }
}

# Configure the Twilio provider. Credentials can be found at www.twilio.com/console
provider "twilio" {
  //  account_sid defaults to TWILIO_ACCOUNT_SID env var
  //  auth_token  defaults to TWILIO_AUTH_TOKEN env var
}

# Create a new API key resource
resource "twilio_api_accounts_keys_v2010" "key_name" {
  friendly_name = "terraform key"
}

output "messages" {
  value = twilio_api_accounts_keys_v2010.key_name
}
```

then execute the following in your terminal to initialize and apply the changes to your Twilio infrastructure:

```bash
# Initialize a new or existing Terraform working directory
$ terraform init
# Generate a new infrastructure plan and create or update infrastructure accordingly
$ terraform apply
```

### Specify an Edge Location

You can define the [Edge](https://www.twilio.com/docs/global-infrastructure/edge-locations#public-edge-locations) and/or [Region](https://www.twilio.com/docs/global-infrastructure/edge-locations/legacy-regions) by setting the environment variables `TWILIO_EDGE` and/or `TWILIO_REGION`. However, the resource configuration in your Terraform configuration file takes precedence.

```terraform
provider "twilio" {
  //  account_sid defaults to TWILIO_ACCOUNT_SID env var
  //  auth_token  defaults to TWILIO_AUTH_TOKEN env var
  region = "au1"
  edge   = "sydney"
}
```

Setting this configuration will result in the Twilio client hostname transforming from `api.twilio.com` to `api.sydney.au1.twilio.com`.

A Twilio client constructed without these parameters will also look for `TWILIO_REGION` and `TWILIO_EDGE` variables inside the current environment.

### Specify a Subaccount

You can specify a subaccount to use with the provider by either setting the `TWILIO_SUBACCOUNT_SID` environment variable or explicitly passing it to the provider like so:

```terraform
provider "twilio" {
  // account_sid     defaults to TWILIO_ACCOUNT_SID env var
  // auth_token      defaults to TWILIO_AUTH_TOKEN env var
  // subaccount_sid  defaults to TWILIO_SUBACCOUNT_SID env var
}
```

```terraform
provider "twilio" {
  account_sid    = "AC00112233445566778899aabbccddeefe"
  auth_token     = "12345678123456781234567812345678"
  subaccount_sid = "AC00112233445566778899aabbccddeeff"
}
```

Alternatively, you can specify the subaccount to use at the resource level:

```terraform
resource "twilio_api_accounts_keys_v2010" "key_name" {
  path_account_sid = "AC00112233445566778899aabbccddeeff"
  friendly_name    = "subaccount key"
}
```

## Examples

### Create and Delete API Keys

```terraform
resource "twilio_api_accounts_keys_v2010" "key_name" {
  friendly_name = "terraform key"
}

output "messages" {
  value = twilio_api_accounts_keys_v2010.key_name
}
```

To delete a specific key in your terraform infrastructure you can use the command:

`terraform destroy -target twilio_api_keys_v2010.<resource name>`

To delete the terraform key created in this example, use:

`terraform destroy -target twilio_api_keys_v2010.key_name`

#### API Key Attributes

- `sid` - The unique string that that we created to identify the Key resource.
- `friendly_name` - The string that you assigned to describe the resource.

For more information see [the API Key documentation](https://www.twilio.com/docs/iam/keys/api-key).

### Buy and configure a Phone Number

```terraform
resource "twilio_api_incoming_phone_numbers_v2010" "phone_number" {
  area_code     = "415"
  friendly_name = "terraform phone number"
  sms_url       = "https://demo.twilio.com/welcome/sms/reply"
  voice_url     = "https://demo.twilio.com/docs/voice.xml"
}
```

### Import a Twilio Phone Number

```terraform
resource "twilio_api_accounts_incoming_phone_numbers_v2010" "imported_phone_number" {
  phone_number = "+14444444444"
}
```

### Define a Studio Flow

```terraform
resource "twilio_studio_flows_v2" "studio_flow" {
  commit_message = "first draft"
  friendly_name  = "terraform flow"
  status         = "draft"
  definition = jsonencode({
    description = "A New Flow",
    states = [
      {
        name        = "Trigger"
        type        = "trigger"
        transitions = []
        properties  = {
          offset = {
            x = 0
            y = 0
          }
        }
    }]
    initial_state = "Trigger"
    flags = {
      allow_concurrent_calls = true
    }
  })
}
```

After creating a studio flow, you can make changes to your infrastructure by changing the values in your configuration file. Run `terraform apply` to apply these changes to your infrastructure.

### Importing an existing Flex project

For guidance on how to import resources from an existing Flex project, please reference our [Flex example documentation](examples/flex/v1/README.md).

<!-- ## Proxy Service

`twilio_proxy_service` provides a Twilio Proxy Service resource.
Twilio Proxy Service is the top-level scope of all other resources in the Proxy Service REST API.

### Usage

```terraform
resource "twilio_proxy_service" "default" {
  unique_name = "Unique Proxy Service"
  chat_instance_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  callback_url = "https://example.com/proxy/callback/url"
}
```

### Argument Reference

- `unique_name` - An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. This value should not have PII.
- `default_ttl` - The default ttl value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session's last create or last Interaction. The default value of 0 indicates an unlimited Session length. You can override a Session's default TTL value by setting its ttl value.
- `callback_url` - The URL we should call when the interaction status changes.
- `geo_match_level` - Where a proxy number must be located relative to the participant identifier. Can be: `country`, `area-code`, or `extended-area-code`. The default value is `country` and more specific areas than `country` are only available in North America.
- `number_selection_behavior` - The preference for Proxy Number selection in the Service instance. Can be: `prefer-sticky` or `avoid-sticky` and the default is `prefer-sticky`. `prefer-sticky` means that we will try and select the same Proxy Number for a given participant if they have previous Sessions, but we will not fail if that Proxy Number cannot be used. `avoid-sticky` means that we will try to use different Proxy Numbers as long as that is possible within a given pool rather than try and use a previously assigned number.
- `intercept_callback_url` - The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues.
- `out_of_session_callback_url` - The URL we should call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio function) responds with valid TwiML, we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See Out-of-Session Callback Response Guide for more information.
- `chat_instance_sid` - The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship.

For more information see [the Proxy Service documentation](https://www.twilio.com/docs/proxy/api/service). -->
