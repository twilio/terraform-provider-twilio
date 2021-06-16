terraform {
  required_providers {
    twilio = {
      source  = "twilio.com/twilio/twilio"
      version = "0.1.2"
    }
  }
}

provider "twilio" {
  //  account_sid defaults to TWILIO_ACCOUNT_SID env var
  //  auth_token  defaults to TWILIO_AUTH_TOKEN env var
}

resource "twilio_serverless_service_v1" "service" {
  friendly_name = "Terraform service"
  unique_name   = "uniqueName"
}

resource "twilio_serverless_function_v1" "function" {
  service_sid   = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" //service_sid from new service
  friendly_name = "My serverless func"
}

output "services" {
  value = twilio_serverless_service_v1.service
}

output "functions" {
  value = twilio_serverless_function_v1.function
}
