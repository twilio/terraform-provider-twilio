
## twilio_intelligence_operators_custom_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A human readable description of the new Operator, up to 64 characters.
**operator_type** | string | **Required** | Operator Type for this Operator. References an existing Operator Type resource.
**config** | string | **Required** | Operator configuration, following the schema defined by the Operator Type.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this Custom Operator.
**if_match** | string | Optional | The If-Match HTTP request header

## twilio_intelligence_services_operators_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The unique SID identifier of the Service.
**operator_sid** | string | **Required** | The unique SID identifier of the Operator. Allows both Custom and Pre-built Operators.

## twilio_intelligence_services_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**unique_name** | string | **Required** | Provides a unique and addressable name to be assigned to this Service, assigned by the developer, to be optionally used in addition to SID.
**auto_transcribe** | bool | Optional | Instructs the Speech Recognition service to automatically transcribe all recordings made on the account.
**data_logging** | bool | Optional | Data logging allows Twilio to improve the quality of the speech recognition & language understanding services through using customer data to refine, fine tune and evaluate machine learning models. Note: Data logging cannot be activated via API, only via www.twilio.com, as it requires additional consent.
**friendly_name** | string | Optional | A human readable description of this resource, up to 64 characters.
**language_code** | string | Optional | The language code set during Service creation determines the Transcription language for all call recordings processed by that Service. The default is en-US if no language code is set. A Service can only support one language code, and it cannot be updated once it's set.
**auto_redaction** | bool | Optional | Instructs the Speech Recognition service to automatically redact PII from all transcripts made on this service.
**media_redaction** | bool | Optional | Instructs the Speech Recognition service to automatically redact PII from all transcripts media made on this service. The auto_redaction flag must be enabled, results in error otherwise.
**webhook_url** | string | Optional | The URL Twilio will request when executing the Webhook.
**webhook_http_method** | string | Optional | 
**sid** | string | *Computed* | A 34 character string that uniquely identifies this Service.
**if_match** | string | Optional | The If-Match HTTP request header

## twilio_intelligence_transcripts_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The unique SID identifier of the Service.
**channel** | string | **Required** | JSON object describing Media Channel including Source and Participants
**customer_key** | string | Optional | Used to store client provided metadata. Maximum of 64 double-byte UTF8 characters.
**media_start_time** | string | Optional | The date that this Transcript's media was started, given in ISO 8601 format.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this Transcript.

