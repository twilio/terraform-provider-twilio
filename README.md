# Twilio Terraform Provider

**:warning: Note: Issues are currently closed on this repo as we are not quite ready for feedback. Thanks!**

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) v0.15.x
-	[Go](https://golang.org/doc/install) 1.15+ (to build the provider plugin)

## Building The Provider

Clone repository:

```sh
$ git clone git@github.com:twilio/terraform-provider-twilio
```

Enter the provider directory and build the provider

```sh
$ make build
```

## Using the Provider
1. Import the resource you want to use into [twilio/provider.go](twilio/provider.go) and add the resource to the `ResourcesMap`.
```go
//import the v2010 resource to the provider into twilio/provider.go
import apiV2010 "github.com/twilio/terraform-provider-twilio/twilio/resources/api/v2010"

//add the resource to the ResourceMap
ResourcesMap: map[string]*schema.Resource{
"twilio_api_keys_v2010": apiV2010.ResourceAccountsSigningKeys(),
},
```
2. Run `make install` to install and build the twilio-terraform-provider.
3. Configure the Twilio provider with your twilio credentials in your terraform file.
4. Add your resource configurations to your terraform file.
```terraform
# Credentials can be found in the twilio console.
provider "twilio" {
  account_sid = "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  auth_token  = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}

resource "twilio_api_keys_v2010" "key_name" {
    account_sid = "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    friendly_name = "terraform key"
}

output "messages" {
    value = twilio_api_keys_v2010.key_name
}
```
5. Run `terraform init` and `terraform apply`to initialize and apply changes to your twilio infrastructure.

### Create and Delete API Keys
```terraform
resource "twilio_api_keys_v2010" "key_name" {
  account_sid = "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  friendly_name = "terraform key"
}

output "messages" {
    value = twilio_api_keys_v2010.key_name
}
```
To delete a specific key in your terraform configuration you can use the command `terraform destroy -target twilio_api_keys_v2010.<resource name>`. To delete the terraform key created in this example use `terraform destroy -target twilio_api_keys_v2010.key_name`.

### Buy and Configure a Phone Number
```terraform
resource "twilio_api_incoming_phone_numbers_v2010" "buy_phone_number" {
    account_sid = "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    area_code = "415"
    friendly_name = "terraform phone number"
    sms_url = "https://demo.twilio.com/welcome/sms/reply"
    voice_url = "https://demo.twilio.com/docs/voice.xml"
}
```
### Import a Twilio Phone Number
```terraform
resource "twilio_api_incoming_phone_numbers_v2010" "import_purchased_number" {
    account_sid = "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    phone_number = "+14444444444"
}
```
### Define a Studio Flow
```terraform
resource "twilio_studio_flow_v2" "studio_flow" {
    commit_message = "first draft"
    friendly_name = "terraform flow"
    status = "draft"
    definition = "{\"description\": \"A New Flow\", \"states\": [{\"name\": \"Trigger\", \"type\": \"trigger\", \"transitions\": [], \"properties\": {\"offset\": {\"x\": 0, \"y\": 0}}}], \"initial_state\": \"Trigger\", \"flags\": {\"allow_concurrent_calls\": true}}"
}
```
For more examples checkout the [documentation in the usage.md](usage.md) and the [examples folder](examples).

## Developing the Provider

The boilerplate includes the following:
- `Makefile` contains helper functions used to build, package and install the Twilio Terraform Provider. It's currently written for MacOS Terraform provider development, but you can change the variables at the top of the file to match your OS_ARCH.

  The `install` function is configured to install the provider into the ~/.terraform.d/plugins/ folder.
- `examples` contains sample Terraform configuration that can be used to test the Twilio provider
- `twilio` contains the main provider code. This will be where the provider's resources and data source implementations will be defined.

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.15+ is *required*).

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-twilio
...
```

In order to run the full suite of Acceptance tests, run `make testacc`. Provide your Account SID and Auth Token as environment variables to properly configure the test suite.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$  make testacc ACCOUNT_SID=YOUR_ACCOUNT_SID AUTH_TOKEN=YOUR_AUTH_TOKEN
```

An example test file can be found [here](https://github.com/twilio/terraform-provider-twilio/blob/master/twilio/resource_taskrouter_workspace_test.go).

### Debugging
First:
```sh
$ export TF_LOG=TRACE
```
then refer to the [Terraform Debugging Documentation](https://www.terraform.io/docs/internals/debugging.html).
