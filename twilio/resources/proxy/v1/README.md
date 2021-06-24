
## twilio_proxy_services_phone_numbers_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID parent [Service](https://www.twilio.com/docs/proxy/api/service) resource of the new PhoneNumber resource.
**is_reserved** | bool | Optional | Whether the new phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.
**phone_number** | string | Optional | The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
**sid** | string | Optional | The SID of a Twilio [IncomingPhoneNumber](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) resource that represents the Twilio Number you would like to assign to your Proxy Service.

## twilio_proxy_services_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**unique_name** | string | **Required** | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.**
**callback_url** | string | Optional | The URL we should call when the interaction status changes.
**chat_instance_sid** | string | Optional | The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship.
**default_ttl** | int | Optional | The default &#x60;ttl&#x60; value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session&#39;s last create or last Interaction. The default value of &#x60;0&#x60; indicates an unlimited Session length. You can override a Session&#39;s default TTL value by setting its &#x60;ttl&#x60; value.
**geo_match_level** | string | Optional | Where a proxy number must be located relative to the participant identifier. Can be: &#x60;country&#x60;, &#x60;area-code&#x60;, or &#x60;extended-area-code&#x60;. The default value is &#x60;country&#x60; and more specific areas than &#x60;country&#x60; are only available in North America.
**intercept_callback_url** | string | Optional | The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues.
**number_selection_behavior** | string | Optional | The preference for Proxy Number selection in the Service instance. Can be: &#x60;prefer-sticky&#x60; or &#x60;avoid-sticky&#x60; and the default is &#x60;prefer-sticky&#x60;. &#x60;prefer-sticky&#x60; means that we will try and select the same Proxy Number for a given participant if they have previous [Sessions](https://www.twilio.com/docs/proxy/api/session), but we will not fail if that Proxy Number cannot be used.  &#x60;avoid-sticky&#x60; means that we will try to use different Proxy Numbers as long as that is possible within a given pool rather than try and use a previously assigned number.
**out_of_session_callback_url** | string | Optional | The URL we should call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio [function](https://www.twilio.com/functions)) responds with valid [TwiML](https://www.twilio.com/docs/voice/twiml), we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See [Out-of-Session Callback Response Guide](https://www.twilio.com/docs/proxy/out-session-callback-response-guide) for more information.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Service resource to update.

## twilio_proxy_services_sessions_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
**date_expiry** | string | Optional | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the &#x60;ttl&#x60; value.
**fail_on_participant_conflict** | bool | Optional | [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to reject a Session create (with Participants) request that could cause the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. Depending on the context, this could be a 409 error (Twilio error code 80623) or a 400 error (Twilio error code 80604). If not provided, requests will be allowed to succeed and a Debugger notification (80802) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts.
**mode** | string | Optional | The Mode of the Session. Can be: &#x60;message-only&#x60;, &#x60;voice-only&#x60;, or &#x60;voice-and-message&#x60; and the default value is &#x60;voice-and-message&#x60;.
**participants** | list(string) | Optional | The Participant objects to include in the new session.
**status** | string | Optional | The initial status of the Session. Can be: &#x60;open&#x60;, &#x60;in-progress&#x60;, &#x60;closed&#x60;, &#x60;failed&#x60;, or &#x60;unknown&#x60;. The default is &#x60;open&#x60; on create.
**ttl** | int | Optional | The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session&#39;s last Interaction.
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.**
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Session resource to update.

## twilio_proxy_services_short_codes_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
**sid** | string | **Required** | The SID of a Twilio [ShortCode](https://www.twilio.com/docs/sms/api/short-code) resource that represents the short code you would like to assign to your Proxy Service.
**is_reserved** | bool | Optional | Whether the short code should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.

