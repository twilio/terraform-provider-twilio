# terraform-provider-twilio
Terraform Twilio provider

## Building

```bash
 go build -o terraform-provider-twilio
```

## Running

In `main.tf`:
```hcl-terraform
provider "twilio" {
  account_sid = "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  auth_token  = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
resource "twilio_chat_service" "chat-service" {
  friendly_name        = "Terraform'd Chat Service"
  reachability_enabled = false
}
resource "twilio_proxy_service" "proxy-service" {
  unique_name     = "Terraform'd Proxy Service"
  default_ttl     = 0
}
```
then:
```bash
 terraform init
 terraform apply
 terraform destroy
```

## Debugging
First:
```bash
export TF_LOG=TRACE
```
then refer to the [Terraform Debugging Documentation](https://www.terraform.io/docs/internals/debugging.html).