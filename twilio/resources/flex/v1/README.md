
## twilio_flex_channels_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**flex_flow_sid** | string | **Required** | The SID of the Flex Flow.
**identity** | string | **Required** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s chat User.
**chat_user_friendly_name** | string | **Required** | The chat participant&#39;s friendly name.
**chat_friendly_name** | string | **Required** | The chat channel&#39;s friendly name.
**target** | string | Optional | The Target Contact Identity, for example the phone number of an SMS.
**chat_unique_name** | string | Optional | The chat channel&#39;s unique name.
**pre_engagement_data** | string | Optional | The pre-engagement data.
**task_sid** | string | Optional | The SID of the TaskRouter Task. Only valid when integration type is &#x60;task&#x60;. &#x60;null&#x60; for integration types &#x60;studio&#x60; &amp; &#x60;external&#x60;
**task_attributes** | string | Optional | The Task attributes to be added for the TaskRouter Task.
**long_lived** | bool | Optional | Whether to create the channel as long-lived.
**sid** | string | *Computed* | The SID of the Flex chat channel resource to fetch.

## twilio_flex_flex_flows_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A descriptive string that you create to describe the Flex Flow resource.
**chat_service_sid** | string | **Required** | The SID of the chat service.
**channel_type** | string | **Required** | 
**contact_identity** | string | Optional | The channel contact&#39;s Identity.
**enabled** | bool | Optional | Whether the new Flex Flow is enabled.
**integration_type** | string | Optional | 
**integration_flow_sid** | string | Optional | The SID of the Studio Flow. Required when &#x60;integrationType&#x60; is &#x60;studio&#x60;.
**integration_url** | string | Optional | The URL of the external webhook. Required when &#x60;integrationType&#x60; is &#x60;external&#x60;.
**integration_workspace_sid** | string | Optional | The Workspace SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;.
**integration_workflow_sid** | string | Optional | The Workflow SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;.
**integration_channel** | string | Optional | The Task Channel SID (TCXXXX) or unique name (e.g., &#x60;sms&#x60;) to use for the Task that will be created. Applicable and required when &#x60;integrationType&#x60; is &#x60;task&#x60;. The default value is &#x60;default&#x60;.
**integration_timeout** | int | Optional | The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise.
**integration_priority** | int | Optional | The Task priority of a new Task. The default priority is 0. Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise.
**integration_creation_on_message** | bool | Optional | In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging.
**long_lived** | bool | Optional | When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to &#x60;false&#x60;.
**janitor_enabled** | bool | Optional | When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to &#x60;false&#x60;.
**integration_retry_count** | int | Optional | The number of times to retry the Studio Flow or webhook in case of failure. Takes integer values from 0 to 3 with the default being 3. Optional when &#x60;integrationType&#x60; is &#x60;studio&#x60; or &#x60;external&#x60;, not applicable otherwise.
**sid** | string | *Computed* | The SID of the Flex Flow resource to update.

## twilio_flex_web_channels_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**flex_flow_sid** | string | **Required** | The SID of the Flex Flow.
**identity** | string | **Required** | The chat identity.
**customer_friendly_name** | string | **Required** | The chat participant&#39;s friendly name.
**chat_friendly_name** | string | **Required** | The chat channel&#39;s friendly name.
**chat_unique_name** | string | Optional | The chat channel&#39;s unique name.
**pre_engagement_data** | string | Optional | The pre-engagement data.
**sid** | string | *Computed* | The SID of the WebChannel resource to update.
**chat_status** | string | Optional | 
**post_engagement_data** | string | Optional | The post-engagement data.

