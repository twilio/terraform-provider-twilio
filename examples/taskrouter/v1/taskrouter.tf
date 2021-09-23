terraform {
  required_providers {
    twilio = {
      source  = "twilio/twilio"
      version = ">=0.4.0"
    }
  }
}

provider "twilio" {
  //  account_sid defaults to TWILIO_ACCOUNT_SID with TWILIO_API_KEY as the fallback env var
  //  auth_token  defaults to TWILIO_AUTH_TOKEN with TWILIO_API_SECRET as the fallback env var
}

resource "twilio_taskrouter_workspaces_v1" "demo" {
  friendly_name = "demoWorkspace"
}

resource "twilio_taskrouter_workspaces_activities_v1" "offline" {
  workspace_sid = twilio_taskrouter_workspaces_v1.demo.sid
  friendly_name = "OfflineCustom"
  available     = false
}

resource "twilio_taskrouter_workspaces_activities_v1" "idle" {
  workspace_sid = twilio_taskrouter_workspaces_v1.demo.sid
  friendly_name = "Idle"
  available     = true
}

resource "twilio_taskrouter_workspaces_activities_v1" "busy" {
  workspace_sid = twilio_taskrouter_workspaces_v1.demo.sid
  friendly_name = "Busy"
  available     = false
}

resource "twilio_taskrouter_workspaces_activities_v1" "reserved" {
  workspace_sid = twilio_taskrouter_workspaces_v1.demo.sid
  friendly_name = "Reserved"
  available     = false
}

resource "twilio_taskrouter_workspaces_workers_v1" "alice" {
  workspace_sid = twilio_taskrouter_workspaces_v1.demo.sid
  friendly_name = "Alice"
  activity_sid  = twilio_taskrouter_workspaces_activities_v1.offline.sid
  attributes = jsonencode({
    "languages" : [
      "en",
      "es"
    ]
  })
}

resource "twilio_taskrouter_workspaces_workers_v1" "bob" {
  workspace_sid = twilio_taskrouter_workspaces_v1.demo.sid
  friendly_name = "Bob"
  activity_sid  = twilio_taskrouter_workspaces_activities_v1.offline.sid
  attributes = jsonencode({
    "languages" : [
      "en"
    ]
  })
}

resource "twilio_taskrouter_workspaces_task_queues_v1" "spanish" {
  workspace_sid  = twilio_taskrouter_workspaces_v1.demo.sid
  friendly_name  = "Spanish"
  target_workers = "languages HAS \"es\""
}

resource "twilio_taskrouter_workspaces_task_queues_v1" "english" {
  workspace_sid  = twilio_taskrouter_workspaces_v1.demo.sid
  friendly_name  = "English"
  target_workers = "languages HAS \"en\""
}

resource "twilio_taskrouter_workspaces_workflows_v1" "flow" {
  workspace_sid = twilio_taskrouter_workspaces_v1.demo.sid
  friendly_name = "Flow"
  configuration = jsonencode({
    task_routing : {
      filters : [
        {
          filter_friendly_name : "Language - Spanish",
          expression : "selected_language == \"es\"",
          targets : [
            {
              queue : twilio_taskrouter_workspaces_task_queues_v1.spanish.sid
            }
          ]
        },
        {
          filter_friendly_name : "Language - English",
          expression : "selected_language == \"en\"",
          targets : [
            {
              queue : twilio_taskrouter_workspaces_task_queues_v1.english.sid
            }
          ]
        },
      ]
    }
  })
}
