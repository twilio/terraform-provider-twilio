terraform {
  required_providers {
    twilio = {
      source  = "twilio/twilio"
      version = ">=0.4.0"
    }
  }
}

provider "twilio" {
  //  account_sid defaults to TWILIO_ACCOUNT_SID with TWILIO_API_KEY as the fallback env var
  //  auth_token  defaults to TWILIO_AUTH_TOKEN with TWILIO_API_SECRET as the fallback env var
}

resource "twilio_serverless_services_v1" "service" {
  friendly_name = "Terraform service"
  unique_name   = "uniqueName"
}

resource "twilio_serverless_services_functions_v1" "function" {
  service_sid   = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" //service_sid from new service
  friendly_name = "My serverless func"
}

output "services" {
  value = twilio_serverless_services_v1.service
}

output "functions" {
  value = twilio_serverless_services_functions_v1.function
}
