# Upgrade guide

[2022-06-01] 0.17.x to 0.18.x
------------------------------

### CHANGED - Use the full resource names for AWS, SIP, and Usa2p resources.

Multiple resources were missing the full resource name when an acronym was used. These parts have been added to the
resource name.

#### Create a SIP domain

```terraform
// 0.17.x
resource "twilio_api_accounts_domains" "sip_domain" {
  domain_name = "my-domain-sip.twilio.com"
}
```

```terraform
// 0.18.x
resource "twilio_api_accounts_sip_domains" "sip_domain" {
  domain_name = "my-domain-sip.twilio.com"
}
```

[2022-05-20] 0.16.x to 0.17.x
------------------------------
### CHANGED - Renamed ApiV2010 to Api.
ApiV2010 has now been renamed to Api. This has caused a breaking change for all endpoints located under `rest/api/2010`.

#### Buy and Configure a Phone Number
```terraform
// 0.16.x
terraform {
  required_providers {
    twilio = {
      source  = "twilio/twilio"
      version = "0.16.0"
    }
  }
}

# Credentials can be found at www.twilio.com/console.
provider "twilio" {
  //  username defaults to TWILIO_API_KEY with TWILIO_ACCOUNT_SID as the fallback env var
  //  password  defaults to TWILIO_API_SECRET with TWILIO_AUTH_TOKEN as the fallback env var
}

resource "twilio_api_accounts_incoming_phone_numbers_v2010" "phone_number" {
  area_code     = "415"
  friendly_name = "terraform phone number"
  sms_url       = "https://demo.twilio.com/welcome/sms/reply"
  voice_url     = "https://demo.twilio.com/docs/voice.xml"
}

output "phone_numbers" {
  value = twilio_api_accounts_incoming_phone_numbers_v2010.phone_number
}
```
```terraform
// 0.17.x
terraform {
  required_providers {
    twilio = {
      source  = "twilio/twilio"
      version = "0.17.0"
    }
  }
}

# Credentials can be found at www.twilio.com/console.
provider "twilio" {
  //  username defaults to TWILIO_API_KEY with TWILIO_ACCOUNT_SID as the fallback env var
  //  password  defaults to TWILIO_API_SECRET with TWILIO_AUTH_TOKEN as the fallback env var
}

resource "twilio_api_accounts_incoming_phone_numbers" "phone_number" {
  area_code     = "415"
  friendly_name = "terraform phone number"
  sms_url       = "https://demo.twilio.com/welcome/sms/reply"
  voice_url     = "https://demo.twilio.com/docs/voice.xml"
}

output "phone_numbers" {
  value = twilio_api_accounts_incoming_phone_numbers.phone_number
}
```
