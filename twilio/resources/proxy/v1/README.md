
## twilio_proxy_services_sessions_participants_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
**session_sid** | string | **Required** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource.
**identifier** | string | **Required** | The phone number of the Participant.
**friendly_name** | string | Optional | The string that you assigned to describe the participant. This value must be 255 characters or fewer. **This value should not have PII.**
**proxy_identifier** | string | Optional | The proxy phone number to use for the Participant. If not specified, Proxy will select a number from the pool.
**proxy_identifier_sid** | string | Optional | The SID of the Proxy Identifier to assign to the Participant.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Participant resource to fetch.

## twilio_proxy_services_phone_numbers_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID parent [Service](https://www.twilio.com/docs/proxy/api/service) resource of the new PhoneNumber resource.
**sid** | string | Optional | The SID of a Twilio [IncomingPhoneNumber](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) resource that represents the Twilio Number you would like to assign to your Proxy Service.
**phone_number** | string | Optional | The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
**is_reserved** | bool | Optional | Whether the new phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.

## twilio_proxy_services_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**unique_name** | string | **Required** | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.**
**default_ttl** | int | Optional | The default &#x60;ttl&#x60; value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session&#39;s last create or last Interaction. The default value of &#x60;0&#x60; indicates an unlimited Session length. You can override a Session&#39;s default TTL value by setting its &#x60;ttl&#x60; value.
**callback_url** | string | Optional | The URL we should call when the interaction status changes.
**geo_match_level** | string | Optional | 
**number_selection_behavior** | string | Optional | 
**intercept_callback_url** | string | Optional | The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues.
**out_of_session_callback_url** | string | Optional | The URL we should call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio [function](https://www.twilio.com/functions)) responds with valid [TwiML](https://www.twilio.com/docs/voice/twiml), we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See [Out-of-Session Callback Response Guide](https://www.twilio.com/docs/proxy/out-session-callback-response-guide) for more information.
**chat_instance_sid** | string | Optional | The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Service resource to update.

## twilio_proxy_services_sessions_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.**
**date_expiry** | string | Optional | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the &#x60;ttl&#x60; value.
**ttl** | int | Optional | The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session&#39;s last Interaction.
**mode** | string | Optional | 
**status** | string | Optional | 
**participants** | list(string) | Optional | The Participant objects to include in the new session.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Session resource to update.

## twilio_proxy_services_short_codes_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
**sid** | string | **Required** | The SID of a Twilio [ShortCode](https://www.twilio.com/docs/sms/api/short-code) resource that represents the short code you would like to assign to your Proxy Service.
**is_reserved** | bool | Optional | Whether the short code should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.

