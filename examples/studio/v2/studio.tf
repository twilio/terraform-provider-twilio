terraform {
  required_providers {
    twilio = {
      source  = "twilio.com/twilio/twilio"
      version = "0.1.0"
    }
  }
}

provider "twilio" {
  //  account_sid defaults to TWILIO_ACCOUNT_SID env var
  //  auth_token  defaults to TWILIO_AUTH_TOKEN env var
}

resource "twilio_studio_flow_v2" "flow" {
  commit_message = "first draft"
  friendly_name  = "terraform flow"
  status         = "draft"
  definition     = "{\"description\": \"A New Flow\", \"states\": [{\"name\": \"Trigger\", \"type\": \"trigger\", \"transitions\": [], \"properties\": {\"offset\": {\"x\": 0, \"y\": 0}}}], \"initial_state\": \"Trigger\", \"flags\": {\"allow_concurrent_calls\": true}}"
}


output "flows" {
  value = twilio_studio_flow_v2.flow
}

