terraform {
  required_providers {
    twilio = {
      source  = "twilio.com/twilio/twilio"
      version = ">=0.2.0"
    }
  }
}

provider "twilio" {
  //  account_sid defaults to TWILIO_ACCOUNT_SID env var
  //  auth_token  defaults to TWILIO_AUTH_TOKEN env var
}

resource "twilio_autopilot_assistants_v1" "assistant" {
  unique_name       = "GreetingsBot"
  development_stage = "in-production"
}

resource "twilio_autopilot_assistants_tasks_v1" "greeting" {
  unique_name   = "greeting"
  assistant_sid = twilio_autopilot_assistants_v1.assistant.sid
  actions = jsonencode(
    {
      actions : [
        {
          collect : {
            name : "get_name",
            questions : [
              {
                name : "name",
                question : "Hi! Who do I have the pleasure of speaking to?"
              }
            ],
            on_complete : {
              redirect : "https://example.com/collect"
            }
          }
        }
      ]
  })
}

variable "greeting_samples" {
  type = list(string)
  default = [
    "hi",
    "hello",
  ]
}

resource "twilio_autopilot_assistants_tasks_samples_v1" "greeting_sample" {
  for_each      = toset(var.greeting_samples)
  assistant_sid = twilio_autopilot_assistants_v1.assistant.sid
  task_sid      = twilio_autopilot_assistants_tasks_v1.greeting.sid
  language      = "en-US"
  tagged_text   = each.value
}

resource "twilio_autopilot_assistants_model_builds_v1" "build" {
  assistant_sid = twilio_autopilot_assistants_v1.assistant.sid
  depends_on = [
    twilio_autopilot_assistants_tasks_samples_v1.greeting_sample
  ]
}

output "assistant" {
  value = twilio_autopilot_assistants_v1.assistant
}
