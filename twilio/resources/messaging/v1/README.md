
## twilio_messaging_services_alpha_senders_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.
**alpha_sender** | string | **Required** | The Alphanumeric Sender ID string. Can be up to 11 characters long. Valid characters are A-Z, a-z, 0-9, space, hyphen &#x60;-&#x60;, plus &#x60;+&#x60;, underscore &#x60;_&#x60; and ampersand &#x60;&amp;&#x60;. This value cannot contain only numbers.
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
**inbound_request_url** | string | Optional | The URL we call using &#x60;inbound_method&#x60; when a message is received by any phone number or short code in the Service. When this property is &#x60;null&#x60;, receiving inbound messages is disabled. All messages sent to the Twilio phone number or short code will not be logged and received on the Account. If the &#x60;use_inbound_webhook_on_number&#x60; field is enabled then the webhook url defined on the phone number will override the &#x60;inbound_request_url&#x60; defined for the Messaging Service.
**inbound_method** | string | Optional | The HTTP method we should use to call &#x60;inbound_request_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
**fallback_url** | string | Optional | The URL that we call using &#x60;fallback_method&#x60; if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. If the &#x60;use_inbound_webhook_on_number&#x60; field is enabled then the webhook url defined on the phone number will override the &#x60;fallback_url&#x60; defined for the Messaging Service.
**fallback_method** | string | Optional | The HTTP method we should use to call &#x60;fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**status_callback** | string | Optional | The URL we should call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery.
**sticky_sender** | bool | Optional | Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance.
**mms_converter** | bool | Optional | Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance.
**smart_encoding** | bool | Optional | Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance.
**scan_message_content** | string | Optional | 
**fallback_to_long_code** | bool | Optional | Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance.
**area_code_geomatch** | bool | Optional | Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance.
**validity_period** | int | Optional | How long, in seconds, messages sent from the Service are valid. Can be an integer from &#x60;1&#x60; to &#x60;14,400&#x60;.
**synchronous_validation** | bool | Optional | Reserved.
**usecase** | string | Optional | A string that describes the scenario in which the Messaging Service will be used. Examples: [notification, marketing, verification, poll ..].
**use_inbound_webhook_on_number** | bool | Optional | A boolean value that indicates either the webhook url configured on the phone number will be used or &#x60;inbound_request_url&#x60;/&#x60;fallback_url&#x60; url will be called when a message is received from the phone number. If this field is enabled then the webhook url defined on the phone number will override the &#x60;inbound_request_url&#x60;/&#x60;fallback_url&#x60; defined for the Messaging Service.
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
**messaging_service_sid** | string | **Required** | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to create the resources from.
**brand_registration_sid** | string | **Required** | A2P Brand Registration SID
**description** | string | **Required** | A short description of what this SMS campaign does.
**message_samples** | list(string) | **Required** | Message samples, at least 2 and up to 5 sample messages, &lt;&#x3D;1024 chars each.
**us_app_to_person_usecase** | string | **Required** | A2P Campaign Use Case. Examples: [ 2FA, EMERGENCY, MARKETING..]
**has_embedded_links** | bool | **Required** | Indicates that this SMS campaign will send messages that contain links.
**has_embedded_phone** | bool | **Required** | Indicates that this SMS campaign will send messages that contain phone numbers.
**sid** | string | *Computed* | The SID of the US A2P Compliance resource to fetch &#x60;QE2c6890da8086d771620e9b13fadeba0b&#x60;.

