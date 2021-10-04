
## twilio_trunking_trunks_credential_lists_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**trunk_sid** | string | **Required** | The SID of the Trunk to associate the credential list with.
**credential_list_sid** | string | **Required** | The SID of the [Credential List](https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource) that you want to associate with the trunk. Once associated, we will authenticate access to the trunk against this list.
**sid** | string | *Computed* | The unique string that we created to identify the CredentialList resource to fetch.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**friendly_name** | string | *Computed* | The string that you assigned to describe the resource
**url** | string | *Computed* | The absolute URL of the resource

## twilio_trunking_trunks_ip_access_control_lists_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**trunk_sid** | string | **Required** | The SID of the Trunk to associate the IP Access Control List with.
**ip_access_control_list_sid** | string | **Required** | The SID of the [IP Access Control List](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource) that you want to associate with the trunk.
**sid** | string | *Computed* | The unique string that we created to identify the IpAccessControlList resource to fetch.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**friendly_name** | string | *Computed* | The string that you assigned to describe the resource
**url** | string | *Computed* | The absolute URL of the resource

## twilio_trunking_trunks_origination_urls_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**trunk_sid** | string | **Required** | The SID of the Trunk to associate the resource with.
**enabled** | bool | **Required** | Whether the URL is enabled. The default is &#x60;true&#x60;.
**friendly_name** | string | **Required** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**priority** | int | **Required** | The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI.
**sip_url** | string | **Required** | The SIP address you want Twilio to route your Origination calls to. This must be a &#x60;sip:&#x60; schema.
**weight** | int | **Required** | The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority.
**sid** | string | *Computed* | The unique string that we created to identify the OriginationUrl resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**url** | string | *Computed* | The absolute URL of the resource

## twilio_trunking_trunks_phone_numbers_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**trunk_sid** | string | **Required** | The SID of the Trunk to associate the phone number with.
**phone_number_sid** | string | **Required** | The SID of the [Incoming Phone Number](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) that you want to associate with the trunk.
**sid** | string | *Computed* | The unique string that we created to identify the PhoneNumber resource to fetch.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**address_requirements** | string | *Computed* | Whether the phone number requires an Address registered with Twilio
**api_version** | string | *Computed* | The API version used to start a new TwiML session
**beta** | string | *Computed* | Whether the phone number is new to the Twilio platform
**capabilities** | string | *Computed* | Indicate if a phone can receive calls or messages
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**friendly_name** | string | *Computed* | The string that you assigned to describe the resource
**links** | string | *Computed* | The URLs of related resources
**phone_number** | string | *Computed* | The phone number in E.164 format
**sms_application_sid** | string | *Computed* | The SID of the application that handles SMS messages sent to the phone number
**sms_fallback_method** | string | *Computed* | The HTTP method used with sms_fallback_url
**sms_fallback_url** | string | *Computed* | The URL that we call when an error occurs while retrieving or executing the TwiML
**sms_method** | string | *Computed* | The HTTP method to use with sms_url
**sms_url** | string | *Computed* | The URL we call when the phone number receives an incoming SMS message
**status_callback** | string | *Computed* | The URL to send status information to your application
**status_callback_method** | string | *Computed* | The HTTP method we use to call status_callback
**url** | string | *Computed* | The absolute URL of the resource
**voice_application_sid** | string | *Computed* | The SID of the application that handles calls to the phone number
**voice_caller_id_lookup** | string | *Computed* | Whether to lookup the caller&#39;s name
**voice_fallback_method** | string | *Computed* | The HTTP method that we use to call voice_fallback_url
**voice_fallback_url** | string | *Computed* | The URL we call when an error occurs in TwiML
**voice_method** | string | *Computed* | The HTTP method used with the voice_url
**voice_url** | string | *Computed* | The URL we call when the phone number receives a call

## twilio_trunking_trunks_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**cnam_lookup_enabled** | bool | Optional | Whether Caller ID Name (CNAM) lookup should be enabled for the trunk. If enabled, all inbound calls to the SIP Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
**disaster_recovery_method** | string | Optional | The HTTP method we should use to call the &#x60;disaster_recovery_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**disaster_recovery_url** | string | Optional | The URL we should call using the &#x60;disaster_recovery_method&#x60; if an error occurs while sending SIP traffic towards the configured Origination URL. We retrieve TwiML from the URL and execute the instructions like any other normal TwiML call. See [Disaster Recovery](https://www.twilio.com/docs/sip-trunking#disaster-recovery) for more information.
**domain_name** | string | Optional | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and &#x60;-&#x60; and must end with &#x60;pstn.twilio.com&#x60;. See [Termination Settings](https://www.twilio.com/docs/sip-trunking#termination) for more information.
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**secure** | bool | Optional | Whether Secure Trunking is enabled for the trunk. If enabled, all calls going through the trunk will be secure using SRTP for media and TLS for signaling. If disabled, then RTP will be used for media. See [Secure Trunking](https://www.twilio.com/docs/sip-trunking#securetrunking) for more information.
**transfer_mode** | string | Optional | The call transfer settings for the trunk. Can be: &#x60;enable-all&#x60;, &#x60;sip-only&#x60; and &#x60;disable-all&#x60;. See [Transfer](https://www.twilio.com/docs/sip-trunking/call-transfer) for more information.
**sid** | string | *Computed* | The unique string that we created to identify the OriginationUrl resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**auth_type** | string | *Computed* | The types of authentication mapped to the domain
**auth_type_set** | list(string) | *Computed* | Reserved
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**links** | string | *Computed* | The URLs of related resources
**recording** | string | *Computed* | The recording settings for the trunk
**url** | string | *Computed* | The absolute URL of the resource

