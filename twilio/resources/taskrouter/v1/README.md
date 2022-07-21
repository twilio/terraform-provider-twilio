
## twilio_taskrouter_workspaces_activities_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**workspace_sid** | string | **Required** | The SID of the Workspace that the new Activity belongs to.
**friendly_name** | string | **Required** | A descriptive string that you create to describe the Activity resource. It can be up to 64 characters long. These names are used to calculate and expose statistics about Workers, and provide visibility into the state of each Worker. Examples of friendly names include: &#x60;on-call&#x60;, &#x60;break&#x60;, and &#x60;email&#x60;.
**available** | bool | Optional | Whether the Worker should be eligible to receive a Task when it occupies the Activity. A value of &#x60;true&#x60;, &#x60;1&#x60;, or &#x60;yes&#x60; specifies the Activity is available. All other values specify that it is not. The value cannot be changed after the Activity is created.
**sid** | string | *Computed* | The SID of the Activity resource to update.

## twilio_taskrouter_workspaces_tasks_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**workspace_sid** | string | **Required** | The SID of the Workspace that the new Task belongs to.
**timeout** | int | Optional | The amount of time in seconds the new task can live before being assigned. Can be up to a maximum of 2 weeks (1,209,600 seconds). The default value is 24 hours (86,400 seconds). On timeout, the &#x60;task.canceled&#x60; event will fire with description &#x60;Task TTL Exceeded&#x60;.
**priority** | int | Optional | The priority to assign the new task and override the default. When supplied, the new Task will have this priority unless it matches a Workflow Target with a Priority set. When not supplied, the new Task will have the priority of the matching Workflow Target. Value can be 0 to 2^31^ (2,147,483,647).
**task_channel** | string | Optional | When MultiTasking is enabled, specify the TaskChannel by passing either its &#x60;unique_name&#x60; or &#x60;sid&#x60;. Default value is &#x60;default&#x60;.
**workflow_sid** | string | Optional | The SID of the Workflow that you would like to handle routing for the new Task. If there is only one Workflow defined for the Workspace that you are posting the new task to, this parameter is optional.
**attributes** | string | Optional | A URL-encoded JSON string with the attributes of the new task. This value is passed to the Workflow&#39;s &#x60;assignment_callback_url&#x60; when the Task is assigned to a Worker. For example: &#x60;{ \\\&quot;task_type\\\&quot;: \\\&quot;call\\\&quot;, \\\&quot;twilio_call_sid\\\&quot;: \\\&quot;CAxxx\\\&quot;, \\\&quot;customer_ticket_number\\\&quot;: \\\&quot;12345\\\&quot; }&#x60;.
**sid** | string | *Computed* | The SID of the Task resource to update.
**if_match** | string | Optional | If provided, applies this mutation if (and only if) the [ETag](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag) header of the Task matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
**assignment_status** | string | Optional | 
**reason** | string | Optional | The reason that the Task was canceled or completed. This parameter is required only if the Task is canceled or completed. Setting this value queues the task for deletion and logs the reason.

## twilio_taskrouter_workspaces_task_channels_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**workspace_sid** | string | **Required** | The SID of the Workspace that the new Task Channel belongs to.
**friendly_name** | string | **Required** | A descriptive string that you create to describe the Task Channel. It can be up to 64 characters long.
**unique_name** | string | **Required** | An application-defined string that uniquely identifies the Task Channel, such as &#x60;voice&#x60; or &#x60;sms&#x60;.
**channel_optimized_routing** | bool | Optional | Whether the Task Channel should prioritize Workers that have been idle. If &#x60;true&#x60;, Workers that have been idle the longest are prioritized.
**sid** | string | *Computed* | The SID of the Task Channel resource to update.

## twilio_taskrouter_workspaces_task_queues_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**workspace_sid** | string | **Required** | The SID of the Workspace that the new TaskQueue belongs to.
**friendly_name** | string | **Required** | A descriptive string that you create to describe the TaskQueue. For example &#x60;Support-Tier 1&#x60;, &#x60;Sales&#x60;, or &#x60;Escalation&#x60;.
**target_workers** | string | Optional | A string that describes the Worker selection criteria for any Tasks that enter the TaskQueue. For example, &#x60;&#39;\\\&quot;language\\\&quot; &#x3D;&#x3D; \\\&quot;spanish\\\&quot;&#39;&#x60;. The default value is &#x60;1&#x3D;&#x3D;1&#x60;. If this value is empty, Tasks will wait in the TaskQueue until they are deleted or moved to another TaskQueue. For more information about Worker selection, see [Describing Worker selection criteria](https://www.twilio.com/docs/taskrouter/api/taskqueues#target-workers).
**max_reserved_workers** | int | Optional | The maximum number of Workers to reserve for the assignment of a Task in the queue. Can be an integer between 1 and 50, inclusive and defaults to 1.
**task_order** | string | Optional | 
**reservation_activity_sid** | string | Optional | The SID of the Activity to assign Workers when a task is reserved for them.
**assignment_activity_sid** | string | Optional | The SID of the Activity to assign Workers when a task is assigned to them.
**sid** | string | *Computed* | The SID of the TaskQueue resource to update.

