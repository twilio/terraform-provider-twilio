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

resource "twilio_verify_services_v2" "verify" {
  friendly_name                = "Verify"
  code_length                  = 8
  do_not_share_warning_enabled = false
  psd2_enabled                 = false
  push_include_date            = false
  lookup_enabled               = false
  skip_sms_to_landlines        = false
}

output "verification" {
  value = twilio_verify_services_v2.verify
}
