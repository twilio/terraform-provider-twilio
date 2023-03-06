
## twilio_ai_services_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**unique_name** | string | **Required** | Provides a unique and addressable name to be assigned to this Service, assigned by the developer, to be optionally used in addition to SID.
**auto_transcribe** | bool | Optional | Instructs the Speech Recognition service to automatically transcribe all recordings made on the account.
**data_logging** | bool | Optional | Data logging allows Twilio to improve the quality of the speech recognition through using customer data to refine its speech recognition models.
**environment** | string | Optional | The environment where the audio is coming from. Values can be 'telephony', 'meeting_room', or 'media_broadcast'.
**friendly_name** | string | Optional | A human readable description of this resource, up to 64 characters.
**language_locale** | string | Optional | The default language locale of the audio.
**word_alternates** | bool | Optional | Displays the next best results for each word of the transcript.
**auto_redaction** | bool | Optional | Instructs the Speech Recognition service to automatically redact PII from all transcripts made on this service.
**media_redaction** | bool | Optional | Instructs the Speech Recognition service to automatically redact PII from all transcripts media made on this service. The auto_redaction flag must be enabled, results in error otherwise.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this Service.
**default_custom_model_sid** | string | Optional | The unique SID identifier of the Default Custom Model.

## twilio_ai_services_transcripts_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The unique SID identifier of the Service.
**data_logging** | bool | Optional | Data logging allows Twilio to improve the quality of the speech recognition through using customer data to refine its speech recognition models.
**model_sid** | string | Optional | The unique SID identifier of the Model to use.
**media_url** | string | Optional | The URL of the media to transcribe.
**recording_sid** | string | Optional | The unique SID identifier of the Recording to use.
**participants** | string | Optional | The array containing the transcript's participants. Participant fields: id, channel, type,role, full_name, email and image_url
**call_direction** | string | Optional | 
**sid** | string | *Computed* | A 34 character string that uniquely identifies this Transcript.

