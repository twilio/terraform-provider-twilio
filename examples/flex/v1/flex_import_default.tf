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

resource "twilio_chat_services_v2" "flex_chat_service" {}

resource "twilio_taskrouter_workspaces_v1" "flex_task_assignment" {}

resource "twilio_taskrouter_workspaces_task_queues_v1" "everyone" {
  workspace_sid = twilio_taskrouter_workspaces_v1.flex_task_assignment.sid
}

resource "twilio_taskrouter_workspaces_workflows_v1" "assign_to_anyone" {
  workspace_sid = twilio_taskrouter_workspaces_v1.flex_task_assignment.sid
}

resource "twilio_taskrouter_workspaces_activities_v1" "offline" {
  workspace_sid = twilio_taskrouter_workspaces_v1.flex_task_assignment.sid
}

resource "twilio_taskrouter_workspaces_activities_v1" "available" {
  workspace_sid = twilio_taskrouter_workspaces_v1.flex_task_assignment.sid
}

resource "twilio_taskrouter_workspaces_activities_v1" "unavailable" {
  workspace_sid = twilio_taskrouter_workspaces_v1.flex_task_assignment.sid
}

resource "twilio_taskrouter_workspaces_activities_v1" "break" {
  workspace_sid = twilio_taskrouter_workspaces_v1.flex_task_assignment.sid
}
