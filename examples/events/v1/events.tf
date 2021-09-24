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

# Step 1: Create an Event Sink
# There are two types of Sinks you can create: Kinesis and Webhook

# Type 1: Kinesis Sink
resource "twilio_events_sinks_v1" "kinesis" {
  description = "Kinesis Sink"
  sink_configuration = jsonencode({
    arn : "arn:aws:kinesis:us-east-1:111111111:stream/test",
    role_arn : "arn:aws:iam::111111111:role/Role",
    external_id : "1234567890"
  })
  sink_type = "kinesis"
}

# Type 2: Webhook Sink
resource "twilio_events_sinks_v1" "webhook" {
  description = "Webhook Sink"
  sink_configuration = jsonencode({
    destination : "http://example.org/webhook",
    method : "POST",
    batch_events : false
  })
  sink_type = "webhook"
}

# Step 2: Subscribe to Twilio Events
# For simplicity, this example will continue with our Webhook sink
resource "twilio_events_subscriptions_v1" "message_status" {
  description = "Message Status Events Subscription"
  sink_sid    = twilio_events_sinks_v1.webhook.sid
  types = [jsonencode({
    type = "com.twilio.messaging.message.queued"
  })]
}

resource "twilio_events_subscriptions_v1" "call_summary" {
  description = "Call Summary Events Subscription"
  sink_sid    = twilio_events_sinks_v1.webhook.sid
  types = [jsonencode({
    type = "com.twilio.voice.insights.call-summary.partial"
  })]
}

variable "message_status_event_types" {
  type = list(string)
  default = [
    "com.twilio.messaging.message.sent",
    "com.twilio.messaging.message.delivered",
    "com.twilio.messaging.message.failed",
    "com.twilio.messaging.message.undelivered",
    "com.twilio.messaging.message.read",
  ]
}

resource "twilio_events_subscriptions_subscribed_events_v1" "message_status_event" {
  subscription_sid = twilio_events_subscriptions_v1.message_status.sid
  for_each         = toset(var.message_status_event_types)
  type             = each.key
}

variable "call_summary_event_types" {
  type = list(string)
  default = [
    "com.twilio.voice.insights.call-summary.predicted-complete",
    "com.twilio.voice.insights.call-summary.complete",
  ]
}

resource "twilio_events_subscriptions_subscribed_events_v1" "call_summary_event" {
  subscription_sid = twilio_events_subscriptions_v1.call_summary.sid
  for_each         = toset(var.call_summary_event_types)
  type             = each.key
}

output "sink" {
  value = twilio_events_sinks_v1.webhook
}

output "message_subscription" {
  value = twilio_events_subscriptions_v1.message_status
}

output "call_subscription" {
  value = twilio_events_subscriptions_v1.call_summary
}
