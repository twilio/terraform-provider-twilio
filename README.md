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

data "twilio_available_phone_numbers_local" "local_numbers" {
}

resource "twilio_incoming_phone_number" "incoming_phone_number" {
  phone_number = data.twilio_available_phone_numbers_local.local_numbers.available_phone_numbers[0].phone_number
}

resource "twilio_chat_service" "chat-service" {
  friendly_name        = "Terraform'd Chat Service"
  reachability_enabled = false
}

resource "twilio_proxy_service" "proxy-service" {
  unique_name     = "Terraform'd Proxy Service"
  default_ttl     = 0
}

resource "twilio_workspace" "my-workspace" {
friendly_name = "Flex Task Assignment"
}
resource "twilio_taskqueue" "my-taskqueue" {
friendly_name = "Everyone"
workspace_sid = twilio_workspace.my-workspace.id
}
resource "twilio_workflow" "my-workflow" {
friendly_name = "Assign to Anyone"
configuration = jsonencode(
            {
               task_routing = {
                   default_filter = {
                       task_queue_sid = twilio_taskqueue.my-taskqueue.id
                    }
                }
            }
        )
workspace_sid = twilio_workspace.my-workspace.id
depends_on = [
        twilio_taskqueue.my-taskqueue,
   ]
}
resource "twilio_activity" "my-activity" {
friendly_name = "Break"
workspace_sid = twilio_workspace.my-workspace.id
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