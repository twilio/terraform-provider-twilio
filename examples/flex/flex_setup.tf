terraform {
  required_providers {
    twilio = {
      source  = "twilio/twilio"
      version = ">=0.4.0"
    }
  }
}

provider "twilio" {
  //  username defaults to TWILIO_API_KEY with TWILIO_ACCOUNT_SID as the fallback env var
  //  password  defaults to TWILIO_API_SECRET with TWILIO_AUTH_TOKEN as the fallback env var
}

resource "twilio_chat_services_v2" "chat_service" {
  friendly_name = "Flex Chat Service"
  webhook_filters = [
    "onMessageSent",
    "onChannelDestroyed",
  "onChannelUpdated"]
  reachability_enabled = true
  read_status_enabled  = true
}

resource "twilio_studio_flows_v2" "flow" {
  commit_message = "first draft"
  friendly_name  = "terraform flow 1"
  status         = "draft"
  definition = jsonencode({
    description = "A New Flow",
    states = [
      {
        name        = "Trigger"
        type        = "trigger"
        transitions = []
        properties = {
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

resource "twilio_flex_flex_flows_v1" "flows" {
  friendly_name        = "Test Twilio flex flow"
  chat_service_sid     = twilio_chat_services_v2.chat_service.id
  channel_type         = "sms"
  integration_flow_sid = twilio_studio_flows_v2.flow.id
  contact_identity     = "true"
  enabled              = true
}

output "flex_flows" {
  value = twilio_flex_flex_flows_v1.flows
}

output "studio_flows" {
  value = twilio_studio_flows_v2.flow
}

output "chat_service" {
  value = twilio_chat_services_v2.chat_service
}
