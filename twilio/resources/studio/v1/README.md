
## twilio_studio_flows_engagements_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**flow_sid** | string | **Required** | The SID of the Flow.
**from** | string | **Required** | The Twilio phone number to send messages or initiate calls from during the Flow Engagement. Available as variable &#x60;{{flow.channel.address}}&#x60;
**to** | string | **Required** | The Contact phone number to start a Studio Flow Engagement, available as variable &#x60;{{contact.channel.address}}&#x60;.
**parameters** | string | Optional | A JSON string we will add to your flow&#39;s context and that you can access as variables inside your flow. For example, if you pass in &#x60;Parameters&#x3D;{&#39;name&#39;:&#39;Zeke&#39;}&#x60; then inside a widget you can reference the variable &#x60;{{flow.data.name}}&#x60; which will return the string &#39;Zeke&#39;. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode your JSON string.
**sid** | string | *Computed* | The SID of the Engagement resource to fetch.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**contact_channel_address** | string | *Computed* | The phone number, SIP address or Client identifier that triggered this Engagement
**contact_sid** | string | *Computed* | The SID of the Contact
**context** | string | *Computed* | The current state of the execution flow
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the Engagement was created
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the Engagement was last updated
**links** | string | *Computed* | The URLs of the Engagement&#39;s nested resources
**status** | string | *Computed* | The status of the Engagement
**url** | string | *Computed* | The absolute URL of the resource

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
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**contact_channel_address** | string | *Computed* | The phone number, SIP address or Client identifier that triggered the Execution
**contact_sid** | string | *Computed* | The SID of the Contact
**context** | string | *Computed* | The current state of the flow
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was last updated
**links** | string | *Computed* | Nested resource URLs
**url** | string | *Computed* | The absolute URL of the resource

