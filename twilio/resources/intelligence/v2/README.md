
## twilio_intelligence_services_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**unique_name** | string | **Required** | Provides a unique and addressable name to be assigned to this Service, assigned by the developer, to be optionally used in addition to SID.
**auto_transcribe** | bool | Optional | Instructs the Speech Recognition service to automatically transcribe all recordings made on the account.
**data_logging** | bool | Optional | Data logging allows Twilio to improve the quality of the speech recognition through using customer data to refine its speech recognition models.
**friendly_name** | string | Optional | A human readable description of this resource, up to 64 characters.
**language_code** | string | Optional | The default language code of the audio.
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

