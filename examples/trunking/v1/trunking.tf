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

resource "twilio_trunking_trunks_v1" "trunk" {
  friendly_name            = "ElasticTrunk"
  disaster_recovery_method = "GET"
  disaster_recovery_url    = "https://www.twilio.com/docs/sip-trunking/api/trunk-resource"
  cnam_lookup_enabled      = true
  secure                   = true
  transfer_mode            = "enable-all"
}

# can only associate original url and not Credential lists/ IP Access Control Lists or Phone Numbers since they do not support an Update
resource "twilio_trunking_trunks_origination_urls_v1" "origination_urls" {
  trunk_sid     = twilio_trunking_trunks_v1.trunk.sid
  enabled       = true
  friendly_name = "TrunkOriginationUrl"
  priority      = 50
  sip_url       = "sip:username@host"
  weight        = 50
}

output "trunk" {
  value = twilio_trunking_trunks_v1.trunk
}

output "origination_url" {
  value = twilio_trunking_trunks_origination_urls_v1.origination_urls
}
