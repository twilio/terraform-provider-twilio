terraform {
  required_providers {
    twilio = {
      source  = "twilio/twilio"
      version = ">=0.4.0"
    }
  }
}

provider "twilio" {
  //  account_sid defaults to TWILIO_ACCOUNT_SID/TWILIO_API_KEY env var
  //  auth_token  defaults to TWILIO_AUTH_TOKEN/TWILIO_API_SECRET env var
}

resource "twilio_studio_flows_v2" "flow" {
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

output "flows" {
  value = twilio_studio_flows_v2.flow
}
