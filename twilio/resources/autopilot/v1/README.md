
## twilio_autopilot_assistants_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**callback_events** | string | Optional | Reserved.
**callback_url** | string | Optional | Reserved.
**defaults** | string | Optional | A JSON object that defines the Assistant&#39;s [default tasks](https://www.twilio.com/docs/autopilot/api/assistant/defaults) for various scenarios, including initiation actions and fallback tasks.
**friendly_name** | string | Optional | A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long.
**log_queries** | bool | Optional | Whether queries should be logged and kept after training. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;. If &#x60;true&#x60;, queries are stored for 30 days, and then deleted. If &#x60;false&#x60;, no queries are stored.
**style_sheet** | string | Optional | The JSON string that defines the Assistant&#39;s [style sheet](https://www.twilio.com/docs/autopilot/api/assistant/stylesheet)
**unique_name** | string | Optional | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Assistant resource to update.
**development_stage** | string | Optional | A string describing the state of the assistant.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**latest_model_build_sid** | string | *Computed* | Reserved
**links** | string | *Computed* | A list of the URLs of the Assistant&#39;s related resources
**needs_model_build** | bool | *Computed* | Whether model needs to be rebuilt
**url** | string | *Computed* | The absolute URL of the Assistant resource

## twilio_autopilot_assistants_tasks_fields_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**assistant_sid** | string | **Required** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the new resource.
**task_sid** | string | **Required** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the new Field resource.
**field_type** | string | **Required** | The Field Type of the new field. Can be: a [Built-in Field Type](https://www.twilio.com/docs/autopilot/built-in-field-types), the &#x60;unique_name&#x60;, or the &#x60;sid&#x60; of a custom Field Type.
**unique_name** | string | **Required** | An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Field resource to fetch.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**url** | string | *Computed* | The absolute URL of the Field resource

## twilio_autopilot_assistants_field_types_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**assistant_sid** | string | **Required** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.
**unique_name** | string | **Required** | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique.
**friendly_name** | string | Optional | A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the FieldType resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**links** | string | *Computed* | A list of the URLs of related resources
**url** | string | *Computed* | The absolute URL of the FieldType resource

## twilio_autopilot_assistants_field_types_field_values_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**assistant_sid** | string | **Required** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the new resource.
**field_type_sid** | string | **Required** | The SID of the Field Type associated with the Field Value.
**language** | string | **Required** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: &#x60;en-US&#x60;
**value** | string | **Required** | The Field Value data.
**synonym_of** | string | Optional | The string value that indicates which word the field value is a synonym of.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the FieldValue resource to fetch.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**url** | string | *Computed* | The absolute URL of the FieldValue resource

## twilio_autopilot_assistants_model_builds_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**assistant_sid** | string | **Required** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.
**status_callback** | string | Optional | The URL we should call using a POST method to send status information to your application.
**unique_name** | string | Optional | An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the ModelBuild resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**build_duration** | int | *Computed* | The time in seconds it took to build the model
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**error_code** | int | *Computed* | More information about why the model build failed, if &#x60;status&#x60; is &#x60;failed&#x60;
**status** | string | *Computed* | The status of the model build process
**url** | string | *Computed* | The absolute URL of the ModelBuild resource

## twilio_autopilot_assistants_queries_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**assistant_sid** | string | **Required** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.
**language** | string | **Required** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the new query. For example: &#x60;en-US&#x60;.
**query** | string | **Required** | The end-user&#39;s natural language input. It can be up to 2048 characters long.
**model_build** | string | Optional | The SID or unique name of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) to be queried.
**tasks** | string | Optional | The list of tasks to limit the new query to. Tasks are expressed as a comma-separated list of task &#x60;unique_name&#x60; values. For example, &#x60;task-unique_name-1, task-unique_name-2&#x60;. Listing specific tasks is useful to constrain the paths that a user can take.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Query resource to update.
**sample_sid** | string | Optional | The SID of an optional reference to the [Sample](https://www.twilio.com/docs/autopilot/api/task-sample) created from the query.
**status** | string | Optional | The new status of the resource. Can be: &#x60;pending-review&#x60;, &#x60;reviewed&#x60;, or &#x60;discarded&#x60;
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**dialogue_sid** | string | *Computed* | The SID of the [Dialogue](https://www.twilio.com/docs/autopilot/api/dialogue).
**model_build_sid** | string | *Computed* | The SID of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) queried
**results** | string | *Computed* | The natural language analysis results that include the Task recognized and a list of identified Fields
**source_channel** | string | *Computed* | The communication channel from where the end-user input came
**url** | string | *Computed* | The absolute URL of the Query resource

## twilio_autopilot_assistants_tasks_samples_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**assistant_sid** | string | **Required** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the new resource.
**task_sid** | string | **Required** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to create.
**language** | string | **Required** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the new sample. For example: &#x60;en-US&#x60;.
**tagged_text** | string | **Required** | The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging).
**source_channel** | string | Optional | The communication channel from which the new sample was captured. Can be: &#x60;voice&#x60;, &#x60;sms&#x60;, &#x60;chat&#x60;, &#x60;alexa&#x60;, &#x60;google-assistant&#x60;, &#x60;slack&#x60;, or null if not included.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Sample resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**url** | string | *Computed* | The absolute URL of the Sample resource

## twilio_autopilot_assistants_tasks_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**assistant_sid** | string | **Required** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.
**unique_name** | string | **Required** | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. This value must be unique and 64 characters or less in length.
**actions** | string | Optional | The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task. It is optional and not unique.
**actions_url** | string | Optional | The URL from which the Assistant can fetch actions.
**friendly_name** | string | Optional | A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Task resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**links** | string | *Computed* | A list of the URLs of related resources
**url** | string | *Computed* | The absolute URL of the Task resource

## twilio_autopilot_assistants_webhooks_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**assistant_sid** | string | **Required** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.
**events** | string | **Required** | The list of space-separated events that this Webhook will subscribe to.
**unique_name** | string | **Required** | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. This value must be unique and 64 characters or less in length.
**webhook_url** | string | **Required** | The URL associated with this Webhook.
**webhook_method** | string | Optional | The method to be used when calling the webhook&#39;s URL.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Webhook resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**url** | string | *Computed* | The absolute URL of the Webhook resource

