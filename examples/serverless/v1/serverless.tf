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

resource "twilio_serverless_service_v1" "create_new_service" {
  friendly_name = "Terraform service"
  unique_name = "uniqueName"
}

resource "twilio_serverless_function_v1" "create_serverless_function" {
  service_sid   = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" //service_sid from new service
  friendly_name = "My serverless 1"
}
