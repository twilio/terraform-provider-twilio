# Create a SMS conversation between two people
terraform {
  required_providers {
    twilio = {
      source  = "twilio/twilio"
      version = ">=0.4.0"
    }
  }
}

# Credentials can be found at www.twilio.com/console.
provider "twilio" {
  // account_sid = "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" // removing account_sid defaults to TWILIO_ACCOUNT_SID env var
  // auth_token  = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" // removing auth_token defaults to TWILIO_AUTH_TOKEN env var
}

# Create the conversation
resource "twilio_conversations_conversations_v1" "create_conversation" {
  friendly_name = "Deep Conversation 00"
}

# Add a SMS Participant
resource "twilio_conversations_conversations_participants_v1" "add_user_00_to_conversation" {
  conversation_sid = twilio_conversations_conversations_v1.create_conversation.sid
  messaging_binding_address = "<User 00 Personal Mobile Number>"
  messaging_binding_proxy_address = "<Your purchased Twilio Phone Number>"
}

# Add a SMS Participant
resource "twilio_conversations_conversations_participants_v1" "add_user_01_to_conversation" {
  messaging_binding_address = "<User 01 Personal Mobile Number>"
  messaging_binding_proxy_address = "<Your purchased Twilio Phone Number>"
}

output "create_conversation" {
  value = twilio_conversations_conversations_v1.create_conversation
}

output "add_user_00_to_conversation" {
  value = twilio_conversations_conversations_participants_v1.add_user_00_to_conversation
}

output "add_user_01_to_conversation" {
  value = twilio_conversations_conversations_participants_v1.add_user_01_to_conversation
}

# To add a Chat participant: https://www.twilio.com/docs/conversations/quickstart#configure-the-conversations-demo-application-using-codesandboxio