terraform {
  required_providers {
    twilio = {
      source  = "twilio/twilio"
      version = ">=0.4.0"
    }
  }
}

provider "twilio" {
  //  account_sid defaults to TWILIO_ACCOUNT_SID env var
  //  auth_token  defaults to TWILIO_AUTH_TOKEN env var
}

resource "twilio_notify_services_v1" "service" {
}

resource "twilio_notify_services_bindings_v1" "binding" {
  service_sid  = twilio_notify_services_v1.service.sid
  address      = "9999999999"
  binding_type = "apn"
  identity     = "test"
}


output "services" {
  value = twilio_notify_services_v1.service
}

output "jobs" {
  value = twilio_notify_services_bindings_v1.binding
}

