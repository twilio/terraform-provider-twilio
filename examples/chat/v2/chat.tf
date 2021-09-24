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

# Configure chat service
resource "twilio_chat_services_v2" "chat_service" {
  friendly_name                                 = "Chat Service"
  read_status_enabled                           = false
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

# Create channels
resource "twilio_chat_services_channels_v2" "customer_channel" {
  service_sid   = twilio_chat_services_v2.chat_service.sid
  friendly_name = "Customer Chat"
}
resource "twilio_chat_services_channels_v2" "employee_channel" {
  service_sid   = twilio_chat_services_v2.chat_service.sid
  friendly_name = "Employee Chat"
}

# Create an employee role
resource "twilio_chat_services_roles_v2" "employee_role" {
  service_sid   = twilio_chat_services_v2.chat_service.sid
  friendly_name = "Employee"
  type          = "deployment"
  permission    = ["createChannel", "deleteAnyMessage", "destroyChannel", "removeMember"]
}

# Create members
resource "twilio_chat_services_users_v2" "alice_chat_member" {
  service_sid   = twilio_chat_services_v2.chat_service.sid
  identity      = "Alice"
  friendly_name = "Alice Smith"
}

resource "twilio_chat_services_users_v2" "bob_chat_member" {
  service_sid   = twilio_chat_services_v2.chat_service.sid
  role_sid      = twilio_chat_services_roles_v2.employee_role.sid
  identity      = "Bob"
  friendly_name = "Bob Williams"
}

resource "twilio_chat_services_users_v2" "lisa_chat_member" {
  service_sid   = twilio_chat_services_v2.chat_service.sid
  role_sid      = twilio_chat_services_roles_v2.employee_role.sid
  identity      = "Lisa"
  friendly_name = "Lisa Doe"
}

# Add members to Customer Chat channel
resource "twilio_chat_services_channels_members_v2" "add_alice_to_customer_channel" {
  service_sid = twilio_chat_services_v2.chat_service.sid
  channel_sid = twilio_chat_services_channels_v2.customer_channel.sid
  identity    = "Alice"
}

resource "twilio_chat_services_channels_members_v2" "add_bob_to_customer_channel" {
  service_sid = twilio_chat_services_v2.chat_service.sid
  channel_sid = twilio_chat_services_channels_v2.customer_channel.sid
  identity    = "Bob"
}

# Add members to Employee Chat channel
resource "twilio_chat_services_channels_members_v2" "add_bob_to_employee_channel" {
  service_sid = twilio_chat_services_v2.chat_service.sid
  channel_sid = twilio_chat_services_channels_v2.employee_channel.sid
  identity    = "Bob"
}

resource "twilio_chat_services_channels_members_v2" "add_lisa_to_employee_channel" {
  service_sid = twilio_chat_services_v2.chat_service.sid
  channel_sid = twilio_chat_services_channels_v2.employee_channel.sid
  identity    = "Lisa"
}

# Populate with initial chat
resource "twilio_chat_services_channels_messages_v2" "initial_message" {
  service_sid = twilio_chat_services_v2.chat_service.sid
  channel_sid = twilio_chat_services_channels_v2.customer_channel.sid
  body        = "Welcome to our service chat! How can I help you today?"
  from        = "Bob"
}
