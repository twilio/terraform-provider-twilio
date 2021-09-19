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


resource "twilio_bulkexports_exports_jobs_v1" "job" {
  resource_type = "Messages"
  end_day       = "2021-09-09"
  friendly_name = "Bulk test"
  start_day     = "2021-04-04"
}


output "jobs" {
  value = twilio_bulkexports_exports_jobs_v1.job
}