## twilio_taskrouter_workspaces_workers_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**workspace_sid** | string | **Required** | The SID of the Workspace that the new Worker belongs to.
**friendly_name** | string | **Required** | A descriptive string that you create to describe the new Worker. It can be up to 64 characters long.
**activity_sid** | string | Optional | The SID of a valid Activity that will describe the new Worker&#39;s initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information. If not provided, the new Worker&#39;s initial state is the &#x60;default_activity_sid&#x60; configured on the Workspace.
**attributes** | string | Optional | A valid JSON string that describes the new Worker. For example: &#x60;{ \\\&quot;email\\\&quot;: \\\&quot;Bob@example.com\\\&quot;, \\\&quot;phone\\\&quot;: \\\&quot;+5095551234\\\&quot; }&#x60;. This data is passed to the &#x60;assignment_callback_url&#x60; when TaskRouter assigns a Task to the Worker. Defaults to {}.
**sid** | string | *Computed* | The SID of the Worker resource to update.
**if_match** | string | Optional | The If-Match HTTP request header
**reject_pending_reservations** | bool | Optional | Whether to reject the Worker&#39;s pending reservations. This option is only valid if the Worker&#39;s new [Activity](https://www.twilio.com/docs/taskrouter/api/activity) resource has its &#x60;availability&#x60; property set to &#x60;False&#x60;.

## twilio_taskrouter_workspaces_workflows_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**workspace_sid** | string | **Required** | The SID of the Workspace that the new Workflow to create belongs to.
**friendly_name** | string | **Required** | A descriptive string that you create to describe the Workflow resource. For example, &#x60;Inbound Call Workflow&#x60; or &#x60;2014 Outbound Campaign&#x60;.
**configuration** | string | **Required** | A JSON string that contains the rules to apply to the Workflow. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information.
**assignment_callback_url** | string | Optional | The URL from your application that will process task assignment events. See [Handling Task Assignment Callback](https://www.twilio.com/docs/taskrouter/handle-assignment-callbacks) for more details.
**fallback_assignment_callback_url** | string | Optional | The URL that we should call when a call to the &#x60;assignment_callback_url&#x60; fails.
**task_reservation_timeout** | int | Optional | How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to &#x60;86,400&#x60; (24 hours) and the default is &#x60;120&#x60;.
**sid** | string | *Computed* | The SID of the Workflow resource to update.
**re_evaluate_tasks** | string | Optional | Whether or not to re-evaluate Tasks. The default is &#x60;false&#x60;, which means Tasks in the Workflow will not be processed through the assignment loop again.

## twilio_taskrouter_workspaces_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A descriptive string that you create to describe the Workspace resource. It can be up to 64 characters long. For example: &#x60;Customer Support&#x60; or &#x60;2014 Election Campaign&#x60;.
**event_callback_url** | string | Optional | The URL we should call when an event occurs. If provided, the Workspace will publish events to this URL, for example, to collect data for reporting. See [Workspace Events](https://www.twilio.com/docs/taskrouter/api/event) for more information. This parameter supports Twilio&#39;s [Webhooks (HTTP callbacks) Connection Overrides](https://www.twilio.com/docs/usage/webhooks/webhooks-connection-overrides).
**events_filter** | string | Optional | The list of Workspace events for which to call event_callback_url. For example, if &#x60;EventsFilter&#x3D;task.created, task.canceled, worker.activity.update&#x60;, then TaskRouter will call event_callback_url only when a task is created, canceled, or a Worker activity is updated.
**multi_task_enabled** | bool | Optional | Whether to enable multi-tasking. Can be: &#x60;true&#x60; to enable multi-tasking, or &#x60;false&#x60; to disable it. However, all workspaces should be created as multi-tasking. The default is &#x60;true&#x60;. Multi-tasking allows Workers to handle multiple Tasks simultaneously. When enabled (&#x60;true&#x60;), each Worker can receive parallel reservations up to the per-channel maximums defined in the Workers section. In single-tasking mode (legacy mode), each Worker will only receive a new reservation when the previous task is completed. Learn more at [Multitasking](https://www.twilio.com/docs/taskrouter/multitasking).
**template** | string | Optional | An available template name. Can be: &#x60;NONE&#x60; or &#x60;FIFO&#x60; and the default is &#x60;NONE&#x60;. Pre-configures the Workspace with the Workflow and Activities specified in the template. &#x60;NONE&#x60; will create a Workspace with only a set of default activities. &#x60;FIFO&#x60; will configure TaskRouter with a set of default activities and a single TaskQueue for first-in, first-out distribution, which can be useful when you are getting started with TaskRouter.
**prioritize_queue_order** | string | Optional | 
**sid** | string | *Computed* | The SID of the Workspace resource to update.
**default_activity_sid** | string | Optional | The SID of the Activity that will be used when new Workers are created in the Workspace.
**timeout_activity_sid** | string | Optional | The SID of the Activity that will be assigned to a Worker when a Task reservation times out without a response.

