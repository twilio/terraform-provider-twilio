
## twilio_studio_flows_executions_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**flow_sid** | string | **Required** | The SID of the Excecution&#39;s Flow.
**from** | string | **Required** | The Twilio phone number to send messages or initiate calls from during the Flow&#39;s Execution. Available as variable &#x60;{{flow.channel.address}}&#x60;. For SMS, this can also be a Messaging Service SID.
**to** | string | **Required** | The Contact phone number to start a Studio Flow Execution, available as variable &#x60;{{contact.channel.address}}&#x60;.
**parameters** | string | Optional | JSON data that will be added to the Flow&#39;s context and that can be accessed as variables inside your Flow. For example, if you pass in &#x60;Parameters&#x3D;{\\\&quot;name\\\&quot;:\\\&quot;Zeke\\\&quot;}&#x60;, a widget in your Flow can reference the variable &#x60;{{flow.data.name}}&#x60;, which returns \\\&quot;Zeke\\\&quot;. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string.
**sid** | string | *Computed* | The SID of the Execution resource to update.
**status** | string | Optional | The status of the Execution. Can only be &#x60;ended&#x60;.

