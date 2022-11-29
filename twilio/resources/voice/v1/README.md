
## twilio_voice_byoc_trunks_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**voice_url** | string | Optional | The URL we should call when the BYOC Trunk receives a call.
**voice_method** | string | Optional | The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`.
**voice_fallback_url** | string | Optional | The URL that we should call when an error occurs while retrieving or executing the TwiML from `voice_url`.
**voice_fallback_method** | string | Optional | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
**status_callback_url** | string | Optional | The URL that we should call to pass status parameters (such as call ended) to your application.
**status_callback_method** | string | Optional | The HTTP method we should use to call `status_callback_url`. Can be: `GET` or `POST`.
**cnam_lookup_enabled** | bool | Optional | Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
**connection_policy_sid** | string | Optional | The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure.
**from_domain_sid** | string | Optional | The SID of the SIP Domain that should be used in the `From` header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \\\"call back\\\" an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \\\"sip.twilio.com\\\".
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the BYOC Trunk resource to update.

## twilio_voice_connection_policies_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**sid** | string | *Computed* | The unique string that we created to identify the Connection Policy resource to update.

## twilio_voice_connection_policies_targets_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**connection_policy_sid** | string | **Required** | The SID of the Connection Policy that owns the Target.
**target** | string | **Required** | The SIP address you want Twilio to route your calls to. This must be a `sip:` schema. `sips` is NOT supported.
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**priority** | int | Optional | The relative importance of the target. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important target.
**weight** | int | Optional | The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. Targets with higher values receive more load than those with lower ones with the same priority.
**enabled** | bool | Optional | Whether the Target is enabled. The default is `true`.
**sid** | string | *Computed* | The unique string that we created to identify the Target resource to update.

## twilio_voice_ip_records_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**ip_address** | string | **Required** | An IP address in dotted decimal notation, IPv4 only.
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**cidr_prefix_length** | int | Optional | An integer representing the length of the [CIDR](https://tools.ietf.org/html/rfc4632) prefix to use with this IP address. By default the entire IP address is used, which for IPv4 is value 32.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the IP Record resource to update.

## twilio_voice_source_ip_mappings_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**ip_record_sid** | string | **Required** | The Twilio-provided string that uniquely identifies the IP Record resource to map from.
**sip_domain_sid** | string | **Required** | The SID of the SIP Domain that the IP Record should be mapped to.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the IP Record resource to update.

