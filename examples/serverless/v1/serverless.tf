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

resource "twilio_serverless_v1" "default" {
  service_sid   = "ZS97e9bafac602e0c2ce4046f14c7e8c1c"
  friendly_name = "My serverless 1"
}
