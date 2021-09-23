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
