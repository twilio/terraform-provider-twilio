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
  //  subaccount_sid  defaults to TWILIO_SUBACCOUNT_SID env var
}


resource "twilio_api_incoming_phone_numbers_v2010" "phone_number" {
  area_code     = "415"
  friendly_name = "terraform phone number"
  sms_url       = "https://demo.twilio.com/welcome/sms/reply"
  voice_url     = "https://demo.twilio.com/docs/voice.xml"
}


resource "twilio_api_message_v2010" "message" {
  path_account_sid = "AC123456123456123456123456123456"
  from             = "9999999999"
  to               = "4444444444"
  body             = "Hello from Twilio Terraform!"
}

resource "twilio_api_call_v2010" "call" {
  from  = "9999999999"
  to    = "4444444444"
  twiml = "<?xml version=\"1.0\" encoding=\"UTF-8\"?><Response><Say>Hello World</Say></Response>"
}

output "phone_numbers" {
  value = twilio_api_incoming_phone_numbers_v2010.phone_number
}

output "messages" {
  value = twilio_api_message_v2010.message
}

output "calls" {
  value = twilio_api_call_v2010.call
}

