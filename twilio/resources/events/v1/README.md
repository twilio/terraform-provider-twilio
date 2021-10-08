
## twilio_events_sinks_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**description** | string | **Required** | A human readable description for the Sink **This value should not contain PII.**
**sink_configuration** | string | **Required** | The information required for Twilio to connect to the provided Sink encoded as JSON.
**sink_type** | string | **Required** | The Sink type. Can only be \\\&quot;kinesis\\\&quot; or \\\&quot;webhook\\\&quot; currently.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this Sink.
**date_created** | string | *Computed* | The date this Sink was created
**date_updated** | string | *Computed* | The date this Sink was updated
**links** | string | *Computed* | Nested resource URLs.
**status** | string | *Computed* | The Status of this Sink
**url** | string | *Computed* | The URL of this resource.

## twilio_events_subscriptions_subscribed_events_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**subscription_sid** | string | **Required** | The unique SID identifier of the Subscription.
**type** | string | **Required** | Type of event being subscribed to.
**schema_version** | int | Optional | The schema version that the subscription should use.
**account_sid** | string | *Computed* | Account SID.
**url** | string | *Computed* | The URL of this resource.

## twilio_events_subscriptions_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**description** | string | **Required** | A human readable description for the Subscription **This value should not contain PII.**
**sink_sid** | string | **Required** | The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created.
**types** | list(string) | **Required** | An array of objects containing the subscribed Event Types
**sid** | string | *Computed* | A 34 character string that uniquely identifies this Subscription.
**account_sid** | string | *Computed* | Account SID.
**date_created** | string | *Computed* | The date this Subscription was created
**date_updated** | string | *Computed* | The date this Subscription was updated
**links** | string | *Computed* | Nested resource URLs.
**url** | string | *Computed* | The URL of this resource.

