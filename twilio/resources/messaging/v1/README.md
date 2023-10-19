
## twilio_messaging_services_alpha_senders_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.
**alpha_sender** | string | **Required** | The Alphanumeric Sender ID string. Can be up to 11 characters long. Valid characters are A-Z, a-z, 0-9, space, hyphen `-`, plus `+`, underscore `_` and ampersand `&`. This value cannot contain only numbers.
**sid** | string | *Computed* | The SID of the AlphaSender resource to fetch.

## twilio_messaging_services_phone_numbers_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.
**phone_number_sid** | string | **Required** | The SID of the Phone Number being added to the Service.
**sid** | string | *Computed* | The SID of the PhoneNumber resource to fetch.

## twilio_messaging_services_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**inbound_request_url** | string | Optional | The URL we call using `inbound_method` when a message is received by any phone number or short code in the Service. When this property is `null`, receiving inbound messages is disabled. All messages sent to the Twilio phone number or short code will not be logged and received on the Account. If the `use_inbound_webhook_on_number` field is enabled then the webhook url defined on the phone number will override the `inbound_request_url` defined for the Messaging Service.
**inbound_method** | string | Optional | The HTTP method we should use to call `inbound_request_url`. Can be `GET` or `POST` and the default is `POST`.
**fallback_url** | string | Optional | The URL that we call using `fallback_method` if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. If the `use_inbound_webhook_on_number` field is enabled then the webhook url defined on the phone number will override the `fallback_url` defined for the Messaging Service.
**fallback_method** | string | Optional | The HTTP method we should use to call `fallback_url`. Can be: `GET` or `POST`.
**status_callback** | string | Optional | The URL we should call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery.
**sticky_sender** | bool | Optional | Whether to enable [Sticky Sender](https://www.twilio.com/docs/messaging/services#sticky-sender) on the Service instance.
**mms_converter** | bool | Optional | Whether to enable the [MMS Converter](https://www.twilio.com/docs/messaging/services#mms-converter) for messages sent through the Service instance.
**smart_encoding** | bool | Optional | Whether to enable [Smart Encoding](https://www.twilio.com/docs/messaging/services#smart-encoding) for messages sent through the Service instance.
**scan_message_content** | string | Optional | 
**fallback_to_long_code** | bool | Optional | [OBSOLETE] Former feature used to fallback to long code sender after certain short code message failures.
**area_code_geomatch** | bool | Optional | Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/messaging/services#area-code-geomatch) on the Service Instance.
**validity_period** | int | Optional | How long, in seconds, messages sent from the Service are valid. Can be an integer from `1` to `14,400`.
**synchronous_validation** | bool | Optional | Reserved.
**usecase** | string | Optional | A string that describes the scenario in which the Messaging Service will be used. Possible values are `notifications`, `marketing`, `verification`, `discussion`, `poll`, `undeclared`.
**use_inbound_webhook_on_number** | bool | Optional | A boolean value that indicates either the webhook url configured on the phone number will be used or `inbound_request_url`/`fallback_url` url will be called when a message is received from the phone number. If this field is enabled then the webhook url defined on the phone number will override the `inbound_request_url`/`fallback_url` defined for the Messaging Service.
**sid** | string | *Computed* | The SID of the Service resource to update.

## twilio_messaging_services_short_codes_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.
**short_code_sid** | string | **Required** | The SID of the ShortCode resource being added to the Service.
**sid** | string | *Computed* | The SID of the ShortCode resource to fetch.

## twilio_messaging_services_compliance_usa2p_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**messaging_service_sid** | string | **Required** | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/api/service-resource) to create the resources from.
**brand_registration_sid** | string | **Required** | A2P Brand Registration SID
**description** | string | **Required** | A short description of what this SMS campaign does. Min length: 40 characters. Max length: 4096 characters.
**message_flow** | string | **Required** | Required for all Campaigns. Details around how a consumer opts-in to their campaign, therefore giving consent to receive their messages. If multiple opt-in methods can be used for the same campaign, they must all be listed. 40 character minimum. 2048 character maximum.
**message_samples** | list(string) | **Required** | An array of sample message strings, min two and max five. Min length for each sample: 20 chars. Max length for each sample: 1024 chars.
**us_app_to_person_usecase** | string | **Required** | A2P Campaign Use Case. Examples: [ 2FA, EMERGENCY, MARKETING..]
**has_embedded_links** | bool | **Required** | Indicates that this SMS campaign will send messages that contain links.
**has_embedded_phone** | bool | **Required** | Indicates that this SMS campaign will send messages that contain phone numbers.
**opt_in_message** | string | Optional | If end users can text in a keyword to start receiving messages from this campaign, the auto-reply messages sent to the end users must be provided. The opt-in response should include the Brand name, confirmation of opt-in enrollment to a recurring message campaign, how to get help, and clear description of how to opt-out. This field is required if end users can text in a keyword to start receiving messages from this campaign. 20 character minimum. 320 character maximum.
**opt_out_message** | string | Optional | Upon receiving the opt-out keywords from the end users, Twilio customers are expected to send back an auto-generated response, which must provide acknowledgment of the opt-out request and confirmation that no further messages will be sent. It is also recommended that these opt-out messages include the brand name. This field is required if managing opt out keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). 20 character minimum. 320 character maximum.
**help_message** | string | Optional | When customers receive the help keywords from their end users, Twilio customers are expected to send back an auto-generated response; this may include the brand name and additional support contact information. This field is required if managing help keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). 20 character minimum. 320 character maximum.
**opt_in_keywords** | list(string) | Optional | If end users can text in a keyword to start receiving messages from this campaign, those keywords must be provided. This field is required if end users can text in a keyword to start receiving messages from this campaign. Values must be alphanumeric. 255 character maximum.
**opt_out_keywords** | list(string) | Optional | End users should be able to text in a keyword to stop receiving messages from this campaign. Those keywords must be provided. This field is required if managing opt out keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). Values must be alphanumeric. 255 character maximum.
**help_keywords** | list(string) | Optional | End users should be able to text in a keyword to receive help. Those keywords must be provided as part of the campaign registration request. This field is required if managing help keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). Values must be alphanumeric. 255 character maximum.
**sid** | string | *Computed* | The SID of the US A2P Compliance resource to fetch `QE2c6890da8086d771620e9b13fadeba0b`.

