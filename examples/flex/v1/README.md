# Configuring Flex with Terraform

## Importing an existing Flex project

To import existing Flex resources into terraform, we need to know the SIDs of the default resources that are created for you when a new Flex instance is created. Start by exporting your Flex configuration using the Twilio CLI:
```bash
twilio api:flex:v1:configuration:fetch \
  --properties=chatServiceInstanceSid,taskrouterWorkspaceSid,taskrouterTargetTaskqueueSid,taskrouterTargetWorkflowSid,taskrouterOfflineActivitySid
```

Next, we need to add placeholders for these resources in our Twilio Terraform configuration file
```terraform
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

```

Taskrouter also creates some other default Activities. We'll need to import these as well to our configuration:
```terraform
resource "twilio_taskrouter_workspaces_activities_v1" "available" {
    workspace_sid = twilio_taskrouter_workspaces_v1.flex_task_assignment.sid
}

resource "twilio_taskrouter_workspaces_activities_v1" "unavailable" {
    workspace_sid = twilio_taskrouter_workspaces_v1.flex_task_assignment.sid
}

resource "twilio_taskrouter_workspaces_activities_v1" "break" {
    workspace_sid = twilio_taskrouter_workspaces_v1.flex_task_assignment.sid
}
```

We can use the Twilio CLI to obtain the SIDs for these resources:
```bash
twilio api:taskrouter:v1:workspaces:activities:list --workspace-sid=<WORKSPACE_SID>
```
Note that the `<WORKSPACE_SID>` used in this command is the "Taskrouter Workspace SID" value obtained from the CLI call above.

Once you have configurations for each of your existing resources, you'll need to import each resource using `terraform import`, mapping each one its corresponding SID CLI output:
```bash
terraform import twilio_chat_services_v2.flex_chat_service <CHAT_SERVICE_INSTANCE_SID>
terraform import twilio_taskrouter_workspaces_v1.flex_task_assignment <TASKROUTER_WORKSPACE_SID>
terraform import twilio_taskrouter_workspaces_task_queues_v1.everyone <TASKROUTER_WORKSPACE_SID>/<TASKROUTER_TARGET_TASK_QUEUE_SID>
terraform import twilio_taskrouter_workspaces_workflows_v1.assign_to_everyone <TASKROUTER_WORKSPACE_SID>/<TASKROUTER_TARGET_WORKFLOW_SID>
terraform import twilio_taskrouter_workspaces_activities_v1.offline <TASKROUTER_WORKSPACE_SID>/<TASKROUTER_OFFLINE_ACTIVITY_SID>
```

