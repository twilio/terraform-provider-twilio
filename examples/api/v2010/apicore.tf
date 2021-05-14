terraform {
  required_providers {
    twilio = {
      source  = "twilio.com/twilio/twilio"
      version = "0.1.0"
    }
  }
}

provider "twilio" {
  account_sid = "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  auth_token  = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}


resource "twilio_api_incoming_phone_numbers_v2010" "buy_and_configure_phone_number" {
  account_sid = "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  area_code = "415"
  friendly_name = "terraform phone number"
  sms_url = "https://demo.twilio.com/welcome/sms/reply"
  voice_url = "https://demo.twilio.com/docs/voice.xml"
}

resource "twilio_api_message_v2010" "send_message" {
  account_sid = "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  from        = "9999999999"
  to          = "4444444444"
  body        = "Hello from Twilio Terraform!"
}

resource "twilio_api_call_v2010" "make_call" {
  account_sid = "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  from        = "+19999999999"
  to          = "+14444444444"
  twiml       = "<?xml version=\"1.0\" encoding=\"UTF-8\"?><Response><Say>Hello World</Say></Response>"
}


output "messages" {
  value = twilio_api_message_v2010.send_message
}
