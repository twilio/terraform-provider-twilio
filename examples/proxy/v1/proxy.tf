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
}

# Step 1: Purchase 2 phone numbers -- these are going to be the participants
resource "twilio_api_accounts_incoming_phone_numbers_v2010" "alice_phone" {
  area_code     = "415"
  friendly_name = "Alice's Phone"
}

resource "twilio_api_accounts_incoming_phone_numbers_v2010" "bob_phone" {
  area_code     = "908"
  friendly_name = "Bob's Phone"
}

# Create some proxy phone numbers to be added into the service/session pool
resource "twilio_api_accounts_incoming_phone_numbers_v2010" "proxy_phone_dev_1" {
  area_code     = "908"
  friendly_name = "Proxy Phone Dev 1"
}

resource "twilio_api_accounts_incoming_phone_numbers_v2010" "proxy_phone_dev_2" {
  area_code     = "908"
  friendly_name = "Proxy Phone Dev 2"
}

resource "twilio_api_accounts_incoming_phone_numbers_v2010" "proxy_phone_stage_1" {
  area_code     = "415"
  friendly_name = "Proxy Phone Stage 1"
}

resource "twilio_api_accounts_incoming_phone_numbers_v2010" "proxy_phone_stage_2" {
  area_code     = "415"
  friendly_name = "Proxy Phone Stage 2"
}

# Create chat services
resource "twilio_chat_services_v2" "chat_service_dev" {
  friendly_name                                 = "Chat Dev"
  read_status_enabled                           = true
  reachability_enabled                          = true
  typing_indicator_timeout                      = 60
  notifications_new_message_enabled             = true
  notifications_new_message_badge_count_enabled = true
  notifications_added_to_channel_enabled        = true
  notifications_added_to_channel_template       = true
  media_compatibility_message                   = "Media message has no text!"
  limits_user_channels                          = 50
  limits_channel_members                        = 50
  webhook_filters                               = ["onMessageSend", "onMessageSent"]
}

resource "twilio_chat_services_v2" "chat_service_stage" {
  friendly_name                                 = "Chat Stage"
  read_status_enabled                           = true
  reachability_enabled                          = true
  typing_indicator_timeout                      = 60
  notifications_new_message_enabled             = true
  notifications_new_message_badge_count_enabled = true
  notifications_added_to_channel_enabled        = true
  notifications_added_to_channel_template       = true
  media_compatibility_message                   = "Media message has no text!"
  limits_user_channels                          = 50
  limits_channel_members                        = 50
  webhook_filters                               = ["onMessageSend", "onMessageSent"]
}


# Create 2 proxy services for both
resource "twilio_proxy_services_v1" "dev_proxy_service" {
  unique_name               = "Dev Proxy Service"
  chat_instance_sid         = twilio_chat_services_v2.chat_service_dev.sid
  number_selection_behavior = "prefer-sticky"
  geo_match_level           = "country"
  default_ttl               = 3600
}

resource "twilio_proxy_services_v1" "stage_proxy_service" {
  unique_name               = "Stage Proxy Service"
  chat_instance_sid         = twilio_chat_services_v2.chat_service_stage.sid
  number_selection_behavior = "avoid-sticky"
  geo_match_level           = "country"
  default_ttl               = 3600
}

# Add the proxy phone numbers to the service
resource "twilio_proxy_services_phone_numbers_v1" "proxy_phone_dev_service_1" {
  service_sid  = twilio_proxy_services_v1.dev_proxy_service.sid
  phone_number = twilio_api_accounts_incoming_phone_numbers_v2010.proxy_phone_dev_1.phone_number
  is_reserved  = false
}

resource "twilio_proxy_services_phone_numbers_v1" "proxy_phone_dev_service_2" {
  service_sid  = twilio_proxy_services_v1.dev_proxy_service.sid
  phone_number = twilio_api_accounts_incoming_phone_numbers_v2010.proxy_phone_dev_2.phone_number
  is_reserved  = false
}

resource "twilio_proxy_services_phone_numbers_v1" "proxy_phone_stage_service_1" {
  service_sid  = twilio_proxy_services_v1.stage_proxy_service.sid
  phone_number = twilio_api_accounts_incoming_phone_numbers_v2010.proxy_phone_stage_1.phone_number
  is_reserved  = false
}

resource "twilio_proxy_services_phone_numbers_v1" "proxy_phone_stage_service_2" {
  service_sid  = twilio_proxy_services_v1.stage_proxy_service.sid
  phone_number = twilio_api_accounts_incoming_phone_numbers_v2010.proxy_phone_stage_2.phone_number
  is_reserved  = false
}

# Create a session and add Alice and Bob as participants so they can talk to each other via the proxy phone numbers
resource "twilio_proxy_services_sessions_v1" "dev_session" {
  service_sid = twilio_proxy_services_v1.dev_proxy_service.sid
  unique_name = "Alice Proxy Session"
  depends_on = [
    twilio_proxy_services_phone_numbers_v1.proxy_phone_dev_service_1,
    twilio_proxy_services_phone_numbers_v1.proxy_phone_dev_service_2
  ]
  participants = [
    jsonencode({
      friendly_name : "Alice"
      identifier : twilio_api_accounts_incoming_phone_numbers_v2010.alice_phone.phone_number
    }),
    jsonencode({
      friendly_name : "Bob"
      identifier : twilio_api_accounts_incoming_phone_numbers_v2010.bob_phone.phone_number
    }),
  ]
}

resource "twilio_proxy_services_sessions_v1" "stage_session" {
  service_sid = twilio_proxy_services_v1.stage_proxy_service.sid
  unique_name = "Bob Proxy Session"
  depends_on = [
    twilio_proxy_services_phone_numbers_v1.proxy_phone_stage_service_1,
    twilio_proxy_services_phone_numbers_v1.proxy_phone_stage_service_2
  ]
  participants = [
    jsonencode({
      friendly_name : "Alice"
      identifier : twilio_api_accounts_incoming_phone_numbers_v2010.alice_phone.phone_number
    }),
    jsonencode({
      friendly_name : "Bob"
      identifier : twilio_api_accounts_incoming_phone_numbers_v2010.bob_phone.phone_number
    }),
  ]
}

output "dev_service" {
  value = twilio_proxy_services_v1.dev_proxy_service
}

output "stage_service" {
  value = twilio_proxy_services_v1.stage_proxy_service
}
