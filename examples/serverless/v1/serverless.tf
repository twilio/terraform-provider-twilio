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

resource "twilio_serverless_services_v1" "service" {
  friendly_name = "Terraform service"
  unique_name   = "uniqueName"
}

resource "twilio_serverless_services_functions_v1" "function" {
  service_sid   = twilio_serverless_services_v1.service.sid //service_sid from new service
  friendly_name = "My serverless func"
}

output "services" {
  value = twilio_serverless_services_v1.service
}

output "functions" {
  value = twilio_serverless_services_functions_v1.function
}
