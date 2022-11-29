
## twilio_studio_flows_executions_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**flow_sid** | string | **Required** | The SID of the Excecution's Flow.
**to** | string | **Required** | The Contact phone number to start a Studio Flow Execution, available as variable `{{contact.channel.address}}`.
**from** | string | **Required** | The Twilio phone number to send messages or initiate calls from during the Flow's Execution. Available as variable `{{flow.channel.address}}`. For SMS, this can also be a Messaging Service SID.
**parameters** | string | Optional | JSON data that will be added to the Flow's context and that can be accessed as variables inside your Flow. For example, if you pass in `Parameters={\\\"name\\\":\\\"Zeke\\\"}`, a widget in your Flow can reference the variable `{{flow.data.name}}`, which returns \\\"Zeke\\\". Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string.
**sid** | string | *Computed* | The SID of the Execution resource to update.
**status** | string | Optional | 

## twilio_studio_flows_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | The string that you assigned to describe the Flow.
**status** | string | **Required** | 
**definition** | string | **Required** | JSON representation of flow definition.
**commit_message** | string | Optional | Description of change made in the revision.
**sid** | string | *Computed* | The SID of the Flow resource to fetch.

